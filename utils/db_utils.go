package utils

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

func InitConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

var (
	DB  *gorm.DB
	RDB *redis.Client
)

func InitMySql() {
	//自定义日志输出模板，打印sql语句
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags|log.Lshortfile),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Info,
			Colorful:      true,
		},
	)
	dbUser := viper.GetString("mysql.username")
	dbPass := viper.GetString("mysql.password")
	dbHost := viper.GetString("mysql.host")
	dbPort := viper.GetString("mysql.port")
	dbName := viper.GetString("mysql.database")
	my := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser, dbPass, dbHost, dbPort, dbName)
	db, err := gorm.Open(mysql.Open(my), &gorm.Config{Logger: newLogger})
	if err != nil {
		panic(err)
	}
	DB = db
}

func InitRedis() {
	redisHost := viper.GetString("redis.host")
	redisPort := viper.GetString("redis.port")
	redisPass := viper.GetString("redis.password")

	RDB = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", redisHost, redisPort),
		Password: redisPass,
	})

	_, err := RDB.Ping().Result()
	if err != nil {
		panic(err)
	}
}
