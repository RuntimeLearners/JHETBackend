package userService

import (
	"JHETBackend/common/exception"
	"JHETBackend/configs/database"
	"JHETBackend/models"
	"crypto/sha256"
	"fmt"
)

func ChangePwd(accountInfo *models.AccountInfo, oldPassword, newPassword string) error {
	// 更改密码
	var err error

	oldPasswordHash := sha256.Sum256([]byte(oldPassword))
	hashedOldPassword := fmt.Sprintf("%x", oldPasswordHash)
	// 检查旧密码是否正确
	if accountInfo.PasswordHash != hashedOldPassword {
		return exception.UsrPasswordErr
	}

	// 更新密码
	newPasswordHash := sha256.Sum256([]byte(newPassword))
	hashedNewPassword := fmt.Sprintf("%x", newPasswordHash)

	err = database.DataBase.Table("account_infos").Where("id = ?", accountInfo.ID).Updates(map[string]interface{}{"password_hash": hashedNewPassword}).Error
	if err != nil {
		return err
	}

	return nil
}
