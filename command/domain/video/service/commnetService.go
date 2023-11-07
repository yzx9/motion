package service

import (
	"errors"
	"github.com/jinzhu/copier"
	"github.com/yzx9/motion/command/domain/video/adapter/dto"
	"github.com/yzx9/motion/command/infra/dao/PO"
)

/**
*@Description:
*@Author: BZ
*@date: 2023/11/6 19:52
*@Version: V1.0
 */
var commentDao = PO.NewCommentDB()

type CommentService interface {
	Comment(*dto.CommentDto) (int, error)
	GetVideoComments(videoId int64) ([]dto.CommentDto, error)
}

type CommentServiceImpl struct {
}

func (service *CommentServiceImpl) Comment(dto *dto.CommentDto) (int, error) {
	if dto == nil {
		return -1, errors.New("无评论信息")
	}
	if len(dto.Content) == 0 {
		return -1, errors.New("无评论信息")
	}
	var comment PO.Comment
	if err := dto.ToComment(&comment); err != nil {
		return -1, err
	}
	return commentDao.InsertComment(comment)
}
func (service *CommentServiceImpl) GetVideoComments(videoId int64) ([]dto.CommentDto, error) {
	comments, err := commentDao.SelectCommentByVideoId(videoId)
	if err != nil {
		return nil, err
	}
	commentsDto := make([]dto.CommentDto, len(comments))
	for i := 0; i < len(comments); i++ {
		var dto1 dto.CommentDto
		temp := comments[i]
		if err := copier.Copy(&dto1, temp); err != nil {
			return nil, err
		}
		commentsDto[i] = dto1
	}
	return commentsDto, nil

}
