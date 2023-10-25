package user

import "mime/multipart"

// Aggregate root
type User interface {
	Id() string
	Avatar() Avatar
	Favorites() []Favorite

	SetAvatar(f *multipart.FileHeader) string
	Follow(user User) error
}

func NewUser(nickname string) User {
	return nil // TODO: not implement
}
