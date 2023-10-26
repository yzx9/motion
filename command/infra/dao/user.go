package dao

import (
	"database/sql"
	"errors"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/yzx9/motion/domain/user/adpter/dto"
)

// PO
type user struct {
	Id       string    `db:"id"`
	Nickname string    `db:"nickname"`
	Avatar   string    `db:"avatar"`
	Mobile   string    `db:"mobile"`
	CreateAt time.Time `db:"create_at"` // DB handled
	UpdateAt time.Time `db:"update_at"` // DB handled
}

type userDAO struct{ db *sqlx.DB }

func newUserDB(db *sqlx.DB) userDAO { return userDAO{db} }

func (dao userDAO) Upsert(u user) error {
	const query = `
	INSERT INTO tbl_user (
		 id, name, avatar, mobile, create_at, update_at)
	VALUES (
		:id,:name,:avatar,:mobile, NOW(),	  NOW())
	ON DUPLICATE KEY UPDATE
		name = :name,
		avatar = :avatar,
		mobile = :mobile,
		update_at = NOW();
	`

	if _, err := dao.db.NamedExec(query, args{
		"id":     u.Id,
		"name":   u.Nickname,
		"avatar": u.Avatar,
		"mobile": u.Mobile,
	}); err != nil {
		return wrapErr(err)
	}
	return nil
}

func (dao userDAO) Get(id string) (user, error) {
	const query = `
SELECT
	id, name, avatar, mobile, create_at, update_at
FROM tbl_user
WHERE id = ?;
`
	return dao.get(query, id)
}

func (dao userDAO) GetByMobile(openId string) (user, error) {
	const query = `
SELECT
	id, name, avatar, mobile, create_at, update_at
FROM tbl_user
WHERE mobile = ?;
`
	return dao.get(query, openId)
}

func (dao userDAO) get(query string, args ...any) (user, error) {
	var u user
	if err := dao.db.Get(&u, query, args...); err != nil {
		return user{}, wrapErr(err)
	}
	return u, nil
}

func (dao userDAO) List(limit, offset int) ([]user, error) {
	const query = `
SELECT
	id, name, avatar, mobile, create_at, update_at
FROM tbl_user
LIMIT ? OFFSET ?;
`
	return dao.list(query, limit, offset)
}

func (dao userDAO) ListByName(search string, limit, offset int) ([]user, error) {
	const query = `
SELECT
	id, name, avatar, mobile, create_at, update_at
FROM tbl_user
WHERE name LIKE ?
LIMIT ? OFFSET ?;
`
	search = "%" + search + "%"
	return dao.list(query, search, limit, offset)
}

func (dao userDAO) list(query string, args ...interface{}) ([]user, error) {
	var users []user
	if err := dao.db.Select(&users, query, args...,
	); err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, wrapErr(err)
	}
	return users, nil
}

func (dao userDAO) Count() (int64, error) {
	const query = `
SELECT COUNT(id)
FROM tbl_user;
`
	return dao.count(query)
}

func (dao userDAO) CountByName(search string) (int64, error) {
	const query = `
SELECT COUNT(id)
FROM tbl_user
WHERE name LIKE ?;
`
	search = "%" + search + "%"
	return dao.count(query, search)
}

func (dao userDAO) count(query string, args ...interface{}) (int64, error) {
	var count int64
	if err := dao.db.Get(&count, query, args...); err != nil {
		return 0, wrapErr(err)
	}
	return count, nil
}

/**
 * user domain
 */

type UserDBForUser struct {
	userDAO
}

func NewUser(db *sqlx.DB) UserDBForUser { return UserDBForUser{userDAO{db}} }

func (UserDBForUser) fromDTO(u dto.User) user {
	return user{
		Id:       u.Id,
		Nickname: u.Nickname,
		Avatar:   u.Avatar,
		Mobile:   u.Mobile,
	}
}

func (UserDBForUser) toDTO(u user) dto.User {
	return dto.User{
		Id:       u.Id,
		Nickname: u.Nickname,
		Avatar:   u.Avatar,
		Mobile:   u.Mobile,
	}
}
