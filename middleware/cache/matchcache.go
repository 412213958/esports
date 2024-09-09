package cache

import (
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
	"log"
	"net/http"
	"strings"
	"time"
)

// UpdateCache 依次读取 Redis 缓存的键，转换为请求并重新写入 Redis 缓存
func UpdateCache(redisClient *redis.Client) {
	// 获取所有的键
	keys, err := redisClient.Keys("*").Result()
	if err != nil {
		log.Println("Error getting keys from Redis:", err)
		return
	}

	for _, key := range keys {
		updateRequest(key, redisClient)
	}
}

// UpdateAllCache 依次读取 Redis 缓存的键，转换为请求并重新写入 Redis 缓存
func UpdateAllCache(redisClient *redis.Client) {
	// 获取所有的键
	keys := []string{"/api/v1/match/GetRaceCalendarListByABCDGET?types=A", "/api/v1/match/GetRaceCalendarGET", "/api/v1/match/GetRaceCalendarListByTitleGET?title=2024LPL夏季赛",
		"/api/v1/match/GetVideoListGET?pagenum=1&pagesize=10&type=1", "/api/v1/match/GetAllDataGET", "/api/v1/imageprocess/GetImagesGET", "/api/v1/team/GetPlayerAndTeamDataByTeamGET?team=UP",
		"/api/v1/match/GetVideoListGET?type=1", "/api/v1/match/GetMatches1GET", "/api/v1/match/GetMatchesGET", "/api/v1/match/GetMatches1GET", "/api/v1/match/GetMatchesGET", "/api/v1/match/GetMatchesGET",
		"/api/v1/team/GetLegacyListGET", "/api/v1/match/GetRaceCalendarListWFGET?typeSF=1", "/api/v1/team/GetAllPlayerAndTeamDataByTeamGET", "/api/v1/match/GetAllRaceCalendarListGET",
		"/api/v1/tournament/GetArticles1DataGET", "/api/v1/tournament/GetArticlesDataGET", "/api/v1/tournament/GetCommentsDataGET",
	}

	for _, key := range keys {
		updateAllRequest(key, redisClient)
	}
}

// updateRequest 模拟对请求结构体的更新操作，并将结果写入 Redis
func updateRequest(key string, redisClient *redis.Client) string {
	// 去除从后往前的第一个 "Get" 字符串
	lastGetIndex := strings.LastIndex(key, "GET")
	if lastGetIndex != -1 {
		processedKey := key[:lastGetIndex] + key[lastGetIndex+len("Get"):]
		// 加上指定前缀
		url := "http://121.36.26.12:8081" + processedKey

		fmt.Println("url=", url)

		// 模拟发送 GET 请求
		resp, err := http.Get(url)
		if err != nil {
			fmt.Println("GET request error:", err)
			return ""
		}
		defer resp.Body.Close()

		// 处理响应数据，这里简单返回响应状态码
		result := fmt.Sprintf("Response status code: %d", resp.StatusCode)

		// 尝试解码取出的数据
		var decodedData interface{}
		if err := json.NewDecoder(resp.Body).Decode(&decodedData); err != nil {
			fmt.Println("Error decoding response data:", err)
			return ""
		}

		// 将解码后的响应数据以 JSON 格式存储到 Redis 中
		dataToCache, err := json.Marshal(decodedData)
		if err != nil {
			fmt.Println("Error marshaling data for cache:", err)
			return ""
		}
		if err := redisClient.Set(key, dataToCache, time.Hour); err != nil {
			//fmt.Println("Error setting data in cache:", err)
			return ""
		}

		return result
	}
	return ""
}

// updateRequest 模拟对请求结构体的更新操作，并将结果写入 Redis
func updateAllRequest(key string, redisClient *redis.Client) string {
	// 去除从后往前的第一个 "Get" 字符串
	lastGetIndex := strings.LastIndex(key, "GET")
	if lastGetIndex != -1 {
		processedKey := key[:lastGetIndex] + key[lastGetIndex+len("Get"):]
		// 加上指定前缀
		url := "http://121.36.26.12:8081" + processedKey

		fmt.Println("url=", url)

		// 模拟发送 GET 请求
		resp, err := http.Get(url)
		if err != nil {
			fmt.Println("GET request error:", err)
			return ""
		}
		defer resp.Body.Close()

		// 处理响应数据，这里简单返回响应状态码
		result := fmt.Sprintf("Response status code: %d", resp.StatusCode)

		// 尝试解码取出的数据
		var decodedData interface{}
		if err := json.NewDecoder(resp.Body).Decode(&decodedData); err != nil {
			fmt.Println("Error decoding response data:", err)
			return ""
		}

		// 将解码后的响应数据以 JSON 格式存储到 Redis 中
		dataToCache, err := json.Marshal(decodedData)
		if err != nil {
			fmt.Println("Error marshaling data for cache:", err)
			return ""
		}
		if err := redisClient.Set(key, dataToCache, time.Hour); err != nil {
			//fmt.Println("Error setting data in cache:", err)
			return ""
		}

		return result
	}
	return ""
}
