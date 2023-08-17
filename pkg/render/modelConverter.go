// @description 数据库模型转换成交互模型
// @author zkp15
// @date 2023/8/11 11:17
// @version 1.0

package render

import (
	"github.com/sirupsen/logrus"
	"html"
	"time"
	"v-tiktok/model"
	"v-tiktok/model/comment"
	"v-tiktok/model/message"
	"v-tiktok/model/publish"
	"v-tiktok/model/relation"
	"v-tiktok/model/user"
	"v-tiktok/pkg/strs"
)

func UserConverter(userModel model.User, isFollow bool) user.User {
	return user.User{
		Avatar:          userModel.Avatar,
		BackgroundImage: userModel.BackgroundImage,
		FavoriteCount:   userModel.FavoriteCount,
		FollowCount:     userModel.FollowCount,
		FollowerCount:   userModel.FollowerCount,
		ID:              userModel.ID,
		IsFollow:        isFollow,
		Name:            userModel.Name,
		Signature:       userModel.Signature,
		TotalFavorited:  userModel.TotalFavorited,
		WorkCount:       userModel.WorkCount,
	}
}

func VideoConverter(videoModel model.Video, isFavorite bool, userModel model.User, isFollow bool) publish.Video {
	return publish.Video{
		Id:            videoModel.ID,
		Author:        UserConverter(userModel, isFollow),
		PlayUrl:       videoModel.PlayUrl,
		CoverUrl:      videoModel.CoverUrl,
		FavoriteCount: videoModel.FavoriteCount,
		CommentCount:  videoModel.CommentCount,
		IsFavorite:    isFavorite,
		Title:         videoModel.Title,
	}
}

func CommentConverter(commentModel model.Comment, userModel model.User, isFollow bool) comment.Comment {
	//xss预防
	newComment := html.EscapeString(commentModel.Content)
	if !strs.Equals(newComment, commentModel.Content) {
		logrus.Error("comment contains potential xss risk", commentModel.ID, commentModel.UserId)
	}

	return comment.Comment{
		Content:    newComment,
		CreateDate: commentModel.CreatedAt.Format("01-02"),
		ID:         commentModel.ID,
		User:       UserConverter(userModel, isFollow),
	}
}

func MessageConverter(messageMessage model.Message) message.Message {
	return message.Message{
		ID:         messageMessage.ID,
		Content:    messageMessage.Content,
		CreateTime: time.UnixMilli(messageMessage.CreatedAt).Format("2006-01-02 15:04:05"),
		FromUserID: messageMessage.FromUserID,
		ToUserID:   messageMessage.ToUserID,
	}
}

func FriendConverter(userModel model.User, isFollow bool, content string, msgType int64) relation.FriendUserList {
	return relation.FriendUserList{
		Avatar:          userModel.Avatar,
		BackgroundImage: userModel.BackgroundImage,
		FavoriteCount:   userModel.FavoriteCount,
		FollowCount:     userModel.FollowCount,
		FollowerCount:   userModel.FollowerCount,
		ID:              userModel.ID,
		IsFollow:        isFollow,
		Name:            userModel.Name,
		Signature:       userModel.Signature,
		TotalFavorited:  userModel.TotalFavorited,
		WorkCount:       userModel.WorkCount,
		Message:         content,
		MsgType:         msgType,
	}
}
