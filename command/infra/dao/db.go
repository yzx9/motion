package dao

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"regexp"

	mysql "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/qustavo/sqlhooks/v2"
	"github.com/sirupsen/logrus"
)

type SortBy string

const (
	ASC  SortBy = "ASC"
	DESC SortBy = "DESC"
)

var (
	db     *sqlx.DB
	logger = logrus.WithField("domain", "infra/dao")
)

const MYSQL_DSN = "" // TODO: change to config

func init() {
	sql.Register("mysql-logrus", sqlhooks.Wrap(&mysql.MySQLDriver{}, &logrusHook{logger}))
}

func New(dsn string, debug bool) (*sqlx.DB, error) {
	driverName := "mysql"
	if debug {
		driverName = "mysql-logrus"
	}

	// exactly the same as the built-in
	db, err := sqlx.Open(driverName, dsn)
	if err != nil {
		return nil, err
	}

	// force a connection and test that it worked
	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

func Setup(debug bool) {
	var err error
	if db, err = New(MYSQL_DSN, debug); err != nil {
		logger.WithError(err).Fatal("db down")
	}
	logger.Debug("connect to db")
}

func DB() *sqlx.DB { return db }

var (
	ErrNotFound         = fmt.Errorf("db: not found")
	ErrDataInconsistent = fmt.Errorf("db: data inconsistent")
)

type ErrDBDown struct{ err error }

func (e *ErrDBDown) Error() string {
	return fmt.Sprintf("db: %s", e.err.Error())
}

func wrapErr(err error) error {
	if err == nil {
		return nil
	} else if errors.Is(err, sql.ErrNoRows) {
		return ErrNotFound
	} else if _, ok := err.(*ErrDBDown); ok {
		return err
	} else {
		return &ErrDBDown{err}
	}
}

// hooks

type logrusHook struct{ *logrus.Entry }

// make sure hook implement `sqlhooks.Hooks`
var _ interface{ sqlhooks.Hooks } = (*logrusHook)(nil)

func (h *logrusHook) Before(ctx context.Context, query string, args ...interface{}) (context.Context, error) {
	h.Debugf("> %s %q", readFriendly(query), args)
	return ctx, nil
}

func (h *logrusHook) After(ctx context.Context, query string, args ...interface{}) (context.Context, error) {
	return ctx, nil
}

var reWhitespace = regexp.MustCompile(`\s+`)
var reTrimWhitespace = regexp.MustCompile(`(^\s+)|(\s+)$`)

// helpers

type args = map[string]any

func transact(db *sqlx.DB, txFunc func(*sqlx.Tx) error) (err error) {
	tx, err := db.Beginx()
	if err != nil {
		return wrapErr(err)
	}

	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p) // re-throw panic after Rollback
		} else if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()

	err = txFunc(tx)
	return wrapErr(err)
}

func readFriendly(s string) []byte {
	b := []byte(s)
	b = reWhitespace.ReplaceAll(b, []byte(" "))
	b = reTrimWhitespace.ReplaceAll(b, []byte(""))
	return b
}
