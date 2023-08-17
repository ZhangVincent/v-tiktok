// @description 视频业务
// @author zkp15
// @date 2023/8/15 10:20
// @version 1.0

package service

import (
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"mime/multipart"
	"path/filepath"
	"time"
	"v-tiktok/model"
	"v-tiktok/model/config"
	"v-tiktok/pkg/redis"
	"v-tiktok/pkg/sqls"
	"v-tiktok/pkg/strs"
	"v-tiktok/pkg/upload"
	"v-tiktok/repository"
	"v-tiktok/repository/redisDao"
)

// GetVideosByUserId @description 查询用户发布的视频
// @author zkp15
// @date 2023/8/16 15:22
func GetVideosByUserId(userId int64) ([]model.Video, error) {
	db := sqls.DB()
	//数据校验
	if userId <= 0 || !repository.UserExist(db, userId) {
		return nil, errors.New("user not found")
	}

	//根据userid查询视频
	videos, err := repository.GetVideosByUserId(db, userId)
	if err != nil {
		return nil, err
	}

	//返回
	return videos, nil
}

func GetVideos(latestTime int64) ([]model.Video, error) {
	videoNum := 30
	//查询redis
	videos, err := redisDao.GetVideosByTime(redis.Client(), latestTime, videoNum)
	if err == nil && len(videos) > 0 {
		return videos, nil
	}
	//如果redis中没有最新视频，就去数据库中查
	videos, err = repository.GetVideosByTime(sqls.DB(), latestTime, videoNum)
	if err != nil {
		return nil, err
	}
	//如果查到了，就添加到redis中
	if len(videos) > 0 {
		if err = redisDao.SaveVideos(redis.Client(), videos); err != nil {
			logrus.Error(err)
		}
	} else {
		//如果没查到，那么重置latestTime
		videos, err = redisDao.GetVideosByTime(redis.Client(), 0, videoNum)
		if err != nil {
			return nil, err
		}
	}
	//返回
	return videos, nil
}

func SaveVideoOnMinio(data *multipart.FileHeader, title string, userId int64) error {
	//数据校验
	if strs.IsBlank(title) {
		if strs.IsNotBlank(data.Filename) {
			title = filepath.Base(data.Filename)
		} else {
			title = fmt.Sprintf("%d_%s", userId, time.Now().Format("20060102150405"))
		}
	}

	fileName := strs.UUID()
	//保存视频到minio中
	videoSavePath, err := upload.SaveVideoOnMinio(fileName+".mp4", data)
	if err != nil {
		return errors.New("save video on minio error")
	}

	//从视频流中生成一张图片，并保存在minio中
	imageSavePath, err := upload.SaveImageOnMinio(videoSavePath, fileName+".png", 1)
	if err != nil {
		return errors.New("save cover image on minio error")
	}

	if err := sqls.DB().Transaction(func(tx *gorm.DB) error {
		//向数据库插入一条记录
		if err := repository.SaveVideo(tx, userId, videoSavePath, imageSavePath, title, 0); err != nil {
			return err
		}
		//用户发视频数+1
		if err := repository.UpdateUserWorkCount(tx, userId, 1); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return err
	}

	//返回
	return nil
}

func SaveVideoOnLocal(data *multipart.FileHeader, title string, userId int64) error {
	//数据校验
	if strs.IsBlank(title) {
		if strs.IsNotBlank(data.Filename) {
			title = filepath.Base(data.Filename)
		} else {
			title = fmt.Sprintf("%d_%s", userId, time.Now().Format("20060102150405"))
		}
	}

	fileName := strs.UUID()

	//保存视频到本地
	videoSavePath := filepath.Join(config.Instance.Uploader.Local.VideoPath, fileName+".mp4")
	if err := upload.SaveVideoOnLocal(data, videoSavePath); err != nil {
		return errors.New("save video on local error")
	}

	//从视频流中生成一张图片，并保存在本地
	imageSavePath := filepath.Join(config.Instance.Uploader.Local.ImagePath, fileName+".png")
	if err := upload.SaveImageOnLocal(videoSavePath, imageSavePath, 1); err != nil {
		return errors.New("save video on local error")
	}

	if err := sqls.DB().Transaction(func(tx *gorm.DB) error {
		//向数据库插入一条记录
		if err := repository.SaveVideo(tx, userId, videoSavePath, imageSavePath, title, 1); err != nil {
			return err
		}
		//用户发视频数+1
		if err := repository.UpdateUserWorkCount(tx, userId, 1); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return err
	}

	//返回
	return nil
}
