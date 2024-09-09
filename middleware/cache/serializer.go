package cache

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"net/http"
	"sort"
	"strings"
	"time"
)

// responseRecorder 用于记录响应的数据
type responseRecorder struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

// Write 方法实现了记录写入的数据
func (r *responseRecorder) Write(b []byte) (int, error) {
	r.body.Write(b)
	return r.ResponseWriter.Write(b)
}

// CacheMiddleware 生成一个用于处理缓存的 Gin 中间件函数
func CacheMiddleware(redisClient *redis.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		key := generateCacheKey(c.Request)

		// 尝试从 Redis 中获取缓存数据
		cachedData, err := redisClient.Get(key).Bytes()
		if err == nil && cachedData != nil {
			var responseData interface{}
			// 尝试将缓存数据解析为特定的响应结构
			if err := json.Unmarshal(cachedData, &responseData); err == nil {
				// 以 JSON 格式将缓存数据写入响应
				c.JSON(c.Writer.Status(), responseData)
				c.Abort()
				return
			}
		}

		// 创建自定义的 responseRecorder 来记录响应
		recorder := &responseRecorder{
			ResponseWriter: c.Writer,
			body:           &bytes.Buffer{},
		}
		c.Writer = recorder

		// 继续执行后续的请求处理链
		c.Next()

		// 从 recorder 中获取响应数据
		responseData := recorder.body.Bytes()

		// 尝试解码取出的数据
		var decodedData interface{}
		if err := json.NewDecoder(bytes.NewReader(responseData)).Decode(&decodedData); err != nil {
			fmt.Println("Error decoding response data:", err)
			return
		}

		// 将解码后的响应数据以 JSON 格式存储到 Redis 中
		dataToCache, err := json.Marshal(decodedData)
		if err != nil {
			fmt.Println("Error marshaling data for cache:", err)
			return
		}
		if err := redisClient.Set(key, dataToCache, time.Hour); err != nil {
			//fmt.Println("Error setting data in cache:", err)
			return
		}
	}
}

// generateCacheKey 根据请求生成包含参数的缓存键
func generateCacheKey(r *http.Request) string {
	// 获取请求的路径和方法
	baseKey := r.URL.Path + r.Method

	// 提取请求参数
	values := r.URL.Query()

	// 对参数进行排序以确保键的一致性
	keys := make([]string, 0, len(values))
	for k := range values {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var params []string
	for _, key := range keys {
		for _, value := range values[key] {
			params = append(params, fmt.Sprintf("%s=%s", key, value))
		}
	}

	if len(params) > 0 {
		baseKey += "?" + strings.Join(params, "&")
	}

	return baseKey
}
