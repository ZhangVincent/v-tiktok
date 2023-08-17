package config

import (
	"io/ioutil"

	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

var Instance *Config

type Config struct {
	Env     string `yaml:"Env"`     // 环境：prod、dev
	BaseUrl string `yaml:"BaseUrl"` // base url
	LogFile string `yaml:"LogFile"` // 日志文件

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

	// 上传文件系统
	Uploader struct {
		Enable string `yaml:"Enable"`
		Local  struct {
			VideoPath string `yaml:"VideoPath"`
			ImagePath string `yaml:"ImagePath"`
		} `yaml:"Local"`
		Minio struct {
			Host            string `yaml:"Host"`
			Path            string `yaml:"Path"`
			Endpoint        string `yaml:"Endpoint"`
			AccessKeyID     string `yaml:"AccessKeyID"`
			SecretAccessKey string `yaml:"SecretAccessKey"`
			UseSSL          bool   `yaml:"UseSSL"`
		} `yaml:"Minio"`
	} `yaml:"Uploader"`

	// redis配置
	Redis struct {
		Url        string `yaml:"Url"`
		Key        string `yaml:"Key"`
		ExpireDays int    `yaml:"ExpireDays"`
	} `yaml:"Redis"`
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
