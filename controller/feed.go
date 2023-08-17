// @description 视频流接口
// @author zkp15
// @date 2023/8/15 10:22

package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"v-tiktok/model/constants"
	"v-tiktok/model/feed"
	"v-tiktok/model/publish"
	"v-tiktok/pkg/render"
	"v-tiktok/pkg/strs"
	"v-tiktok/pkg/token"
	"v-tiktok/service"
)

// Feed @description 视频流接口
// @author zkp15
// @date 2023/8/16 15:01
func Feed(c *gin.Context) {
	//获取并校验参数
	var request feed.FeedRequest
	if err := c.ShouldBind(&request); err != nil {
		logrus.Error("input params bind error", err.Error())
	}

	//如果是登录用户就记录id
	var userId int64 = 0
	if strs.IsBlank(request.Token) {
		logrus.Info("未登录用户访问视频")
	} else {
		//解析token
		userClaims, err := token.ParseToken(request.Token)
		if err != nil {
			logrus.Error("token parse error")
		}
		userId = userClaims.UserId
	}

	//查询视频列表
	videos, err := service.GetVideos(request.LatestTime)
	if err != nil {
		render.Error(c, constants.FeedErrorCode, err.Error())
		return
	}

	//封装返货参数
	videoList := make([]publish.Video, len(videos))
	for i, v := range videos {
		author, err := service.GetUserById(v.AuthorId)
		if err != nil {
			author = constants.DefaultUserModel
		}
		videoList[i] = render.VideoConverter(v, service.CheckIsFavorite(userId, v.ID), author, service.CheckIsFollow(userId, v.AuthorId))
	}

	var nextTime int64 = 0
	if len(videos) > 0 {
		nextTime = videos[len(videos)-1].CreatedAt
	}

	render.Response(c, feed.FeedResponse{
		VideoList: videoList,
		NextTime:  nextTime,
	})
}
