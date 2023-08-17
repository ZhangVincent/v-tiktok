// @description 点赞喜欢
// @author zkp15
// @date 2023/8/12 9:37

package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"v-tiktok/model/constants"
	"v-tiktok/model/favorite"
	"v-tiktok/model/publish"
	"v-tiktok/pkg/render"
	"v-tiktok/pkg/strs"
	"v-tiktok/pkg/token"
	"v-tiktok/service"
)

// FavoriteAction @description 对视频的赞操作
// @author zkp15
// @date 2023/8/16 15:26
func FavoriteAction(c *gin.Context) {
	//获取并校验参数
	var actionRequest favorite.ActionRequest
	if err := c.ShouldBind(&actionRequest); err != nil || strs.IsAnyBlank(actionRequest.ActionType, actionRequest.Token) || actionRequest.VideoID <= 0 {
		render.Error(c, constants.FavoriteErrorCode, "input params error")
		return
	}

	//解析token
	userClaims, err := token.ParseToken(actionRequest.Token)
	if err != nil {
		logrus.Error(err.Error())
		render.Error(c, constants.FavoriteErrorCode, "token parse error")
		return
	}

	//向用户视频表中插入一条数据
	if err = service.SaveOrDeleteUserFavorite(actionRequest.ActionType, userClaims.UserId, actionRequest.VideoID); err != nil {
		render.Error(c, constants.FavoriteErrorCode, err.Error())
		return
	}

	//封装并返回
	render.Response(c, favorite.ActionResponse{})
}

// FavoriteList @description 查询用户喜欢的视频列表
// @author zkp15
// @date 2023/8/16 15:37
func FavoriteList(c *gin.Context) {
	//获取并校验参数
	var listRequest favorite.ListRequest
	if err := c.ShouldBind(&listRequest); err != nil || strs.IsBlank(listRequest.Token) || listRequest.UserID <= 0 {
		render.Error(c, constants.FavoriteErrorCode, "input params error")
		return
	}

	//解析token
	userClaims, err := token.ParseToken(listRequest.Token)
	if err != nil {
		logrus.Error(err.Error())
		render.Error(c, constants.FavoriteErrorCode, "token parse error")
		return
	}

	//根据userid查询喜欢列表
	videos, err := service.GetFavoriteList(listRequest.UserID)
	if err != nil {
		render.Error(c, constants.FavoriteErrorCode, err.Error())
		return
	}

	//封装返回
	n := len(videos)
	videoList := make([]publish.Video, n)
	for i, v := range videos {
		author, err := service.GetUserById(v.AuthorId)
		if err != nil {
			author = constants.DefaultUserModel
		}
		videoList[i] = render.VideoConverter(v, service.CheckIsFollow(userClaims.UserId, author.ID), author, service.CheckIsFollow(userClaims.UserId, v.AuthorId))
	}

	render.Response(c, favorite.ListResponse{
		VideoList: videoList,
	})
}
