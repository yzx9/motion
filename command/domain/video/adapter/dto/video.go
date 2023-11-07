package dto

import (
	"github.com/jinzhu/copier"
	"github.com/yzx9/motion/command/infra/dao/PO"
)

/**
*@Description:
*@Author: BZ
*@date: 2023/11/3 16:27
*@Version: V1.0
 */

type VideoDto struct {
	Id          int64  `json:"id" form:"id"`
	Title       string `json:"title" form:"title"`
	Url         string `json:"url" form:"url"`
	Tag         string `json:"tag" form:"tag"`
	Location    string `json:"location" form:"location"`
	Channel     string `json:"channel" form:"channel"`
	Description string `json:"description" form:"description"`
	Cover       string `json:"cover"`
	Status      int8   `json:"status" form:"status"`
}

func (v *VideoDto) ToVideo(video *PO.Video) error {
	return copier.Copy(video, v)
}

func (v *VideoDto) ToDto(video *PO.Video) error {
	return copier.Copy(v, video)
}

type VideoShowDto struct {
	Video    VideoDto `json:"video" form:"video"`
	UserId   int      `json:"user-id" ` //copier:"user-id"
	UserName string   `json:"user-name"`
	Avatar   string   `json:"avatar"`
	IsLike   bool     `json:"is-like"`
	IsFollow bool     `json:"is-follow"`
	Likes    int      `json:"likes"`
	Comments int      `json:"comments"`
	Collects int      `json:"collects"`
	Forwards int      `json:"forwards"`
}
