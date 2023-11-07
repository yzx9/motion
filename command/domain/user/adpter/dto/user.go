package dto

import (
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
	"github.com/yzx9/motion/command/infra/dao/PO"
)

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

type UserDto struct {
	MotionId  uint32 `json:"motionId"`
	UserName  string `json:"user_name"`
	Avatar    string `json:"avatar"`
	Sex       bool   `json:"sex"`
	Age       uint8  `json:"age"`
	Fans      int    `json:"fans"`
	Followers int    `json:"followers"`
	Likes     int    `json:"likes"`
	Collects  int    `json:"collects"`
	Friends   int    `json:"friends"`
}
