package router

import (
	"esports/api/v1/imageprocess"
	"esports/api/v1/match"
	"esports/api/v1/team"
	"esports/api/v1/tournament"
	_ "esports/docs"
	"esports/middleware"
	"esports/middleware/cache"
	"esports/spider"
	"esports/utils"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-co-op/gocron"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"sync"
	"time"
)

// skipCacheDuringUpdate 控制在定时器更新缓存时是否跳过 Redis 缓存
var skipCacheDuringUpdate = false

// lock 用于在缓存更新期间加锁
var lock sync.Mutex

// SetSkipCacheDuringUpdate 设置在定时器更新缓存时是否跳过 Redis 缓存
func SetSkipCacheDuringUpdate(skip bool) {
	skipCacheDuringUpdate = skip
}

func Router() *gin.Engine {
	r := gin.Default()

	// 添加 CORS 中间件
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		AllowCredentials: true,
	}))

	utils.InitConfig()
	utils.InitRedis()

	// 创建 Redis 客户端
	redisClient := utils.RDB
	sqlClient := utils.DB

	// 创建缓存中间件
	cacheMiddleware := cache.CacheMiddleware(redisClient)

	// 应用缓存中间件到整个引擎
	r.Use(func(c *gin.Context) {
		// 在缓存更新期间阻塞
		lock.Lock()
		defer lock.Unlock()

		if skipCacheDuringUpdate {
			// 直接跳过缓存处理
			c.Next()
			return
		}
		cacheMiddleware(c)
	})

	// 创建定时任务调度器
	s := gocron.NewScheduler(time.UTC)

	// 定义更新缓存的任务函数
	s.Every(20).Seconds().Do(func() {
		SetSkipCacheDuringUpdate(true)
		cache.UpdateAllCache(redisClient)
		SetSkipCacheDuringUpdate(false)
		fmt.Println("----------------缓存已更新------------")
	})
	//spider.Transcation()
	// 创建一个定时器，每小时执行一次 Transcation 函数
	timer := time.NewTicker(time.Hour)
	go func() {
		for range timer.C {
			spider.Transcation()
		}
	}()

	// 调用中间件函数
	settleFunc := middleware.SettlePredictionMiddleware(sqlClient)

	// 模拟定时调用
	timer = time.NewTicker(5 * time.Second)
	for range timer.C {
		settleFunc()
	}

	// 启动定时任务调度器
	s.StartAsync()

	// Swagger 文档和静态资源
	setupSwagger(r)
	setupStaticFiles(r)

	// 用户路由
	//userRouter(r)

	// 图片处理路由
	imageProcessRouter(r)

	// 赛事路由
	matchRouter(r)

	// 战队路由
	teamRouter(r)

	//赛事信息路由
	tournamentRouter(r)

	return r
}

func setupSwagger(r *gin.Engine) {
	// Swagger 相关配置
	//url := ginSwagger.URL("/swagger/doc.json") // JSON 文件的 URL 或 `URL("/swagger/doc.yaml")` 如果使用 YAML 文件

	// 添加 Swagger 路由
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

func setupStaticFiles(r *gin.Engine) {
	r.Static("/resource", "resource/")
	//r.LoadHTMLGlob("views/**/*")
}

func userRouter(r *gin.Engine) {
	//userGroup := r.Group("/api/v1/users")
	//userGroup.Use(middleware.AuthMiddleware())
	//{
	//    //userGroup.GET("", user.GetUsers)
	//    //userGroup.GET("/:id", user.GetUserByID)
	//    //userGroup.POST("", user.CreateUser)
	//    //userGroup.PUT("/:id", user.UpdateUser)
	//    //userGroup.DELETE("/:id", user.DeleteUser)
	//}
}

func imageProcessRouter(r *gin.Engine) {
	imageGroup := r.Group("/api/v1/imageprocess")
	{
		imageGroup.GET("/GetImages", imageprocess.GetImages)
	}
}

func matchRouter(r *gin.Engine) {
	tournamentGroup := r.Group("/api/v1/match/")
	{
		tournamentGroup.GET("GetRaceCalendar", match.GetRaceCalendarList)
		tournamentGroup.GET("/GetRaceCalendarListByTitle", match.GetRaceCalendarListByTitle)
		tournamentGroup.GET("/GetRaceCalendarListWF", match.GetRaceCalendarListWF)
		tournamentGroup.GET("/GetRaceCalendarListByABCD", match.GetRaceCalendarListByABCD)
		tournamentGroup.GET("GetAllData", match.GetAllData)
		tournamentGroup.GET("GetVideoList", match.GetVideoList)
		tournamentGroup.GET("GetAllRaceCalendarList", match.GetAllRaceCalendarList)
		tournamentGroup.GET("GetMatchesList", match.GetMatchesList)
		tournamentGroup.GET("GetMatches", match.GetMatches)
		tournamentGroup.GET("GetMatches1", match.GetMatches1)
	}
}

func teamRouter(r *gin.Engine) {
	teamGroup := r.Group("/api/v1/team/")
	{
		teamGroup.GET("GetLegacyList", team.GetLegacyList)
		teamGroup.GET("GetPlayerAndTeamDataByTeam", team.GetPlayerAndTeamDataByTeam)
		teamGroup.GET("GetAllPlayerAndTeamDataByTeam", team.GetAllPlayerAndTeamDataByTeam)
		teamGroup.GET("Dynamic odds", team.Dynamicodds)

	}
}

func tournamentRouter(r *gin.Engine) {
	teamGroup := r.Group("/api/v1/tournament/")
	{
		teamGroup.GET("GetTournaments", tournament.GetTournaments)
		teamGroup.GET("GetWeiboData", tournament.GetWeiboData)
		teamGroup.GET("GetArticlesData", tournament.GetArticlesData)
		teamGroup.GET("GetArticles1Data", tournament.GetArticles1Data)
		teamGroup.GET("GetCommentsData", tournament.GetCommentsData)
		teamGroup.GET("GetYCData", tournament.GetYCData)

	}
}
