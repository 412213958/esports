package utils

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"runtime"
)

// HandleError 处理错误并返回相应的 HTTP 状态码和错误消息
func HandleError(c *gin.Context, err error, statusCode int) {
	// 获取调用栈信息
	pc, file, line, _ := runtime.Caller(1)
	fn := runtime.FuncForPC(pc)

	// 打印详细的错误日志
	fmt.Printf("Error in %s (%s:%d): %v\n", fn.Name(), file, line, err)

	c.JSON(statusCode, gin.H{"error": err.Error()})
	// 或者
	//c.String(statusCode, err.Error())
}

// NewError 创建一个带有自定义消息的错误
func NewError(message string) error {
	return errors.New(message)
}

// WrapError 包装一个错误并添加额外的上下文信息
func WrapError(err error, context string) error {
	return fmt.Errorf("%s: %w", context, err)
}
