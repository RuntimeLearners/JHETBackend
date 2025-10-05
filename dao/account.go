package dao

import (
	"JHETBackend/common/exception"
	"JHETBackend/configs/database"
	"JHETBackend/models"
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func UpdateAccountAvatar(accountID uint64, fileUUID uuid.UUID) error {
	database.DataBase.Model(&models.AccountInfo{}).
		Where("id = ?", accountID).
		Update("avatar_fileuuid", fileUUID)
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
		panic("[!][FATAL][DAO/Account] 更新用户头像时影响了多于一行记录，请检查数据库完整性")
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

func DeleteAccount(accountID uint64) error {
	result := database.DataBase.Delete(&models.AccountInfo{}, accountID)
	//恢复软删除 database.DataBase.Unscoped().Model(&models.AccountInfo{}).Where("id = ?", accountID).Update("deleted_at", nil).Error
	//硬删除 database.DataBase.Unscoped().Delete(&models.AccountInfo{}, accountID).Error
	if result.Error != nil {
		//panic("[!][FATAL][DAO/Account] 删除用户时出错")
		return exception.SysCannotUpdate
	}
	if result.RowsAffected == 0 {
		return exception.UsrNotExisted
	}
	return nil
}
