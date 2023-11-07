package service

import (
	"fmt"
	"github.com/jinzhu/copier"
	"github.com/yzx9/motion/command/domain/video/adapter/dto"
	"github.com/yzx9/motion/command/infra/dao/PO"
)

/**
*@Description:
*@Author: BZ
*@date: 2023/11/3 17:06
*@Version: V1.0
 */
type VideoService interface {
	SaveVideo(PO.Video) error
	SelectVideoByID(id int64, uid int) (PO.Video, error)
}

var videoDao = PO.NewVideoDB()
var userDao = PO.NewUserDB()
var followDao = PO.NewFollowerDB()

type VideoServiceImpl struct {
}

func (v VideoServiceImpl) SaveVideo(video PO.Video) error {
	return videoDao.InsertVideo(video)
}

func (v VideoServiceImpl) SelectVideoByID(id int64, uid int) (dto.VideoShowDto, error) {
	video, err := videoDao.SelectVideoById(id)
	if err != nil {
		return dto.VideoShowDto{}, err
	}
	user, err := userDao.SelectUserById(video.UserId)
	if err != nil {
		return dto.VideoShowDto{}, err
	}
	var videoDto dto.VideoDto
	err = copier.Copy(&videoDto, &video)
	if err != nil {
		return dto.VideoShowDto{}, err
	}
	var videShowDto dto.VideoShowDto
	err = copier.Copy(&videShowDto, &user)
	if err != nil {
		return dto.VideoShowDto{}, err
	}
	err = copier.Copy(&videShowDto, &video)
	if err != nil {
		return dto.VideoShowDto{}, err
	}
	videShowDto.Video = videoDto
	followed, err := followDao.SelectByUserIdAndFollowerId(uid, videShowDto.UserId)
	if err == nil {
		videShowDto.IsFollow = followed.IsFollow
	}
	videShowDto.IsFollow = false
	like, err := likeDao.SelectByUserIdAndLikeVideoId(uid, videShowDto.Video.Id)
	if err == nil {
		videShowDto.IsLike = like.IsLike
	}
	fmt.Println(err.Error())
	videShowDto.IsLike = false
	return videShowDto, nil
}
