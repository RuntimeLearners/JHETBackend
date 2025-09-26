package dao

import (
	"JHETBackend/configs/database"
	"JHETBackend/models"
)

func UpdateAccountAvatar(accountID uint64, fileName string) {
	database.DataBase.Model(&models.AccountInfo{}).
		Where("id = ?", accountID).
		Update("avatar_file", fileName)
}
