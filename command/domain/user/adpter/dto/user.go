package dto

import (
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
	"github.com/yzx9/motion/command/infra/dao/PO"
)

type User struct {
	Id       string
	Nickname string
	Avatar   string
	Mobile   string
}

type LoginDto struct {
	UserName string `json:"user_name" form:"user_name"`
	Mobile   string `json:"mobile" form:"mobile"`
	Avatar   string
	Password string `json:"password" form:"password"`
}

func (dto *LoginDto) ToPoUser(uu *PO.User) error {
	if dto == nil {
		panic("loginDto为空")
	}
	uu.MotionId = uuid.New().ID()
	return copier.Copy(uu, dto)

}
