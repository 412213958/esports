package utils

import (
	"errors"
	model "esports/models"
	"github.com/dgrijalva/jwt-go"
	"time"
)

// SecretKey 用于签署 JWT 令牌的密钥
var SecretKey = []byte("your_secret_key")

// GenerateJWT 生成 JWT 令牌
func GenerateJWT(userID int) (string, error) {
	// 创建令牌声明
	claims := jwt.MapClaims{
		"user": &model.Users{
			// 设置其他用户字段
		},
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	}

	// 创建令牌对象，并使用 HS256 算法进行签署
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 生成签名字符串
	tokenString, err := token.SignedString(SecretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// ParseJWT 解析 JWT 令牌
func ParseJWT(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// 验证签名方法是否为预期的 HS256
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("无效的签名方法")
		}
		return SecretKey, nil
	})

	if err != nil {
		return nil, err
	}

	// 验证令牌是否有效
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("无效的令牌")
}
