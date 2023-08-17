// @description 点赞数据库操作
// @author zkp15
// @date 2023/8/11 16:20
// @version 1.0

package repository

import (
	"errors"
	"gorm.io/gorm"
	"v-tiktok/model"
)

func UserFavoriteExist(db *gorm.DB, userId, videoId int64) bool {
	var count int64
	db.Model(&model.Favorite{}).Where("user_id = ? AND video_id = ?", userId, videoId).Count(&count)
	return count > 0
}

func SaveUserFavorite(db *gorm.DB, userId, videoId int64) error {
	favorite := model.Favorite{
		UserId:  userId,
		VideoId: videoId,
	}
	result := db.Create(&favorite)
	if result.Error != nil {
		return errors.New("save user favorite error")
	}
	return nil
}

func DeleteUserFavorite(db *gorm.DB, userId, videoId int64) error {
	favorite := model.Favorite{
		UserId:  userId,
		VideoId: videoId,
	}
	result := db.Where("user_id = ? AND video_id = ?", favorite.UserId, favorite.VideoId).Delete(&favorite)
	if result.Error != nil {
		return errors.New("delete user favorite error")
	}
	return nil
}

func GetFavoriteList(db *gorm.DB, userId int64) ([]model.Favorite, error) {
	var videos []model.Favorite
	result := db.Where("user_id = ?", userId).Find(&videos)
	if result.Error != nil {
		return videos, errors.New("search user favorite error")
	}
	return videos, nil
}
