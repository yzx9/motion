package service

import (
	"errors"
	"github.com/yzx9/motion/command/domain/video/adapter/dto"
	"github.com/yzx9/motion/command/infra/dao/PO"
)

/**
*@Description:
*@Author: BZ
*@date: 2023/11/6 19:52
*@Version: V1.0
 */
var likeDao = PO.NewLikeDB()

type LikeService interface {
	Like(dto *dto.LikeDto) error
}

type LikeServiceImpl struct {
}

func (service *LikeServiceImpl) Like(dto *dto.LikeDto) error {
	if dto == nil {
		return errors.New("无点赞信息")
	}
	var likeVideo PO.LikeVideo
	if err := dto.ToLike(&likeVideo); err != nil {
		return err
	}
	return likeDao.InsertLikeVideo(&likeVideo)
}
