package dto

import (
	"github.com/jinzhu/copier"
	"github.com/yzx9/motion/command/infra/dao/PO"
)

/**
*@Description:
*@Author: BZ
*@date: 2023/11/6 17:40
*@Version: V1.0
 */

type LikeDto struct {
	UserId  int  `json:"user-id,omitempty" form:"user-id"`
	VideoId int  `json:"video-id,omitempty" form:"video-id"`
	IsLike  bool `json:"is-like,omitempty" form:"is-like"`
}

func (dto *LikeDto) ToLike(uu *PO.LikeVideo) error {
	if dto == nil {
		panic("LikeDto为空")
	}
	return copier.Copy(uu, dto)
}

func (dto *LikeDto) ToDto(uu *PO.LikeVideo) error {
	if uu == nil {
		panic("Like为空")
	}
	return copier.Copy(dto, uu)
}
