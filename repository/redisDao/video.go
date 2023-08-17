// @description 视频信息redis存储
// @author zkp15
// @date 2023/8/17 10:39
// @version 1.0

package redisDao

import (
	"encoding/json"
	"errors"
	"github.com/redis/go-redis/v9"
	"v-tiktok/model"
	"v-tiktok/model/config"
	"v-tiktok/pkg/strs"
)

func GetVideosByTime(client *redis.Client, latestTime int64, videosNum int) ([]model.Video, error) {
	idx, _ := client.ZRank(ctx, generateVideoKey(), strs.ItoA(latestTime*-1)).Result()

	result, err := client.ZRangeWithScores(ctx, generateVideoKey(), idx, idx+int64(videosNum)).Result()
	if err != nil {
		return nil, err
	}

	n := len(result)
	if n == 0 {
		return nil, errors.New("no more new videos")
	}

	videos := make([]model.Video, 0, n)
	for _, z := range result {
		userString, ok := z.Member.(string)
		if !ok {
			continue
		}

		var video model.Video
		err = json.Unmarshal([]byte(userString), &video)
		if err != nil {
			continue
		}

		videos = append(videos, video)
	}

	return videos, nil
}

func SaveVideos(client *redis.Client, videos []model.Video) error {
	for _, v := range videos {
		if err := SaveVideo(client, v); err != nil {
			return err
		}
	}
	return nil
}

func SaveVideo(client *redis.Client, videos model.Video) error {
	marshal, err := json.Marshal(videos)
	if err != nil {
		return err
	}
	userString := string(marshal)
	zAddResult := client.ZAdd(ctx, generateVideoKey(), redis.Z{Score: float64(videos.CreatedAt * (-1)), Member: userString})
	if zAddResult.Err() != nil {
		return zAddResult.Err()
	}
	return nil
}

func generateVideoKey() string {
	return config.Instance.Redis.Key + ":videoList"
}
