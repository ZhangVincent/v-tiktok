// @description 视频点赞业务
// @author zkp15
// @date 2023/8/11 16:09
// @version 1.0

package service

import (
	"errors"
	"gorm.io/gorm"
	"v-tiktok/model"
	"v-tiktok/pkg/sqls"
	"v-tiktok/repository"
)

func SaveOrDeleteUserFavorite(actionType string, userId int64, videoId int64) error {
	// 校验参数
	if userId <= 0 || !repository.UserExist(sqls.DB(), userId) {
		return errors.New("user not found")
	}
	if videoId <= 0 || !repository.VideoExist(sqls.DB(), videoId) {
		return errors.New("video not found")
	}

	// 查询视频信息
	video, err := repository.GetVideo(sqls.DB(), videoId)
	if err != nil {
		return errors.New("query video error")
	}

	// 根据动作类型执行业务
	if actionType == "1" {
		//点赞
		db := sqls.DB()
		// 查询是否存在
		if ifFavorite := repository.UserFavoriteExist(db, userId, videoId); ifFavorite {
			return errors.New("favorite already exit")
		}
		if err := db.Transaction(func(tx *gorm.DB) error {
			// 点赞，插入数据库
			if err := repository.SaveUserFavorite(tx, userId, videoId); err != nil {
				return err
			}
			// 用户喜欢数+1
			if err := repository.UpdateUserFavorite(tx, userId, 1); err != nil {
				return err
			}
			// 视频作者获赞数+1
			if err := repository.UpdateUserFavorited(tx, video.AuthorId, 1); err != nil {
				return err
			}
			// 视频获赞数+1
			if err := repository.UpdateVideoFavorited(tx, videoId, 1); err != nil {
				return err
			}
			return nil
		}); err != nil {
			return err
		}
	} else if actionType == "2" {
		//取消赞
		db := sqls.DB()
		// 查询是否不存在
		if ifFavorite := repository.UserFavoriteExist(db, userId, videoId); !ifFavorite {
			return errors.New("favorite already redo")
		}
		if err := db.Transaction(func(tx *gorm.DB) error {
			// 取消点赞，直接逻辑删除
			if err := repository.DeleteUserFavorite(tx, userId, videoId); err != nil {
				return err
			}
			// 用户喜欢数-1
			if err := repository.UpdateUserFavorite(tx, userId, -1); err != nil {
				return err
			}
			// 视频作者获赞数-1
			if err := repository.UpdateUserFavorited(tx, video.AuthorId, -1); err != nil {
				return err
			}
			// 视频获赞数-1
			if err := repository.UpdateVideoFavorited(tx, videoId, -1); err != nil {
				return err
			}
			return nil
		}); err != nil {
			return err
		}
	} else {
		return errors.New("action type error")
	}

	return nil
}

func GetFavoriteList(userId int64) ([]model.Video, error) {
	// 查询数据库
	favoriteList, err := repository.GetFavoriteList(sqls.DB(), userId)
	if err != nil {
		return nil, err
	}

	// 根据videoId查询视频列表
	n := len(favoriteList)
	if n == 0 {
		return []model.Video{}, nil
	}
	videoIds := make([]int64, n)
	for i, favorite := range favoriteList {
		videoIds[i] = favorite.VideoId
	}

	videoList, err := repository.GetVideosById(sqls.DB(), videoIds)
	if err != nil {
		return nil, err
	}

	return videoList, nil
}

func CheckIsFavorite(userId, videoId int64) bool {
	if userId <= 0 || !repository.UserExist(sqls.DB(), userId) {
		return false
	}
	if videoId <= 0 || repository.VideoExist(sqls.DB(), videoId) {
		return false
	}
	return repository.UserFavoriteExist(sqls.DB(), userId, videoId)
}
