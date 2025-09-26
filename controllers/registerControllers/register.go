package registerControllers

import (
	"JHETBackend/common/exception"
	"JHETBackend/services/userService"
	"JHETBackend/utils"
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
)

type UserInfo struct {
	StudentID   string `json:"studentID"   binding:"required"`
	RealName    string `json:"realname"  binding:"required"`
	Email       string `json:"email"  binding:"required"`
	Password    string `json:"password"   binding:"required"`
	UserName    string `json:"username"`
	Major       string `json:"major"`       //专业
	PhoneNumber string `json:"phoneNumber"` //手机号
	//PermGroupID
}

func CreateStudentUser(c *gin.Context) { //学生用户注册,强制绑定权限组PermGroupID=1
	var postForm UserInfo
	err := c.ShouldBindJSON(&postForm)
	if err != nil {
		c.Error(exception.ApiParamError)
		fmt.Println("参数错误0:", err)
		return
	}
	fmt.Println("注册信息:", postForm)
	user, err := userService.CreateUser(
		postForm.StudentID,
		postForm.Password,
		postForm.RealName,
		postForm.Email,
		postForm.UserName,
		postForm.Major,
		postForm.PhoneNumber,
		1, //用户类型 PermGroupID 需要修改
	)
	if err != nil {
		if errors.Is(err, exception.ApiParamError) {
			fmt.Println("参数错误1:", err)
			c.Error(exception.ApiParamError)
		} else if errors.Is(err, exception.UsrAlreadyExisted) {
			fmt.Println("用户已存在:", err)
			c.Error(exception.UsrAlreadyExisted)
		} else {
			fmt.Println("读取失败0:", err)
			c.Error(exception.SysCannotLoadFromDB)

		}
		return
	}

	utils.JsonSuccessResponse(c, "注册成功", map[string]interface{}{
		"userID": user.ID,
	})
}
