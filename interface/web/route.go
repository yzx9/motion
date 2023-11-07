package web

import (
	"github.com/gin-gonic/gin"
	"github.com/yzx9/motion/interface/handler"
)

/**
*@Description:
*@Author: BZ
*@date: 2023/11/2 19:00
*@Version: V1.0
 */

func (s *Server) SetRouter() {
	r := s.GetGinEngine()
	r.GET("/index", func(c *gin.Context) {
		c.JSON(200, "This is a test api,pong pong shaKaLaKa!")
	})
	groupUser1 := r.Group("/api/v1/user/vail")
	{
		groupUser1.POST("/login", handler.UserLogin)
		groupUser1.POST("/register", handler.UserRegister)
	}
	//groupUser2 := r.Group("/api/v1/user", middleware.JwtAuthMiddleWare())
	groupUser2 := r.Group("/api/v1/user")
	{
		groupUser2.GET("/getUserInfo", handler.UploadAvatar)
		groupUser2.POST("/upload/avatar", handler.UploadAvatar)
		groupUser2.POST("/follow", handler.FollowHandler)
		groupUser2.POST("/like", handler.LikeHandler)
		groupUser2.POST("/comment", handler.CommentHandler)
		groupUser2.POST("/collect", handler.CollectHandler)
		groupUser2.GET("/fans")

	}

	groupVideo := r.Group("/api/v1/video")
	{
		groupVideo.GET("/recommend/videos", handler.GetVideosByRecommendHandler)
		groupVideo.GET("/getVideo/:id", handler.GetVideoByIdHandler)
		groupVideo.GET("/getComments/:videoId", handler.VideoCommentsHandler)
		groupVideo.POST("/uploadVideo", handler.PostVideoHandler)
	}

}
