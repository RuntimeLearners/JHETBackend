package dao

import (
	"JHETBackend/common/exception"
	"JHETBackend/configs/database"
	"JHETBackend/models"
)

func UpdateAccountAvatar(accountID uint64, fileName string) error {
	database.DataBase.Model(&models.AccountInfo{}).
		Where("id = ?", accountID).
		Update("avatar_file", fileName)
	if database.DataBase.Error != nil {
		// 数据库层面报错（如语法错误、连接失败）
		return exception.SysCannotUpdate
	}
	if database.DataBase.RowsAffected == 0 {
		// 没有行被更新：可能是 ID 不存在，或 version 已变化
		return exception.SysCannotUpdate
	}
	return nil
}
