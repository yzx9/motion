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

type CollectDto struct {
	UserId    int  `json:"user-id,omitempty" form:"user-id"`
	VideoId   int  `json:"video-id,omitempty" form:"video-id"`
	IsCollect bool `json:"is-like,omitempty" form:"is-like"`
}

func (dto *CollectDto) ToCollect(uu *PO.CollectVideo) error {
	if dto == nil {
		panic("CollectDto为空")
	}
	return copier.Copy(uu, dto)
}

func (dto *CollectDto) ToDto(uu *PO.CollectVideo) error {
	if uu == nil {
		panic("Collect为空")
	}
	return copier.Copy(dto, uu)
}
