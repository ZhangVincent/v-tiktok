// @description 视频发布查询接口
// @author zkp15
// @date 2023/8/11 10:10

package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"v-tiktok/model/config"
	"v-tiktok/model/constants"
	"v-tiktok/model/publish"
	"v-tiktok/pkg/render"
	"v-tiktok/pkg/strs"
	"v-tiktok/pkg/token"
	"v-tiktok/service"
)

// PublishAction @description 视频投稿
// @author zkp15
// @date 2023/8/16 15:18
func PublishAction(c *gin.Context) {
	//获取并校验参数
	var request publish.ActionRequest
	if err := c.ShouldBind(&request); err != nil || strs.IsBlank(request.Token) || request.Data == nil {
		render.Error(c, constants.PublishErrorCode, "video publish params error")
		return
	}

	//解析token
	userClaims, err := token.ParseToken(request.Token)
	if err != nil {
		logrus.Error("token parse error", request.Token)
		render.Error(c, constants.PublishErrorCode, "token parse error")
		return
	}

	//根据设置选择不同的保存方式
	if config.Instance.Uploader.Enable == "minio" {
		//保存文件系统
		if err := service.SaveVideoOnMinio(request.Data, request.Title, userClaims.UserId); err != nil {
			render.Error(c, constants.PublishErrorCode, err.Error())
			return
		}
	} else if config.Instance.Uploader.Enable == "local" {
		//保存本地
		if err := service.SaveVideoOnLocal(request.Data, request.Title, userClaims.UserId); err != nil {
			render.Error(c, constants.PublishErrorCode, err.Error())
			return
		}
	} else {
		logrus.Error("uploader config error")
		return
	}

	//封装并返回数据
	render.Response(c, publish.ActionResponse{})
}

// PublishList @description 发布列表
// @author zkp15
// @date 2023/8/16 15:20
func PublishList(c *gin.Context) {
	//获取并校验参数
	var request publish.ListRequest
	if err := c.ShouldBind(&request); err != nil || strs.IsBlank(request.Token) || request.UserID <= 0 {
		render.Error(c, constants.PublishErrorCode, "video publish params error")
		return
	}

	//解析token
	userClaims, err := token.ParseToken(request.Token)
	if err != nil {
		logrus.Error("token parse error")
		render.Error(c, constants.PublishErrorCode, "token parse error")
		return
	}

	//查询用户发表过的视频列表
	videos, err := service.GetVideosByUserId(request.UserID)
	if err != nil {
		render.Error(c, constants.PublishErrorCode, err.Error())
		return
	}

	//查询发表视频的人
	author, err := service.GetUserById(request.UserID)
	if err != nil {
		author = constants.DefaultUserModel
	}
	isFollow := service.CheckIsFollow(userClaims.UserId, request.UserID)

	//封装返货参数
	videoList := make([]publish.Video, len(videos))
	for i, v := range videos {
		videoList[i] = render.VideoConverter(v, service.CheckIsFavorite(userClaims.UserId, v.ID), author, isFollow)
	}

	render.Response(c, publish.ListResponse{
		VideoList: videoList,
	})
}
