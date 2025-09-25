package userService

import (
	"JHETBackend/common/exception"
	"JHETBackend/configs/database"
	"JHETBackend/models"
	"errors"
	"strconv"

	"golang.org/x/crypto/bcrypt"
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

	var userID uint64
	var err error
	userID, err = strconv.ParseUint(studentID, 10, 64)
	if err != nil {
		//返回的学生id有误
		return nil, exception.ApiParamError
	}

	_, err = GetUserByID(userID)
	if err == nil {
		return nil, exception.UsrAlreadyExisted
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	hashPassword := []byte(password)
	hashedPassword, err := bcrypt.GenerateFromPassword(hashPassword, 12) //cost
	if err != nil {
		//fmt.Println(err)
		c.Error(NewWithData(exception.SysUknExc, err))
	}
	//fmt.Println(string(hashedPassword))

	user := &models.User{
		Name:      realName,
		UserName:  userName,
		Major:     major,
		Type:      usertype,
		StudentID: studentID,
		Phone:     phoneNumber,
		Email:     email,
	}

	err = EncryptUserKeyInfo(user)
	if err != nil {
		return nil, err
	}
	res := database.DataBase.Create(&user)

	return user, res.Error
}
