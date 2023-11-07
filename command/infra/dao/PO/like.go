package PO

import "github.com/jinzhu/gorm"

/**
*@Description:
*@Author: BZ
*@date: 2023/11/6 17:43
*@Version: V1.0
 */

type LikeVideo struct {
	Id      int   `gorm:"primary_key;type:int auto_increment;not null"`
	UserId  int   `gorm:"not null"`
	VideoId int64 `gorm:"not null"`
	IsLike  bool  `gorm:"not null;type:tinyint(1)"`
}

type LikeVideoDao struct {
	db *gorm.DB
}

func NewLikeDB() LikeVideoDao {
	return LikeVideoDao{
		db: db1,
	}
}

func (dao *LikeVideoDao) InsertLikeVideo(u1 *LikeVideo) error {
	dao.db.AutoMigrate(&LikeVideo{})
	if err := dao.UpdateLikeVideo(u1); err == nil {
		return err
	}
	err := dao.db.Create(u1).Error
	if err != nil {
		return err
	}
	return dao.UpdateLikeVideo(u1)
}

func (dao *LikeVideoDao) SelectLikeVideoById(id int) (LikeVideo, error) {
	var uu LikeVideo
	err := dao.db.Find(&uu, "id=?", id).Error
	return uu, err
}

func (dao *LikeVideoDao) SelectLikeVideoByUserId(id int) ([]LikeVideo, error) {
	var uu []LikeVideo
	err := dao.db.Find(&uu, "UserId=?", id).Error
	return uu, err
}

func (dao *LikeVideoDao) SelectLikeVideoByVideoId(id int64) ([]LikeVideo, error) {
	var uu []LikeVideo
	err := dao.db.Find(&uu, "VideoId=?", id).Error
	return uu, err
}

func (dao *LikeVideoDao) DeleteLikeVideoById(id int) error {
	var uu = LikeVideo{}
	return dao.db.Delete(&uu, "id=?", id).Error
}

func (dao *LikeVideoDao) UpdateLikeVideo(like *LikeVideo) error {
	record, err := dao.SelectByUserIdAndLikeVideoId(like.UserId, like.VideoId)
	if err == nil {
		user, err := userDao.SelectUserById(record.UserId)
		if err != nil {
			return err
		}
		videoDB := NewVideoDB()
		video, err := videoDB.SelectVideoById(like.VideoId)
		if err != nil {
			return err
		}

		if record.IsLike && !like.IsLike {
			if user.Likes > 0 {
				user.Likes -= 1
			}
			if video.Likes > 0 {
				video.Likes -= 1
			}
		}

		if !record.IsLike && like.IsLike {
			user.Likes += 1
			video.Likes += 1
		}

		dao.db.Model(&User{}).Where("id=?", user.Id).Update("Likes", user.Likes)
		dao.db.Model(&Video{}).Where("id=?", video.Id).Update("Likes", video.Likes)
		err = dao.db.Model(&LikeVideo{}).Where("Id=?", record.Id).Update("IsLike", like.IsLike).Error
		if err != nil {
			return err
		}
	}
	return err
}

func (dao *LikeVideoDao) SelectByUserIdAndLikeVideoId(userId int, likeId int64) (LikeVideo, error) {
	var like LikeVideo
	forms := dao.db.Find(&like, "UserId=?", userId, "VideoId=?", likeId)
	return like, forms.Error
}
