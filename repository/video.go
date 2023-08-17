// @description 视频查询
// @author zkp15
// @date 2023/8/11 17:02
// @version 1.0

package repository

import (
	"errors"
	"gorm.io/gorm"
	"time"
	"v-tiktok/model"
)

func VideoExist(db *gorm.DB, videoId int64) bool {
	var count int64
	db.Model(&model.Video{}).Where("id = ?", videoId).Count(&count)
	return count > 0
}

func GetVideo(db *gorm.DB, videoId int64) (model.Video, error) {
	var video model.Video
	result := db.First(&video, videoId)
	if result.Error != nil {
		return video, result.Error
	}
	return video, nil
}

func GetVideosById(db *gorm.DB, videoIds []int64) ([]model.Video, error) {
	var videos []model.Video
	result := db.Order("created_at desc").Find(&videos, videoIds)
	if result.Error != nil {
		return nil, result.Error
	}
	return videos, nil
}

func GetVideosByTime(db *gorm.DB, latestTime int64, videosNum int) ([]model.Video, error) {
	var videos []model.Video
	result := db.Where("created_at > ? AND is_audit = 0", latestTime).Limit(videosNum).Order("created_at desc").Find(&videos)
	if result.Error != nil {
		return nil, result.Error
	}
	return videos, nil
}

func GetVideosByUserId(db *gorm.DB, userId int64) ([]model.Video, error) {
	var videos []model.Video
	result := db.Where("author_id = ? AND is_audit = 0", userId).Order("created_at desc").Find(&videos)
	if result.Error != nil {
		return videos, result.Error
	}
	return videos, nil
}

func SaveVideo(db *gorm.DB, authorId int64, playUrl, coverUrl, title string, isAudit int64) error {
	result := db.Create(&model.Video{
		GormTimeModel: model.GormTimeModel{
			CreatedAt: time.Now().UnixMilli(),
			UpdatedAt: time.Now().UnixMilli(),
		},
		AuthorId: authorId,
		PlayUrl:  playUrl,
		CoverUrl: coverUrl,
		Title:    title,
		IsAudit:  isAudit,
	})
	if result.Error != nil || result.RowsAffected <= 0 {
		return errors.New("save video error")
	}
	return nil
}

func UpdateVideoFavorited(db *gorm.DB, videoId int64, changeScore int64) error {
	return db.Exec("update t_video set favorite_count = favorite_count + ? where id = ?", changeScore, videoId).Error
}

func UpdateVideoComment(db *gorm.DB, videoId int64, changeScore int64) error {
	return db.Exec("update t_video set comment_count = comment_count + ? where id = ?", changeScore, videoId).Error
}
