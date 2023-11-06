package error

import "github.com/gin-gonic/gin"

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

func (e MyError) AbortWithError(c *gin.Context, err error) *gin.Error {
	gEr := c.Error(err)
	c.Abort()
	return gEr
}
