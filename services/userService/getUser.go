package userService

import (
	"JHETBackend/configs/database"
	"JHETBackend/dao"
	"JHETBackend/models"
)

// WARN(MUCHEXD) 这些函数具有严重问题
// 如果用户故意使用已经存在的姓名，会导致可能的攻击
// 例如GetUserByName，用户A注册姓名为"张三"，用户B为了获取A的信息，也注册姓名为"张三"
// 然后调用GetUserByName("张三")，就读取了A的信息
// 请仅仅为前端提供 uid 到信息的接口

// GetAccountInfoByUID 根据AccountID 获取用户信息

func GetAccountInfoByUID(uid uint64) (*models.AccountInfo, error) {
	accountInfo, err := dao.GetAccountInfoByID(uid)
	if err != nil {
		return nil, err
	}
	return accountInfo, nil
}

// WARN(MUCHEXD) 命名不合理
// 根据你的代码，"编号"指的是student_id，请在代码里使用更明确的命名
// GetUserByNum 根据用户编号(非id)获取用户.给学生注册填学号用

// WARN(MUCHEXD) 建议统一命名: User->Account
// 原因是，在我们的系统中，管理员和用户的信息在一个数据表中存储

func GetUserByNum(id string) (*models.AccountInfo, error) {
	var user models.AccountInfo
	result := database.DataBase.Where("student_id = ?", id).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

// GetUserByNumIncludeDeleted 根据用户编号获取用户（包括软删除的记录）
func GetUserByNumIncludeDeleted(id string) (*models.AccountInfo, error) {
	var user models.AccountInfo
	result := database.DataBase.Unscoped().Where("id = ?", id).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

// GetUserByEmail 根据用户邮箱获取用户
func GetUserByEmailIncludeDeleted(email string) (*models.AccountInfo, error) {
	var user models.AccountInfo
	result := database.DataBase.Unscoped().Where("email = ?", email).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

// GetUserByID 根据用户ID获取用户
func GetUserByID(id uint64) (*models.AccountInfo, error) {
	var user models.AccountInfo
	result := database.DataBase.Where("id = ?", id).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

// GetUserByEmail 根据用户邮箱获取用户
func GetUserByEmail(email string) (*models.AccountInfo, error) {
	var user models.AccountInfo
	result := database.DataBase.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

// GetUserByName 根据用户姓名获取用户
func GetUserByName(name string) (*models.AccountInfo, error) {
	var user models.AccountInfo
	result := database.DataBase.Where("realname = ?", name).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
