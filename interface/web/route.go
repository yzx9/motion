package web

import (
	"github.com/yzx9/motion/interface/handler"
	middleware "github.com/yzx9/motion/interface/middleWare"
)

/**
*@Description:
*@Author: BZ
*@date: 2023/11/2 19:00
*@Version: V1.0
 */

func (s *Server) SetRouter() {
	r := s.GetGinEngine()
	r.GET("/index")
	groupUser1 := r.Group("/api/v1/user/vail")
	{
		groupUser1.POST("/login", handler.UserLogin)
		groupUser1.POST("/register", handler.UserRegister)
	}
	groupUser2 := r.Group("/api/v1/user", middleware.JwtAuthMiddleWare())
	{
		groupUser2.GET("/personal-home")
		groupUser2.POST("/upload/avatar", handler.UploadAvatar)
		groupUser2.POST("/follow", handler.FollowHandler)
		groupUser2.POST("/like", handler.LikeHandler)
		groupUser2.GET("/fans")
	}

	groupVideo := r.Group("/api/v1/video", middleware.JwtAuthMiddleWare())
	{
		groupVideo.GET("/recommend/videos", handler.GetVideosByRecommend)
		groupVideo.GET("/getVideo/:id", handler.GetVideoByIdHandler)
		groupVideo.GET("/getComments/:video-id", handler.VideoCommentsHandler)
		groupVideo.POST("/upload-video", handler.PostVideoHandler)
	}
}
