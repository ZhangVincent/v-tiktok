// @description redis配置
// @author zkp15
// @date 2023/8/17 10:35
// @version 1.0

package redis

import (
	"context"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
	"v-tiktok/model/config"
)

type RedisConfig struct {
	Url        string `yaml:"Url"`
	Key        string `yaml:"Key"`
	ExpireDays int    `yaml:"ExpireDays"`
}

var (
	redisClient *redis.Client
	ctx         = context.Background()
	expireDays  int
)

func Open(redisConfig RedisConfig) error {
	opt, err := redis.ParseURL(redisConfig.Url)
	if err != nil {
		return err
	}
	redisClient = redis.NewClient(opt)
	expireDays = config.Instance.Redis.ExpireDays
	return nil
}

func Client() *redis.Client {
	return redisClient
}

func ExpireDays() int {
	return expireDays
}

func Close() {
	if redisClient == nil {
		return
	}
	if err := redisClient.Close(); nil != err {
		logrus.Errorf("Disconnect from redis failed: %s", err.Error())
	}
	logrus.Info("Disconnect from redis success")
}
