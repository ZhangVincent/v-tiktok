package main

import (
	"flag"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"io"
	"os"
	"time"
	"v-tiktok/model"
	"v-tiktok/model/config"
	"v-tiktok/pkg/message"
	"v-tiktok/pkg/sqls"
)

var configFile = flag.String("config", "./v-tiktok.yaml", "配置文件路径")

func init() {
	flag.Parse()

	// 初始化配置
	conf := config.Init(*configFile)

	// 初始化日志
	if file, err := os.OpenFile(conf.LogFile+time.Now().Format("20060102")+".log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666); err == nil {
		logrus.SetOutput(io.MultiWriter(os.Stdout, file))
		logrus.SetLevel(logrus.WarnLevel)
	} else {
		logrus.SetOutput(os.Stdout)
		logrus.Error(err)
	}

	// 连接数据库
	gormConf := &gorm.Config{
		Logger: logger.New(logrus.StandardLogger(), logger.Config{
			SlowThreshold:             time.Second,
			Colorful:                  true,
			LogLevel:                  logger.Warn,
			IgnoreRecordNotFoundError: true,
		}),
	}
	if err := sqls.Open(conf.DB, gormConf, model.Models...); err != nil {
		logrus.Error(err)
	}

}

func main() {
	go message.RunMessageServer()

	Router()
}
