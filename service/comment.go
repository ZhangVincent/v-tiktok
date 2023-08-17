// @description 评论业务
// @author zkp15
// @date 2023/8/12 9:27
// @version 1.0

package service

import (
	"errors"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"v-tiktok/model"
	"v-tiktok/pkg/forbiddenWordCheck"
	"v-tiktok/pkg/sqls"
	"v-tiktok/pkg/strs"
	"v-tiktok/repository"
)

func SaveComment(videoId int64, commentText string, userId int64) (model.Comment, error) {
	//数据校验
	if videoId <= 0 || strs.IsBlank(commentText) || userId <= 0 {
		return model.Comment{}, errors.New("input params error")
	}

	//内容审核
	hitWords := forbiddenWordCheck.Check(&commentText)
	if len(hitWords) != 0 {
		logrus.Errorf("user (id: %v) comment (%v) contains forbidden words (%v)", userId, commentText, hitWords)
	}

	//数据封装
	commentInfo := model.Comment{
		GormModel: model.GormModel{},
		Content:   commentText,
		UserId:    userId,
		VideoId:   videoId,
	}

	//数据库操作
	if err := sqls.DB().Transaction(func(tx *gorm.DB) error {
		//保存评论到数据库
		if err := repository.SaveComment(tx, &commentInfo); err != nil {
			return err
		}
		//video的评论数+1
		if err := repository.UpdateVideoComment(tx, videoId, 1); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return model.Comment{}, err
	}

	return commentInfo, nil
}

func DeleteComment(commentId int64) error {
	//数据校验
	if commentId <= 0 {
		return errors.New("input params error")
	}

	db := sqls.DB()
	//查询是否存在评论
	var commentInfo model.Comment
	var err error
	if commentInfo, err = repository.GetComment(db, commentId); err != nil {
		return err
	}

	//数据库操作
	if err := db.Transaction(func(tx *gorm.DB) error {
		//删除
		if err := repository.DeleteComment(tx, commentId); err != nil {
			return err
		}
		//video的评论数-1
		if err := repository.UpdateVideoComment(tx, commentInfo.VideoId, -1); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return err
	}

	return nil
}

func GetVideoCommentList(videoId int64) ([]model.Comment, error) {
	//数据校验
	if videoId <= 0 {
		return nil, errors.New("input params error")
	}

	//数据库操作
	comments, err := repository.GetCommentList(sqls.DB(), videoId)
	if err != nil {
		return nil, err
	}

	return comments, nil
}
