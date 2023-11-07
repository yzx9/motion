package dto

import (
	"github.com/jinzhu/copier"
	"github.com/yzx9/motion/command/infra/dao/PO"
)

/**
*@Description:
*@Author: BZ
*@date: 2023/11/6 0:34
*@Version: V1.0
 */

type FollowDto struct {
	UserId     int  `json:"user-id,omitempty" form:"user-id"`
	FollowerId int  `json:"follower-id,omitempty" form:"follower-id"`
	IsFollow   bool `json:"is-follow,omitempty" form:"is-follow"`
}

func (dto *FollowDto) ToFollower(uu *PO.Follower) error {
	if dto == nil {
		panic("FollowerDto为空")
	}
	return copier.Copy(uu, dto)
}

func (dto *FollowDto) ToDto(uu *PO.Follower) error {
	if uu == nil {
		panic("Follower为空")
	}
	return copier.Copy(dto, uu)
}
