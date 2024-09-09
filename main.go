package main

import (
	router "esports/routers"
	"esports/utils"
	"log"
)

func main() {
	// 初始化数据库连接
	utils.InitConfig()
	utils.InitMySql()
	//python.AutomateCrawler()
	// 创建 Gin 路由
	r := router.Router()

	// 启动 Gin 服务器
	if err := r.Run(":8081"); err != nil {
		log.Fatalf("服务器启动失败: %v", err)
	}
}
