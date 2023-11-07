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
	Id          int64  `gorm:"primary_key;not null"`
	UserId      int    `gorm:"not null"`
	Title       string `gorm:"type:varchar(255)"`
	Url         string `gorm:"type:varchar(255)"`
	Tag         string `gorm:"type:varchar(255)"`
	Cover       string `gorm:"type:varchar(255)"`
	Description string `gorm:"type:varchar(255)"`
	Location    string `gorm:"type:varchar(255)"`
	Channel     string `gorm:"type:varchar(15)"`
	Status      int8   `gorm:"type:tinyint(1)"`
	Likes       int    `gorm:"type:int"`
	Collects    int    `gorm:"type:int"`
	Comments    int    `gorm:"type:int"`
	Forward     int    `gorm:"type:int"`
	CreatedAt   time.Time
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

func (dao *VideoDao) SelectAllVideos() ([]Video, error) {
	var vv []Video
	err := dao.db.Model(&Video{}).Find(&vv).Error
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

func (dao *VideoDao) GetVideoShowDtoByOnce(userId int, videoId int64) (map[string]interface{}, error) {
	var res map[string]interface{}
	err := dao.db.Raw("SELECT a.Id ,a.Title,a.Url,a.Tag,a.Location,a.Channel,a.Description,a.Cover,a.Status,b.UserId as Id,"+
		" b.UserName as user_name,b.Avatar,c.IsLike,d.IsFollow,a.Likes,a.Comments,a.Forwards "+
		"FROM user a LEFT JOIN video b ON a.userId = b.Id LEFT JOIN like c ON c.VideoId = a.Id"+
		"LEFT JOIN Follower d ON d.FollowerId = b.Id where a.Id <> ? and c.UserId <> ? and d.UserId <> ?", videoId, userId, userId).Scan(&res).Error
	return res, err
}

//func (dao *VideoDao) GetVideoShowDto(userId int, videoId int64) (map[string]interface{}, error) {
//	var res map[string]interface{}
//	err := dao.db.Raw("SELECT a.Id ,a.Title,a.Url,a.Tag,a.Location,a.Channel,a.Description,a.Cover,a.Status,b.UserId as Id,"+
//		" b.UserName as user_name,b.Avatar,a.Likes,a.Comments,a.Forwards "+
//		"FROM user a LEFT JOIN video b ON a.userId = b.Id  where a.Id <> ?", videoId).Scan(&res).Error
//	if err != nil {
//		return nil, err
//	}
//	var follower Follower
//	err = dao.db.Model(&follower).Find(&follower, "")
//	return res, err
//}
