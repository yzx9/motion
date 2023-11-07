package PO

import (
	"github.com/jinzhu/gorm"
	"time"
)

/**
*@Description:
*@Author: BZ
*@date: 2023/11/4 10:12
*@Version: V1.0
 */

type Comment struct {
	Id             int    `gorm:"primary_key;not null;type:auto_increment"`
	UserId         int    `gorm:"not null"`
	VideoId        int64  `gorm:"not null"`
	Content        string `gorm:"type:varchar(255);not null"`
	Level          int8   `gorm:"type:tinyint(1);not null"`
	ParentId       int
	Likes          int
	ResponseUserId int
	CreatedAt      time.Time
}

type CommentDao struct {
	db *gorm.DB
}

func NewCommentDB() CommentDao {
	return CommentDao{
		db: db1,
	}
}

func (dao *CommentDao) InsertComment(v1 Comment) (int, error) {
	var video Video
	dao.db.Model(video).Select("UserId").Find("VideoId=?", v1.UserId)
	v1.ResponseUserId = video.UserId
	dao.db.AutoMigrate(&Comment{})
	err := dao.db.Create(&v1).Error
	return v1.Id, err
}

func (dao *CommentDao) SelectCommentById(id int) (Comment, error) {
	var vv Comment
	err := dao.db.Find(&vv, "id=?", id).Error
	return vv, err
}

func (dao *CommentDao) SelectCommentByVideoId(id int64) ([]Comment, error) {
	var coments []Comment
	err := dao.db.Find(&coments, "VideoId=?", id).Error
	return coments, err
}

func (dao *CommentDao) SelectCommentByUserId(id int) ([]Comment, error) {
	var coments []Comment
	err := dao.db.Find(&coments, "UserId=?", id).Error
	return coments, err
}

func (dao *CommentDao) SelectCommentByResponseId(id int) ([]Comment, error) {
	var coments []Comment
	err := dao.db.Find(&coments, "ResponseUserId=?", id).Error
	return coments, err
}

func (dao *CommentDao) DeleteCommentById(id int) error {
	return dao.db.Delete(Comment{}, "id=?", id).Error
}

func (dao *CommentDao) DeleteCommentByVideoId(id int64) error {
	return dao.db.Delete(Comment{}, "VideoId=?", id).Error
}
