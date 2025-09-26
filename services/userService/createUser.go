package userService

import (
	"JHETBackend/common/exception"
	"JHETBackend/configs/database"
	"JHETBackend/models"
	"crypto/sha256"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

// 向数据库保存用户信息
func CreateUser(
	studentID string,
	password string,
	realName,
	email,
	userName,
	major,
	phoneNumber string,
	permGroupID uint32,
) (*models.AccountInfo, error) {

	// var userID uint64
	var err error
	// userID, err = strconv.ParseUint(studentID, 10, 64)
	// if err != nil {
	// 	//返回的学生id有误
	// 	fmt.Println("参数错误2:", err)
	// 	return nil, exception.ApiParamError
	// }

	_, err = GetUserByNum(studentID) //判断编号是否重复
	if err == nil {
		return nil, exception.UsrAlreadyExisted
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	_, err = GetUserByEmail(email) //判断邮箱是否重复
	if err == nil {
		return nil, exception.UsrAlreadyExisted
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	//hashedPassword, err := bcrypt.GenerateFromPassword(hashPassword, 12) //cost
	hash := sha256.Sum256([]byte(password))

	// 将哈希值转换为十六进制字符串
	hashedPassword := fmt.Sprintf("%x", hash)
	// if err != nil {
	// 	//fmt.Println(err)
	// 	return nil, exception.SysPwdHashFailed
	// }
	//fmt.Println(string(hashedPassword))

	if userName == "" {
		userName = realName
	}

	user := &models.AccountInfo{
		StudentID:    studentID,
		PasswordHash: string(hashedPassword),
		RealName:     realName,
		Email:        email,
		UserName:     userName,
		Major:        major,
		PhoneNumber:  phoneNumber,
		PermGroupID:  permGroupID,
	}

	res := database.DataBase.Create(user)

	return user, res.Error
}
