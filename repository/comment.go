// @description 评论数据库操作
// @author zkp15
// @date 2023/8/12 10:05
// @version 1.0

package repository

import (
	"errors"
	"gorm.io/gorm"
	"v-tiktok/model"
)

func SaveComment(db *gorm.DB, commentInfo *model.Comment) error {
	result := db.Create(&commentInfo)
	if result.Error != nil {
		return errors.New("save comment error")
	}
	return nil
}

func GetCommentList(db *gorm.DB, videoId int64) ([]model.Comment, error) {
	var commnts []model.Comment
	result := db.Where("video_id = ?", videoId).Find(&commnts)
	if result.Error != nil {
		return nil, errors.New("query comments error")
	}
	return commnts, nil
}

func GetComment(db *gorm.DB, commentId int64) (model.Comment, error) {
	var commentInfo model.Comment
	result := db.Find(&commentInfo, commentId)
	if result.Error != nil || result.RowsAffected <= 0 {
		return model.Comment{}, errors.New("query comment error")
	}
	return commentInfo, nil
}

func DeleteComment(db *gorm.DB, commentId int64) error {
	result := db.Delete(&model.Comment{}, commentId)
	if result.Error != nil || result.RowsAffected <= 0 {
		return errors.New("delete comment error")
	}
	return nil
}
