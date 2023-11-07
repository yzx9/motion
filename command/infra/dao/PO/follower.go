package PO

import (
	"github.com/jinzhu/gorm"
)

/**
*@Description:
*@Author: BZ
*@date: 2023/11/5 23:57
*@Version: V1.0
 */

type Follower struct {
	Id         int  `gorm:"primary_key;type:int auto_increment;not null"`
	UserId     int  `gorm:"not null"`
	FollowerId int  `gorm:"not null"`
	IsFollow   bool `gorm:"not null;type:tinyint(1)"`
}

type FollowerDao struct {
	db *gorm.DB
}

var userDao = NewUserDB()

func NewFollowerDB() FollowerDao {
	return FollowerDao{
		db: db1,
	}
}

func (dao *FollowerDao) InsertFollower(u1 *Follower) error {
	dao.db.AutoMigrate(&Follower{})
	if err := dao.UpdateFollower(u1); err == nil {
		return err
	}
	err := dao.db.Create(u1).Error
	if err != nil {
		return err
	}
	return dao.UpdateFollower(u1)
}

func (dao *FollowerDao) SelectFollowerById(id int) (Follower, error) {
	var uu Follower
	err := dao.db.Find(&uu, "id=?", id).Error
	return uu, err
}

func (dao *FollowerDao) SelectFollowerByUserId(id int) ([]Follower, error) {
	var uu []Follower
	err := dao.db.Find(&uu, "UserId=?", id).Error
	return uu, err
}

func (dao *FollowerDao) SelectFollowerByFollowerId(id int) ([]Follower, error) {
	var uu []Follower
	err := dao.db.Find(&uu, "FollowerId=?", id).Error
	return uu, err
}

func (dao *FollowerDao) DeleteFollowerById(id int) error {
	var uu = Follower{}
	return dao.db.Delete(&uu, "id=?", id).Error
}

func (dao *FollowerDao) UpdateFollower(follower *Follower) error {
	record, err := dao.SelectByUserIdAndFollowerId(follower.UserId, follower.FollowerId)
	if err == nil {
		user, err := userDao.SelectUserById(record.FollowerId)
		if err != nil {
			return err
		}
		if record.IsFollow && !follower.IsFollow {
			if user.Followers > 0 {
				user.Followers -= 1
			}
		}

		if !record.IsFollow && follower.IsFollow {
			user.Followers += 1
		}
		dao.db.Model(&User{}).Where("id=?", record.UserId).Update("Followers", user.Followers)
		err = dao.db.Model(&Follower{}).Where("Id=?", record.Id).Update("IsFollow", follower.IsFollow).Error
		if err != nil {
			return err
		}

	}
	return err
}

func (dao *FollowerDao) SelectByUserIdAndFollowerId(userId int, followerId int) (Follower, error) {
	var follower Follower
	forms := dao.db.Find(&follower, "UserId=?", userId, "FollowerId=?", followerId)
	return follower, forms.Error
}
