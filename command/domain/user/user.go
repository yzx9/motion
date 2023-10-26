package user

import (
	"fmt"
	"mime/multipart"

	"github.com/google/uuid"
	"github.com/yzx9/motion/command/domain/user/adpter/dto"
	"github.com/yzx9/motion/command/domain/user/adpter/repository"
)

// Aggregate root
type User interface {
	Id() string
	Avatar() Avatar
	Favorites() []Favorite

	SetAvatar(f *multipart.FileHeader) error
	Follow(user User) error
}

type user struct {
	id       string
	nickname string
	avatar   string
	mobile   string
}

func NewUser(nickname string) (User, error) {
	user := user{
		id:       uuid.NewString(),
		nickname: nickname,
		avatar:   "", // TODO: default
		mobile:   "", // TODO
	}
	if err := user.save(); err != nil {
		return nil, err
	}
	return user, nil
}

func Get(id string) User {
	return nil // TODO: not implement
}

func (u user) Id() string            { return u.id }
func (u user) Avatar() Avatar        { return emptyAvatar } // TODO
func (u user) Favorites() []Favorite { return nil }         // TODO

func (u user) SetAvatar(f *multipart.FileHeader) error { return fmt.Errorf("not implement") }
func (u user) Follow(user User) error                  { return fmt.Errorf("not implement") }

func (u user) save() error {
	err := repository.UserRepo.Upsert(u.toDTO())
	return wrapErr(err)
}

func (u user) toDTO() dto.User {
	return dto.User{
		Id:       u.id,
		Nickname: u.nickname,
		Avatar:   u.avatar,
		Mobile:   u.mobile,
	}
}
