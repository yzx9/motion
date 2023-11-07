package handler

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	dto2 "github.com/yzx9/motion/command/domain/user/adpter/dto"
	"github.com/yzx9/motion/command/domain/user/service"
	error2 "github.com/yzx9/motion/command/infra/common/error"
	"github.com/yzx9/motion/command/infra/common/response"
	"github.com/yzx9/motion/command/infra/common/util"
	"net/http"
	"strconv"
)

/**
*@Description:
*@Author: BZ
*@date: 2023/11/2 19:13
*@Version: V1.0
 */
var userService service.UserServiceImpl
var followerService service.FollowerServiceImpl

func UserLogin(c *gin.Context) {
	var dto dto2.LoginDto
	if err := c.ShouldBind(&dto); err != nil {
		error2.GetMyError().AbortWithError(c, err)
		return
	}
	token, err := userService.Login(dto)
	if err != nil {
		error2.GetMyError().AbortWithError(c, err)
		return
	}
	c.Set("userId", token)
	c.JSON(http.StatusOK, token)
}

func UserRegister(c *gin.Context) {
	var dto dto2.LoginDto
	if err := c.ShouldBind(&dto); err != nil {
		error2.GetMyError().AbortWithError(c, err)
		return
	}
	err := userService.Register(dto)
	if err != nil {
		error2.GetMyError().AbortWithError(c, err)
		return
	}
	UserLogin(c)
}

func UploadAvatar(c *gin.Context) {
	avatar, err := c.FormFile("avatar")
	if err != nil {
		error2.GetMyError().AbortWithError(c, err)
		return
	}
	file, err := util.UploadFile(avatar, util.Avatar)
	id, exists := c.Get("userId")
	if !exists {
		response.ResponseFail(c, "gin can't get userId", nil)
	}
	userId, err := strconv.Atoi(fmt.Sprintf("%v", id))
	if err != nil {
		response.ResponseFail(c, err.Error(), nil)
	}
	if err = userService.UpdateAvatar(userId, file.Uuid, file.Url); err != nil {
		response.ResponseResult(c, http.StatusBadGateway, err.Error(), nil)
	}
	response.ResponseOk(c, "上传成功", file.Url)
}

func FollowHandler(c *gin.Context) {
	var dto dto2.FollowDto
	err := c.ShouldBind(&dto)
	if err != nil {
		error2.GetMyError().AbortWithError(c, err)
		return
	}
	value, exists := c.Get("userId")
	if !exists {
		error2.GetMyError().AbortWithError(c, errors.New("未登录"))
	}
	dto.UserId, err = strconv.Atoi(fmt.Sprintf("%v", value))
	err = followerService.Follow(&dto)
	if err != nil {
		error2.GetMyError().AbortWithError(c, err)
		return
	}
}

func getUserInfoHandler(c *gin.Context) {
	userId, err := GetContextUserId(c)
	if err != nil {
		response.ResponseResult(c, 401, "未登录", err.Error())
		return
	}
	dto, err := userService.GetUserInfo(userId)
	if err != nil {
		response.ResponseFail(c, "出错啦！", err.Error())
		return
	}
	response.ResponseOk(c, "user-info", dto)

}
