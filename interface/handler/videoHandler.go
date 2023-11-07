package handler

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/yzx9/motion/command/domain/video/adapter/dto"
	"github.com/yzx9/motion/command/domain/video/service"
	error2 "github.com/yzx9/motion/command/infra/common/error"
	"github.com/yzx9/motion/command/infra/common/response"
	"github.com/yzx9/motion/command/infra/common/util"
	"github.com/yzx9/motion/command/infra/dao/PO"
	"net/http"
	"strconv"
)

/**
*@Description:
*@Author: BZ
*@date: 2023/11/3 13:58
*@Version: V1.0
 */
var videoService service.VideoService
var likeService service.LikeService
var commentService service.CommentService

const (
	dirVideo = "videos"
)

func PostVideoHandler(c *gin.Context) {
	fileHeader, err := c.FormFile(dirVideo)
	if err != nil {
		response.ResponseResult(c, http.StatusBadGateway, "上传失败", nil)
		return
	}
	ret, err := util.UploadFile(fileHeader, dirVideo)
	if err != nil {
		response.ResponseResult(c, http.StatusBadGateway, "上传失败", nil)
		return
	}
	var videoDto dto.VideoDto
	if err = c.ShouldBind(&videoDto); err != nil {
		error2.GetMyError().AbortWithError(c, err)
		return
	}
	videoDto.Id = ret.Uuid
	videoDto.Url = ret.Url
	videoDto.Cover = ret.CoverURL
	var video PO.Video
	videoDto.ToVideo(&video)
	video.UserId, err = GetContextUserId(c)
	if err = videoService.SaveVideo(video); err != nil {
		error2.GetMyError().AbortWithError(c, err)
		return
	}
	response.ResponseOk(c, "上传成功", videoDto.Id)
}

func GetVideoByIdHandler(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		error2.GetMyError().AbortWithError(c, err)
		return
	}
	userId, err := GetContextUserId(c)
	if err != nil {
		error2.GetMyError().AbortWithError(c, err)
		return
	}
	video, err := videoService.SelectVideoByID(id, userId)
	if err != nil {
		error2.GetMyError().AbortWithError(c, err)
		return
	}
	response.ResponseOk(c, "recommend videos", video)
}

func LikeHandler(c *gin.Context) {
	var likeDto dto.LikeDto
	err := c.ShouldBind(&likeDto)
	if err != nil {
		error2.GetMyError().AbortWithError(c, err)
		return
	}
	value, exists := c.Get("userId")
	if !exists {
		error2.GetMyError().AbortWithError(c, errors.New("未登录"))
	}
	likeDto.UserId, err = strconv.Atoi(fmt.Sprintf("%v", value))
	err = likeService.Like(&likeDto)
	if err != nil {
		error2.GetMyError().AbortWithError(c, err)
		return
	}
}

func CommentHandler(c *gin.Context) {
	var commentDto dto.CommentDto
	err := c.ShouldBind(&commentDto)
	if err != nil {
		error2.GetMyError().AbortWithError(c, err)
		return
	}
	value, exists := c.Get("userId")
	if !exists {
		error2.GetMyError().AbortWithError(c, errors.New("未登录"))
	}
	commentDto.UserId, err = strconv.Atoi(fmt.Sprintf("%v", value))
	id, err := commentService.Comment(&commentDto)
	if err != nil {
		error2.GetMyError().AbortWithError(c, err)
		return
	}
	response.ResponseOk(c, "评论成功,返回评论的id", id)

}

func VideoCommentsHandler(c *gin.Context) {
	videoId := c.Param("video-id")
	id, err := strconv.ParseInt(videoId, 10, 64)
	commentDtos, err := commentService.GetVideoComments(id)
	if err != nil {
		error2.GetMyError().AbortWithError(c, err)
		return
	}
	response.ResponseOk(c, "视频的评论列表", commentDtos)

}

func GetVideosByRecommend(c *gin.Context) {

}

func GetContextUserId(c *gin.Context) (int, error) {
	value, exists := c.Get("userId")
	if !exists {
		error2.GetMyError().AbortWithError(c, errors.New("未登录"))
	}
	return strconv.Atoi(fmt.Sprintf("%v", value))
}
