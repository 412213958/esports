package middleware

import (
	"context"
	"database/sql"
	"github.com/gin-gonic/gin"
)

// DatabaseMiddleware 是用于处理数据库连接的中间件
func DatabaseMiddleware(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 将数据库连接添加到 Gin 的 context 中
		c.Request = c.Request.WithContext(context.WithValue(c.Request.Context(), "db", db))
		c.Next() // 调用链中的下一个处理函数
	}
}
