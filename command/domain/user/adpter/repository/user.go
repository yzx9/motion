package repository

import "github.com/yzx9/motion/command/domain/user/adpter/dto"

var UserRepo interface {
	Upsert(dto.User) error
}
