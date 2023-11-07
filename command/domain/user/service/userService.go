package service

import (
	"errors"
	"github.com/jinzhu/copier"
	"github.com/yzx9/motion/command/domain/user/adpter/dto"
	"github.com/yzx9/motion/command/infra/common/util"
	"github.com/yzx9/motion/command/infra/dao/PO"
)

/**
*@Description:
*@Author: BZ
*@date: 2023/11/2 19:36
*@Version: V1.0
 */

type UserService interface {
	Login(dto dto.LoginDto) (string, error)
	Register(dto dto.LoginDto) error
	UpdateAvatar(userId int, avatarId int64, avatar string) error
}

var userDao = PO.NewUserDB()

type UserServiceImpl struct {
}

func (u UserServiceImpl) Register(dto dto.LoginDto) error {
	isMatch, err2 := util.MatchRegexpString("^1[0-9]{10}", dto.Mobile)
	if err2 != nil {
		return err2
	}
	if dto.UserName == "" || !isMatch || len(dto.Password) < 6 || len(dto.Password) > 20 {
		return errors.New("输入不符合规范")
	}

	var user PO.User
	err2 = dto.ToPoUser(&user)
	if err2 != nil {
		return err2
	}

	var err error
	user.Password, err = util.GetHashPassword(user.Password)
	if err != nil {
		return err
	}
	return userDao.InsertUser(&user)
}

func (u UserServiceImpl) Login(dto dto.LoginDto) (int, error) {
	user, err := userDao.SelectUserByMobile(dto.Mobile)
	if err == nil {
		//密码确认
		if !util.ComparePassword(dto.Password, user.Password) {
			return -1, errors.New("密码错误")
		}
		//token, err1 := util.GenerateToken(int(user.Id), user.UserName)
		//return token, err1
	}
	return int(user.Id), err
}

func (u UserServiceImpl) UpdateAvatar(userId int, avatarId int64, avatar string) error {
	return userDao.UpdateUserAvatar(userId, avatarId, avatar)

}

func (u UserServiceImpl) GetUserInfo(id int) (dto.UserDto, error) {
	user, err := userDao.SelectUserById(id)
	if err != nil {
		return dto.UserDto{}, err
	}
	var dto dto.UserDto
	err = copier.Copy(&dto, &user)
	return dto, err
}
