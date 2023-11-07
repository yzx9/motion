# 议题介绍
网页短视频应用
使用七牛云存储、七牛视频相关产品（如视频截帧等）开发一款Web端短视频应用

基础功能（必须实现）
视频播放：播放、暂停、进度条拖拽
内容分类：视频内容分类页，如热门视频、体育频道
­视频切换：可通过上下键翻看视频
高级功能（可选实现）
账户系统：用户可登录，收藏视频
可参考常见短视频应用自由增加功能，提升完善度，如点赞、分享、关注、搜索等


# Motion

Create, Share, Move with Motion.

## Deployment

```
bash deployment/build.sh
docker compose -f deployment/docker-compose.yaml up
```

## Development

Run server:

```
go run .
```

Run client:

```
pnpm install \
&& pnpm run dev:client 
```


api:

```
/api/v1/user/vail/login 登录
/api/v1/user/vail/login 注册
/api/v1/user/getUserInfo 用户主页信息
/api/v1/user/avatar上传头像


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

```