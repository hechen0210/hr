package config

import (
	"time"

	goredis "github.com/go-redis/redis"
	"github.com/hechen0210/utils/config"
	"github.com/hechen0210/utils/mongo"
	"github.com/hechen0210/utils/mysql"
	"github.com/hechen0210/utils/redis"
	"github.com/jinzhu/gorm"
	"gopkg.in/mgo.v2"
)

type Storage struct {
	mysql *mysql.DB
	redis *goredis.Client
	mgo   *mgo.Session
}

var storage *Storage

func NewMysql(config *config.Section) (*Storage, error) {
	mysql := (&mysql.Config{
		Host:          config.Get("host").ToString(),
		Port:          config.Get("port").ToString(),
		User:          config.Get("user").ToString(),
		Password:      config.Get("password").ToString(),
		DbName:        config.Get("dbName").ToString(),
		SingularTable: true,
	}).New()
	if mysql.Error != nil {
		return nil, mysql.Error
	}
	mysql.Client.Callback().Update().Register("gorm:update_time_stamp", func(scope *gorm.Scope) {
		if _, ok := scope.Get("gorm:update_column"); !ok {
			_ = scope.SetColumn("UpdatedAt", time.Now().Unix())
		}
	})
	storage = &Storage{
		mysql: mysql,
	}
	return storage, nil
}

func (s *Storage) NewRedis(config *config.Section) (*Storage, error) {
	conn := redis.Config{
		Host: config.Get("host").ToString(),
		Port: config.Get("port").ToString(),
		Auth: config.Get("auth").ToString(),
	}.New()
	if conn.Error != nil {
		return nil, conn.Error
	}
	s.redis = conn.Client
	return s, nil
}

func (s *Storage) NewMongo(config *config.Section) (*Storage, error) {
	conn := mongo.Config{
		Host:     config.Get("host").ToSString(),
		DataBase: config.Get("dbName").ToString(),
		User:     config.Get("user").ToString(),
		Password: config.Get("password").ToString(),
	}.New()
	if conn.Error != nil {
		return nil, conn.Error
	}
	s.mgo = conn.Client
	return s, nil
}

func GetStorage() *Storage {
	return storage
}

func GetDB() *mysql.DB {
	return storage.mysql
}

func GetDbClient() *gorm.DB {
	return storage.mysql.Client
}

func GetRedis() *goredis.Client {
	return storage.redis
}

func GetMgo() *mgo.Session {
	return storage.mgo
}
