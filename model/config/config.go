package config

import (
	"io/ioutil"

	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

var Instance *Config

type Config struct {
	Env       string `yaml:"Env"`       // 环境：prod、dev
	BaseUrl   string `yaml:"BaseUrl"`   // base url
	LogFile   string `yaml:"LogFile"`   // 日志文件
	videoPath string `yaml:"VideoPath"` // 静态文件目录

	// 数据库配置
	DB struct {
		Url                    string `yaml:"Url"`
		MaxIdleConns           int    `yaml:"MaxIdleConns"`
		MaxOpenConns           int    `yaml:"MaxOpenConns"`
		ConnMaxIdleTimeSeconds int    `yaml:"ConnMaxIdleTimeSeconds"`
		ConnMaxLifetimeSeconds int    `yaml:"ConnMaxLifetimeSeconds"`
	} `yaml:"DB"`

	Jwt struct {
		SignKey    string `yaml:"SignKey"`
		ExpireDays int    `yaml:"ExpireDays"`
		Issuer     string `yaml:"Issuer"`
	} `yaml:"Jwt"`
}

func Init(filename string) *Config {
	Instance = &Config{}
	if yamlFile, err := ioutil.ReadFile(filename); err != nil {
		logrus.Error(err)
	} else if err = yaml.Unmarshal(yamlFile, Instance); err != nil {
		logrus.Error(err)
	}
	return Instance
}
