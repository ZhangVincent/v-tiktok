package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"v-tiktok/controller"
	"v-tiktok/model/config"
)

func Router() {
	r := gin.Default()

	r.ForwardedByClientIP = true
	r.SetTrustedProxies([]string{config.Instance.BaseUrl})

	initRouter(r)

	r.Run(config.Instance.BaseUrl)
}

func initRouter(r *gin.Engine) {
	// public directory is used to serve static resources
	r.Static("/static", "./public")

	r.Any("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "powered by tiktok-zkp")
	})

	// 抖音业务模块
	apiRouter := r.Group("/douyin")

	// 视频流
	apiRouter.GET("/feed/", controller.Feed)

	// 用户
	userGroup := apiRouter.Group("/user/")
	{
		userGroup.GET("/", controller.UserInfo)
		userGroup.POST("/register/", controller.Register)
		userGroup.POST("/login/", controller.Login)
	}

	// 发布
	publishGroup := apiRouter.Group("/publish/")
	{
		publishGroup.POST("/action/", controller.PublishAction)
		publishGroup.GET("/list/", controller.PublishList)
	}

	// 喜欢
	favoriteGroup := apiRouter.Group("/favorite/")
	{
		favoriteGroup.POST("/action/", controller.FavoriteAction)
		favoriteGroup.GET("/list/", controller.FavoriteList)
	}

	// 评论
	commentGroup := apiRouter.Group("/comment/")
	{
		commentGroup.POST("/action/", controller.CommentAction)
		commentGroup.GET("/list/", controller.CommentList)
	}

	// 关注/粉丝
	relationGroup := apiRouter.Group("/relation/")
	{
		relationGroup.POST("/action/", controller.RelationAction)
		relationGroup.GET("/follow/list/", controller.FollowList)
		relationGroup.GET("/follower/list/", controller.FollowerList)
		relationGroup.GET("/friend/list/", controller.FriendList)
	}

	// 消息
	messageGroup := apiRouter.Group("/message/")
	{
		messageGroup.GET("/chat/", controller.MessageChat)
		messageGroup.POST("/action/", controller.MessageAction)
	}
}
