package error

import (
	"github.com/gin-gonic/gin"
	"github.com/yzx9/motion/command/infra/common/response"
)

/**
*@Description:
*@Author: BZ
*@date: 2023/11/2 19:24
*@Version: V1.0
 */

type MyError struct {
}

var myError = new(MyError)

func GetMyError() *MyError {
	return myError
}

func (e MyError) AbortWithError(c *gin.Context, err error) {
	gEr := c.Error(err)
	c.Abort()
	response.ResponseFail(c, "发生错误", gEr.Error())
}
