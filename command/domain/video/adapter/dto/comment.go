package dto

import (
	"github.com/jinzhu/copier"
	"github.com/yzx9/motion/command/infra/dao/PO"
	"time"
)

/**
*@Description:
*@Author: BZ
*@date: 2023/11/6 20:09
*@Version: V1.0
 */

type CommentDto struct {
	UserId    int
	UserName  string `json:"user-name" form:"user-name"`
	VideoId   int64  `json:"video-id,omitempty" form:"video-id"`
	Content   string `json:"content" form:"content"`
	Level     int8   `json:"level" form:"level"`
	ParentId  int    `json:"parent-id" form:"parent-id"`
	Likes     int    `json:"likes" form:"likes"`
	IsLike    bool   `json:"is-like,omitempty" form:"is-like"`
	CreatedAt time.Time
}

func (dto *CommentDto) ToComment(uu *PO.Comment) error {
	if dto == nil {
		panic("CommentDto为空")
	}
	return copier.Copy(uu, dto)
}

func (dto *CommentDto) ToDto(uu *PO.Comment) error {
	if uu == nil {
		panic("comment为空")
	}
	return copier.Copy(dto, uu)
}
