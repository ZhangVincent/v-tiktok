// @description 用户关系接口
// @author zkp15
// @date 2023/8/12 16:13

package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"v-tiktok/model/constants"
	"v-tiktok/model/relation"
	"v-tiktok/model/user"
	"v-tiktok/pkg/render"
	"v-tiktok/pkg/strs"
	"v-tiktok/pkg/token"
	"v-tiktok/service"
)

// RelationAction @description 关系操作
// @author zkp15
// @date 2023/8/16 18:34
func RelationAction(c *gin.Context) {
	//获取并校验数据
	var request relation.ActionRequest
	if err := c.ShouldBind(&request); err != nil || strs.IsAnyBlank(request.ActionType, request.Token) || request.ToUserID <= 0 {
		render.Error(c, constants.RelationErrorCode, "input params error")
		return
	}

	//解析token
	userClaims, err := token.ParseToken(request.Token)
	if err != nil {
		logrus.Error(err.Error())
		render.Error(c, constants.RelationErrorCode, "token parse error")
		return
	}

	//根据action type执行业务逻辑
	if request.ActionType == "1" {
		//关注
		if err = service.SaveRelation(userClaims.UserId, request.ToUserID); err != nil {
			render.Error(c, constants.RelationErrorCode, err.Error())
			return
		}
	} else if request.ActionType == "2" {
		//取消关注
		if err = service.DeleteRelation(userClaims.UserId, request.ToUserID); err != nil {
			render.Error(c, constants.RelationErrorCode, err.Error())
			return
		}
	} else {
		render.Error(c, constants.RelationErrorCode, "action type error")
		return
	}

	//封装返回
	render.Response(c, relation.ActionResponse{})
}

// FollowList @description 关注列表
// @author zkp15
// @date 2023/8/16 18:35
func FollowList(c *gin.Context) {
	//获取并校验数据
	var request relation.FollowListRequest
	if err := c.ShouldBind(&request); err != nil || strs.IsBlank(request.Token) || request.UserID <= 0 {
		render.Error(c, constants.RelationFollowErrorCode, "input params error")
		return
	}

	//解析token
	userClaims, err := token.ParseToken(request.Token)
	if err != nil {
		logrus.Error("token parse error")
		render.Error(c, constants.RelationFollowErrorCode, "token parse error")
		return
	}

	//查询用户关注列表
	followUsers, err := service.GetFollowList(request.UserID)
	if err != nil {
		render.Error(c, constants.RelationFollowErrorCode, err.Error())
		return
	}

	//封装并返回
	n := len(followUsers)
	followUserList := make([]user.User, n)
	for i, u := range followUsers {
		followUserList[i] = render.UserConverter(u, service.CheckIsFollow(userClaims.UserId, u.ID))
	}

	render.Response(c, relation.FollowListResponse{
		UserList: followUserList,
	})
}

// FollowerList @description 粉丝列表
// @author zkp15
// @date 2023/8/16 18:38
func FollowerList(c *gin.Context) {
	//获取并校验数据
	var request relation.FollowerListRequest
	if err := c.ShouldBind(&request); err != nil || strs.IsBlank(request.Token) || request.UserID <= 0 {
		render.Error(c, constants.RelationFollowerErrorCode, "input params error")
		return
	}

	//解析token
	userClaims, err := token.ParseToken(request.Token)
	if err != nil {
		logrus.Error("token parse error")
		render.Error(c, constants.RelationFollowerErrorCode, "token parse error")
		return
	}

	//查询用户粉丝列表
	followerUsers, err := service.GetFollowerList(request.UserID)
	if err != nil {
		render.Error(c, constants.RelationFollowerErrorCode, err.Error())
		return
	}

	//封装并返回
	n := len(followerUsers)
	followerUserList := make([]user.User, n)
	for i, u := range followerUsers {
		followerUserList[i] = render.UserConverter(u, service.CheckIsFollow(userClaims.UserId, u.ID))
	}

	render.Response(c, relation.FollowerListResponse{
		UserList: followerUserList,
	})
}

// FriendList @description 好友列表
// @author zkp15
// @date 2023/8/16 18:40
func FriendList(c *gin.Context) {
	//获取并校验数据
	var request relation.FriendListRequest
	if err := c.ShouldBind(&request); err != nil || strs.IsBlank(request.Token) || request.UserID <= 0 {
		render.Error(c, constants.RelationFriendErrorCode, "input params error")
		return
	}

	//解析token
	_, err := token.ParseToken(request.Token)
	if err != nil {
		logrus.Error("token parse error")
		render.Error(c, constants.RelationFriendErrorCode, "token parse error")
		return
	}

	//查询用户好友列表及其最新消息
	FriendUsers, err := service.GetFriendList(request.UserID)
	if err != nil {
		render.Error(c, constants.RelationFriendErrorCode, err.Error())
		return
	}

	//封装并返回
	render.Response(c, relation.FriendListResponse{
		UserList: FriendUsers,
	})
}
