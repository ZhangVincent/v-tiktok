// @description 用户信息redis存储
// @author zkp15
// @date 2023/8/17 10:39
// @version 1.0

package redisDao

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
	"v-tiktok/model"
	"v-tiktok/model/config"
)

var (
	ctx = context.Background()
)

func GetUserById(client *redis.Client, id int64) (model.User, error) {
	val, err := client.Get(ctx, generateUserKey(id)).Result()
	if err != nil {
		return model.User{}, err
	}
	var userInfo model.User
	err = json.Unmarshal([]byte(val), &userInfo)
	if err != nil {
		return model.User{}, err
	}
	return userInfo, nil
}

func SaveUser(client *redis.Client, userInfo model.User) error {
	marshal, err := json.Marshal(userInfo)
	if err != nil {
		return err
	}
	userString := string(marshal)
	err = client.Set(ctx, generateUserKey(userInfo.ID), userString, getExpireTime()).Err()
	if err != nil {
		return err
	}
	return nil
}

func getExpireTime() time.Duration {
	return 7 * time.Hour
}

func generateUserKey(id int64) string {
	template := config.Instance.Redis.Key + ":users:%d"
	key := fmt.Sprintf(template, id)
	return key
}
