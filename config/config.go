package config

import (
	"fmt"
	//_ "github.com/go-sql-driver/mysql" // 导入 MySQL 驱动
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

// DBConfig 存储数据库配置信息
type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

// GetDBConfigFromYAML 从 YAML 文件中获取数据库配置
func GetDBConfigFromYAML() (*DBConfig, error) {
	viper.SetConfigType("yaml")     // 设置配置文件类型为 YAML
	viper.AddConfigPath("./config") // 设置配置文件的搜索路径
	err := viper.ReadInConfig()     // 读取配置文件
	if err != nil {
		return nil, fmt.Errorf("读取配置文件失败: %v", err)
	}

	var config DBConfig
	err = viper.Unmarshal(&config)
	if err != nil {
		return nil, fmt.Errorf("解析配置失败: %v", err)
	}

	return &config, nil
}

// GetDSN 返回用于数据库连接的数据源名称（DSN）
func (c *DBConfig) GetDSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", c.User, c.Password, c.Host, c.Port, c.DBName)
}

// NewDatabaseConnection 使用 GORM 初始化并返回一个新的数据库连接
func NewDatabaseConnection() (*gorm.DB, error) {
	// 获取数据库配置
	dbConfig, err := GetDBConfigFromYAML()
	if err != nil {
		return nil, err
	}

	// 创建数据源名称（DSN）
	dsn := dbConfig.GetDSN()

	// 使用 GORM 连接 MySQL 数据库
	gormDB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("连接数据库失败: %v", err)
	}

	// 测试数据库连接
	db, err := gormDB.DB()
	if err != nil {
		return nil, fmt.Errorf("获取底层数据库连接失败: %v", err)
	}
	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, fmt.Errorf("数据库连接失败: %v", err)
	}

	log.Println("数据库连接成功")
	return gormDB, nil
}
