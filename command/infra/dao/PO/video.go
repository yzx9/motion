package PO

import (
	"github.com/jinzhu/gorm"
	"time"
)

/**
*@Description:
*@Author: BZ
*@date: 2023/11/2 21:29
*@Version: V1.0
 */

type Video struct {
	Id          int64      `gorm:"primary_key;not null"`
	UserId      int        `gorm:"not null"`
	Title       string     `gorm:"type:varchar(255)"`
	Url         string     `gorm:"type:varchar(255)"`
	Tag         string     `gorm:"type:varchar(255)"`
	Cover       string     `gorm:"type:varchar(255)"`
	Description string     `gorm:"type:varchar(255)"`
	Location    string     `gorm:"type:varchar(255)"`
	Channel     string     `gorm:"type:varchar(15)"`
	Status      int8       `gorm:"type:tinyint(1)"`
	Likes       int        `gorm:"type:int"`
	Collections int        `gorm:"type:int"`
	Comments    int        `gorm:"type:int"`
	Forward     int        `gorm:"type:int"`
	CreateAt    *time.Time `gorm:"not null"`
}

type VideoDao struct {
	db *gorm.DB
}

func NewVideoDB() VideoDao {
	return VideoDao{
		db: db1,
	}
}

func (dao *VideoDao) InsertVideo(v1 Video) error {
	dao.db.AutoMigrate(&Video{})
	return dao.db.Create(&v1).Error
}

func (dao *VideoDao) SelectVideoById(id int64) (Video, error) {
	var vv Video
	err := dao.db.Find(&vv, "id=?", id).Error
	return vv, err
}

func (dao *VideoDao) SelectVideoByTitle(title string) (Video, error) {
	var vv Video
	err := dao.db.Where("title like ?", "%"+title+"%").Find(&vv).Error
	return vv, err
}

func (dao *VideoDao) DeleteVideoById(id int64) error {
	var vv = Video{}
	return dao.db.Delete(&vv, "id=?", id).Error
}

func (dao *VideoDao) UpdateVideo(video Video) error {
	return dao.db.Model(&Video{}).Omit("Id", "MotionId").Updates(video).Error
}
