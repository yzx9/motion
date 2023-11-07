package PO

import "github.com/jinzhu/gorm"

/**
*@Description:
*@Author: BZ
*@date: 2023/11/6 17:43
*@Version: V1.0
 */

type CollectVideo struct {
	Id        int   `gorm:"primary_key;type:int auto_increment;not null"`
	UserId    int   `gorm:"not null"`
	VideoId   int64 `gorm:"not null"`
	IsCollect bool  `gorm:"not null;type:tinyint(1)"`
}

type CollectVideoDao struct {
	db *gorm.DB
}

func NewCollectDB() CollectVideoDao {
	return CollectVideoDao{
		db: db1,
	}
}

func (dao *CollectVideoDao) InsertCollectVideo(u1 *CollectVideo) error {
	dao.db.AutoMigrate(&CollectVideo{})
	if err := dao.UpdateCollectVideo(u1); err == nil {
		return err
	}
	err := dao.db.Create(u1).Error
	if err != nil {
		return err
	}
	return dao.UpdateCollectVideo(u1)
}

func (dao *CollectVideoDao) SelectCollectVideoById(id int) (CollectVideo, error) {
	var uu CollectVideo
	err := dao.db.Find(&uu, "id=?", id).Error
	return uu, err
}

func (dao *CollectVideoDao) SelectCollectVideoByUserId(id int) ([]CollectVideo, error) {
	var uu []CollectVideo
	err := dao.db.Find(&uu, "UserId=?", id).Error
	return uu, err
}

func (dao *CollectVideoDao) SelectCollectVideoByVideoId(id int64) ([]CollectVideo, error) {
	var uu []CollectVideo
	err := dao.db.Find(&uu, "VideoId=?", id).Error
	return uu, err
}

func (dao *CollectVideoDao) DeleteCollectVideoById(id int) error {
	var uu = CollectVideo{}
	return dao.db.Delete(&uu, "id=?", id).Error
}

func (dao *CollectVideoDao) UpdateCollectVideo(like *CollectVideo) error {
	record, err := dao.SelectByUserIdAndCollectVideoId(like.UserId, like.VideoId)
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

		if record.IsCollect && !like.IsCollect {
			if user.Collects > 0 {
				user.Collects -= 1
			}
			if video.Collects > 0 {
				video.Collects -= 1
			}
		}

		if !record.IsCollect && like.IsCollect {
			user.Collects += 1
			video.Collects += 1
		}

		dao.db.Model(&User{}).Where("id=?", user.Id).Update("Collects", user.Collects)
		dao.db.Model(&Video{}).Where("id=?", video.Id).Update("Collects", video.Collects)
		err = dao.db.Model(&CollectVideo{}).Where("Id=?", record.Id).Update("IsCollect", like.IsCollect).Error
		if err != nil {
			return err
		}
	}
	return err
}

func (dao *CollectVideoDao) SelectByUserIdAndCollectVideoId(userId int, likeId int64) (CollectVideo, error) {
	var like CollectVideo
	forms := dao.db.Find(&like, "UserId=?", userId, "VideoId=?", likeId)
	return like, forms.Error
}
