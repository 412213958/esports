package utils

import (
	"errors"
	"regexp"
)

// ValidateEmail 验证电子邮件格式
func ValidateEmail(email string) error {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(email) {
		return errors.New("无效的电子邮件格式")
	}
	return nil
}

// ValidateUsername 验证用户名格式（例如，只允许字母、数字和下划线，长度在 3 到 20 个字符之间）
func ValidateUsername(username string) error {
	usernameRegex := regexp.MustCompile(`^[a-zA-Z0-9_]{3,20}$`)
	if !usernameRegex.MatchString(username) {
		return errors.New("无效的用户名格式")
	}
	return nil
}

// ValidatePassword 验证密码强度（例如，至少 8 个字符，包含大写字母、小写字母、数字和特殊字符）
func ValidatePassword(password string) error {
	passwordRegex := regexp.MustCompile(`^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)(?=.*[@$!%*?&])[A-Za-z\d@$!%*?&]{8,}$`)
	if !passwordRegex.MatchString(password) {
		return errors.New("密码强度不足，至少 8 个字符，包含大写字母、小写字母、数字和特殊字符")
	}
	return nil
}
