// @description 消息数据库操作
// @author zkp15
// @date 2023/8/12 22:13
// @version 1.0

package repository

import (
	"gorm.io/gorm"
	"time"
	"v-tiktok/model"
)

func GetLatestMessages(db *gorm.DB, fromUserId, toUserId int64) (model.Message, error) {
	var message model.Message
	result := db.Where("from_user_id = ? AND to_user_id = ?", fromUserId, toUserId).Last(&message)
	if result.Error != nil {
		return message, result.Error
	}
	return message, nil
}

func SaveMessage(db *gorm.DB, fromUserId, toUserId int64, content string) error {
	result := db.Create(&model.Message{
		GormTimeModel: model.GormTimeModel{
			CreatedAt: time.Now().UnixMilli(),
			UpdatedAt: time.Now().UnixMilli(),
		},
		Content:    content,
		FromUserID: fromUserId,
		ToUserID:   toUserId,
	})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func GetMessage(db *gorm.DB, fromUserId, toUserId, preMsgTime int64) ([]model.Message, error) {
	var messages []model.Message
	result := db.Where("from_user_id = ? AND to_user_id = ? AND created_at > ?", fromUserId, toUserId, preMsgTime).Or("from_user_id = ? AND to_user_id = ? AND created_at > ?", toUserId, fromUserId, preMsgTime).Find(&messages)
	if result.Error != nil {
		return nil, result.Error
	}
	return messages, nil
}
