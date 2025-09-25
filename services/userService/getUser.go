package userService

import (
	"JHETBackend/configs/database"
	"JHETBackend/models"
)

// GetUserByID 根据用户ID获取用户
func GetUserByID(id uint64) (*models.AccountInfo, error) {
	user := models.AccountInfo{}
	result := database.DataBase.Where(
		&models.AccountInfo{
			ID: id,
		},
	).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

// GetUserByName 根据用户ID获取用户
func GetUserByName(name string) (*models.AccountInfo, error) {
	user := models.AccountInfo{}
	result := database.DataBase.Where(
		&models.AccountInfo{
			RealName: name,
		},
	).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
