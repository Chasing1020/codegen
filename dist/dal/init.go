
/*
Copyright © 2022 zjc <chasing1020@gmail.com>
Time: 2022/4/7-11:18
File: init.go
*/

package dal

import (
	"crud/conf"
	"github.com/go-redis/redis/v8"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

var DB *gorm.DB
var RDB *redis.Client

func init() {
	InitDB()
	InitRedis()
}

func InitDB() {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Info,
			IgnoreRecordNotFoundError: true,
			Colorful:                  false,
		},
	)
	var err error
	DB, err = gorm.Open(mysql.New(mysql.Config{DSN: conf.Conf.MySQL.DSN}), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		log.Println("DataBase connection failed: ", err)
		panic(err)
	}
	sqlDB, err := DB.DB()
	if err != nil {
		log.Println("Get generic database object sql.DB failed: ", err)
		panic(err)
	}
	sqlDB.SetConnMaxIdleTime(conf.Conf.MySQL.ConnMaxIdleTime)
	sqlDB.SetConnMaxLifetime(conf.Conf.MySQL.ConnMaxLifetime)
	sqlDB.SetMaxIdleConns(conf.Conf.MySQL.MaxIdleConn)
	sqlDB.SetMaxOpenConns(conf.Conf.MySQL.MaxOpenConn)
}

func InitRedis() {
	RDB = redis.NewClient(&redis.Options{
		Addr:         conf.Conf.Redis.Addr,
		Password:     conf.Conf.Redis.Password,
		DB:           conf.Conf.Redis.DB,
		MaxRetries:   conf.Conf.Redis.MaxRetries,
		ReadTimeout:  conf.Conf.Redis.ReadTimeout,
		WriteTimeout: conf.Conf.Redis.WriteTimeout,
	})
}
