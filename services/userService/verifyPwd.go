package userService

import (
	"JHETBackend/common/exception"
	"JHETBackend/models"
	"crypto/sha256"
	"fmt"
)

// Verify Password using SHA256
func VerifyPwd(user *models.AccountInfo, password string) error {
	// 计算输入密码的 SHA256 哈希值
	hash := sha256.Sum256([]byte(password))

	// 将哈希值转换为十六进制字符串
	hashedPassword := fmt.Sprintf("%x", hash)

	// 比较存储的密码哈希和计算的哈希值
	if user.PasswordHash != hashedPassword {
		return exception.UsrPasswordErr
	}

	return nil
}
