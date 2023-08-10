// @description 数据库连接
// @author zkp15
// @date 2023/8/9 11:16
// @version 1.0

package sqls

import (
	"database/sql"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"time"
)

type DbConfig struct {
	Url                    string `yaml:"Url"`
	MaxIdleConns           int    `yaml:"MaxIdleConns"`
	MaxOpenConns           int    `yaml:"MaxOpenConns"`
	ConnMaxIdleTimeSeconds int    `yaml:"ConnMaxIdleTimeSeconds"`
	ConnMaxLifetimeSeconds int    `yaml:"ConnMaxLifetimeSeconds"`
}

var (
	db    *gorm.DB
	sqlDB *sql.DB
)

func Open(dbConfig DbConfig, config *gorm.Config, models ...interface{}) (err error) {
	if config == nil {
		config = &gorm.Config{}
	}

	if config.NamingStrategy == nil {
		config.NamingStrategy = schema.NamingStrategy{
			TablePrefix:   "t_",
			SingularTable: true,
		}
	}

	if db, err = gorm.Open(mysql.Open(dbConfig.Url), config); err != nil {
		logrus.Errorf("opens database failed: %s", err.Error())
		return
	}

	if sqlDB, err = db.DB(); err == nil {
		sqlDB.SetMaxIdleConns(dbConfig.MaxIdleConns)
		sqlDB.SetMaxOpenConns(dbConfig.MaxOpenConns)
		sqlDB.SetConnMaxIdleTime(time.Duration(dbConfig.ConnMaxIdleTimeSeconds) * time.Second)
		sqlDB.SetConnMaxLifetime(time.Duration(dbConfig.ConnMaxLifetimeSeconds) * time.Second)
	} else {
		logrus.Error(err)
	}

	if err = db.AutoMigrate(models...); nil != err {
		logrus.Errorf("auto migrate tables failed: %s", err.Error())
	}
	return
}

func DB() *gorm.DB {
	return db
}

func Close() {
	if sqlDB == nil {
		return
	}
	if err := sqlDB.Close(); nil != err {
		logrus.Errorf("Disconnect from database failed: %s", err.Error())
	}
}
