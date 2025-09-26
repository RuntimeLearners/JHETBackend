package userService

import (
	"JHETBackend/configs/database"
	"JHETBackend/models"
)

// GetUserByNum 根据用户编号(非id)获取用户.给学生注册填学号用
func GetUserByNum(id string) (*models.AccountInfo, error) {
	user := models.AccountInfo{}
	result := database.DataBase.Where(
		&models.AccountInfo{
			StudentID: id,
		},
	).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

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

// GetUserByEmail 根据用户邮箱获取用户
func GetUserByEmail(email string) (*models.AccountInfo, error) {
	user := models.AccountInfo{}
	result := database.DataBase.Where(
		&models.AccountInfo{
			Email: email,
		},
	).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

// GetUserByName 根据用户姓名获取用户
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
