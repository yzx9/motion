package service

import (
	"errors"
	"github.com/yzx9/motion/command/domain/user/adpter/dto"
	"github.com/yzx9/motion/command/infra/dao/PO"
)

/**
*@Description:
*@Author: BZ
*@date: 2023/11/6 0:32
*@Version: V1.0
 */

type FollowerService interface {
	Follow(dto *dto.FollowDto) error
}

var followerDao = PO.NewFollowerDB()

type FollowerServiceImpl struct {
}

func (service *FollowerServiceImpl) Follow(dto *dto.FollowDto) error {
	if dto == nil {
		return errors.New("无关注信息")
	}
	var follower PO.Follower
	if err := dto.ToFollower(&follower); err != nil {
		return err
	}
	return followerDao.InsertFollower(&follower)
}
