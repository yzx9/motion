package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/**
*@Description:
*@Author: BZ
*@date: 2023/11/4 22:05
*@Version: V1.0
 */

func ResponseResult(c *gin.Context, code int, msg string, data any) {
	c.JSON(code, map[string]interface{}{
		"code": code,
		"msg":  msg,
		"data": data,
	})
}

func ResponseOk(c *gin.Context, msg string, data any) {
	ResponseResult(c, http.StatusOK, msg, data)
}

func ResponseFail(c *gin.Context, msg string, data any) {
	ResponseResult(c, http.StatusBadRequest, msg, data)
}
