// @description 关注表数据库操作
// @author zkp15
// @date 2023/8/12 22:06
// @version 1.0

package repository

import (
	"errors"
	"gorm.io/gorm"
	"v-tiktok/model"
)

func RelationExist(db *gorm.DB, fromUserId, toUserId int64) bool {
	var count int64
	db.Model(&model.Follow{}).Where("from_user_id = ? AND to_user_id = ?", fromUserId, toUserId).Count(&count)
	return count > 0
}

func GetFollowsByFromUserId(db *gorm.DB, fromUserId int64) ([]model.Follow, error) {
	var follows []model.Follow
	result := db.Where("from_user_id = ?", fromUserId).Find(&follows)
	if result.Error != nil {
		return follows, result.Error
	}
	return follows, nil
}

func GetFollowersByToUserId(db *gorm.DB, toUserId int64) ([]model.Follow, error) {
	var follows []model.Follow
	result := db.Where("to_user_id = ?", toUserId).Find(&follows)
	if result.Error != nil {
		return follows, result.Error
	}
	return follows, nil
}

func SaveRelation(db *gorm.DB, fromUserId, toUserId int64) error {
	result := db.Create(&model.Follow{
		FromUserId: fromUserId,
		ToUserId:   toUserId,
	})
	if result.Error != nil || result.RowsAffected <= 0 {
		return errors.New("save relation error")
	}
	return nil
}

func DeleteRelation(db *gorm.DB, fromUserId, toUserId int64) error {
	result := db.Delete(model.Follow{}, "from_user_id = ? AND to_user_id = ?", fromUserId, toUserId)
	if result.Error != nil || result.RowsAffected <= 0 {
		return errors.New("delete relation error")
	}
	return nil
}
