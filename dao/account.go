package dao

import (
	"JHETBackend/common/exception"
	"JHETBackend/configs/database"
	"JHETBackend/models"
	"errors"

	"gorm.io/gorm"
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
	if database.DataBase.RowsAffected > 1 {
		// 更新了多于一行，说明有严重问题
		panic("[!]FATAL][DAO/Account] 更新用户头像时影响了多于一行记录，请检查数据库完整性")
	}
	return nil
}

func GetAccountInfoByID(accountID uint64) (*models.AccountInfo, error) {
	var accountInfo []models.AccountInfo
	err := database.DataBase.Where("id = ?", accountID).Find(&accountInfo).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, exception.UsrNotExisted
		}
		return nil, exception.SysCannotReadDB
	}
	if len(accountInfo) > 1 {
		// 一般情况下不可能出现，出现了数据库包有问题的情况
		panic("[!][FATAL][DAO/Account] 查询用户时返回了多于一行记录，请检查数据库完整性")
	}
	return &accountInfo[0], nil
}
