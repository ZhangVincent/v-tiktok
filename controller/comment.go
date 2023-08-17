package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"v-tiktok/model"
	"v-tiktok/model/comment"
	"v-tiktok/model/constants"
	"v-tiktok/pkg/render"
	"v-tiktok/pkg/strs"
	"v-tiktok/pkg/token"
	"v-tiktok/service"
)

// CommentAction @description 评论操作
// @author zkp15
// @date 2023/8/16 15:41
func CommentAction(c *gin.Context) {
	// 获取并校验参数
	var commentRequest comment.ActionRequest
	if err := c.ShouldBind(&commentRequest); err != nil || strs.IsAnyBlank(commentRequest.Token, commentRequest.ActionType) {
		render.Error(c, constants.CommentErrorCode, "input params error")
		return
	}

	//解析token
	userClaims, err := token.ParseToken(commentRequest.Token)
	if err != nil {
		logrus.Error(err.Error())
		render.Error(c, constants.FavoriteErrorCode, "token parse error")
		return
	}

	var commentInfo model.Comment
	//数据库操作
	if commentRequest.ActionType == "1" {
		if commentInfo, err = service.SaveComment(commentRequest.VideoID, commentRequest.CommentText, userClaims.UserId); err != nil {
			render.Error(c, constants.FavoriteErrorCode, err.Error())
			return
		}
	} else if commentRequest.ActionType == "2" {
		if err := service.DeleteComment(commentRequest.CommentID); err != nil {
			render.Error(c, constants.FavoriteErrorCode, err.Error())
			return
		}
	} else {
		render.Error(c, constants.FavoriteErrorCode, "action type error")
		return
	}

	author, err := service.GetUserById(commentInfo.UserId)
	if err != nil {
		author, err = service.GetUserById(userClaims.UserId)
		if err != nil {
			author = constants.DefaultUserModel
		}
	}

	//封装并返回数据
	render.Response(c, comment.ActionResponse{
		Comment: render.CommentConverter(commentInfo, author, service.CheckIsFollow(userClaims.UserId, author.ID)),
	})
}

// CommentList @description 视频的评论列表
// @author zkp15
// @date 2023/8/16 18:31
func CommentList(c *gin.Context) {
	// 获取并校验参数
	var commentRequest comment.ListRequest
	if err := c.ShouldBind(&commentRequest); err != nil || commentRequest.VideoID <= 0 {
		render.Error(c, constants.CommentErrorCode, "input params error")
		return
	}

	//解析token
	//userClaims, err := token.ParseToken(commentRequest.Token)
	//if err != nil {
	//	logrus.Error(err.Error())
	//	render.Error(c, constants.FavoriteErrorCode, "token parse error")
	//	return
	//}

	//查询评论
	comments, err := service.GetVideoCommentList(commentRequest.VideoID)
	if err != nil {
		render.Error(c, constants.FavoriteErrorCode, err.Error())
		return
	}

	//数据格式封装
	commentList := make([]comment.Comment, len(comments))
	for i, c := range comments {
		author, err := service.GetUserById(c.UserId)
		if err != nil {
			author = constants.DefaultUserModel
		}
		//commentList[i] = render.CommentConverter(c, author, service.CheckIsFollow(userClaims.UserId, author.ID))
		commentList[i] = render.CommentConverter(c, author, false)
	}

	//封装并返回数据
	render.Response(c, comment.ListResponse{
		CommentList: commentList,
	})
}
