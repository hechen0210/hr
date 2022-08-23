package config

import (
	"github.com/hechen0210/utils/config"
)

type MyError struct {
	Code    int
	Message string
}

var configData *config.ConfigData

// 加载配置文件
func LoadConfig(file string) error {
	var err error
	configData, err = config.Load(&config.Config{
		FileName: file,
		Use:      "chiefByFile",
	})
	if err != nil {
		return err
	}
	return nil
}

// 初始化数据库、Redis连接
func InitStorage() error {
	storage, err := NewMysql(configData.GetSection("mysql"))
	if err != nil {
		return err
	}
	storage, err = storage.NewRedis(configData.GetSection("redis"))
	if err != nil {
		return err
	}
	return nil
}

// 获取配置信息
func GetConfig() *config.ConfigData {
	return configData
}
