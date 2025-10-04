package userService

import (
	//	"JHETBackend/common/exception"
	"JHETBackend/configs/database"
)

func UpdateUser(
	userID uint64,
	updateForm interface{},
) error {
	//更新用户信息
	err := database.DataBase.Table("account_infos").Where("id = ?", userID).Updates(updateForm).Error
	if err != nil {
		return err
	}
	return nil
}
