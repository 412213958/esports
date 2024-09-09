package test

import (
	"esports/spider"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"testing"
	"time"
)

var DB *gorm.DB

func TestLoadTest(t *testing.T) {
	//自定义日志输出模板，打印sql语句
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags|log.Lshortfile),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Info,
			Colorful:      true,
		},
	)
	my := fmt.Sprintf("esports:123456@tcp(192.168.174.151:3306)/esports?charset=utf8mb4&parseTime=True&loc=Local")
	db, err := gorm.Open(mysql.Open(my), &gorm.Config{Logger: newLogger})
	if err != nil {
		panic(err)
	}
	DB = db
	spider.Transcation()
}
