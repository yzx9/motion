package middleware

import (
	"github.com/gin-gonic/gin"
	error2 "github.com/yzx9/motion/command/infra/common/error"
	"github.com/yzx9/motion/command/infra/common/util"
	"net/http"
)

const authorizationHead = "AuthorizationKey"

func JwtAuthMiddleWare() func(ctx *gin.Context) {
	return func(c *gin.Context) {
		token := c.Request.Header.Get(authorizationHead)
		claims, err := util.ParseToken(token)
		if err != nil {
			error2.GetMyError().AbortWithError(c, err)
			c.Redirect(http.StatusUnauthorized, "/api/v1/user/login")
			return
		}
		c.Set("userId", claims.UserID)
		c.Next()
	}
}
