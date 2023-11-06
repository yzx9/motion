package PO

import (
	"errors"
	"github.com/jinzhu/gorm"
	gorm2 "github.com/yzx9/motion/command/infra/dao/gorm"
)

/**
*@Description:
*@Author: BZ
*@date: 2023/11/2 21:29
*@Version: V1.0
 */

type User struct {
	Id        int32  `gorm:"primary_key;type:int auto_increment;not null" copier:"UserId"`
	MotionId  uint32 `gorm:"not null;unique"`
	UserName  string `gorm:"type:varchar(26);notnull;unique"`
	Password  string `gorm:"type:varchar(200);notnull"`
	Avatar    string `gorm:"type:varchar(256)"`
	AvatarId  int64
	Mobile    string `gorm:"type:varchar(256);not"`
	Sex       bool   `gorm:"type:tinyint(1)"`
	Age       uint8  `gorm:"type:tinyint"`
	Fans      int    `gorm:"type:int"`
	Followers int    `gorm:"type:int"`
	Likes     int    `gorm:"type:int"`
	Friends   int    `gorm:"type:int"`
}

var db1 = gorm2.GetDB()

type UserDao struct {
	db *gorm.DB
}

func NewUserDB() UserDao {
	return UserDao{
		db: db1,
	}
}

func (dao *UserDao) InsertUser(u1 *User) error {
	dao.db.AutoMigrate(&User{})
	_, err := dao.SelectUserByMobile(u1.Mobile)
	if err == nil {
		return errors.New("用户已存在")
	}
	return dao.db.Create(u1).Error
}

func (dao *UserDao) SelectUserById(id int) (User, error) {
	var uu User
	err := dao.db.Find(&uu, "id=?", id).Error
	return uu, err
}

func (dao *UserDao) SelectUserByMobile(mobile string) (User, error) {
	var uu User
	err := dao.db.Find(&uu, "mobile=?", mobile).Error
	return uu, err
}

func (dao *UserDao) DeleteUserById(id int) error {
	var uu = User{}
	return dao.db.Delete(&uu, "id=?", id).Error
}

func (dao *UserDao) UpdateUser(user *User) error {
	return dao.db.Model(&User{}).Omit("Id", "MotionId").Updates(user).Error
}
func (dao *UserDao) UpdateUserAvatar(userId int, avatarId int64, avatar string) error {
	return dao.db.Model(&User{}).Where("Id=?", userId).Updates(map[string]interface{}{
		"Avatar":   avatar,
		"AvatarId": avatarId,
	}).Error
}
