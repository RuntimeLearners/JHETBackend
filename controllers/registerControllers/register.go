package registerControllers

import (
	"JHETBackend/common/exception"
	"JHETBackend/services/userService"
	"JHETBackend/utils"
	"errors"
	"fmt"
	"regexp"

	"github.com/gin-gonic/gin"
)

type UserInfo struct {
	StudentID   string `json:"studentID"   binding:"required"`
	RealName    string `json:"realname"  binding:"required"`
	Email       string `json:"email"  binding:"required"`
	Password    string `json:"password"` //管理员注册时允许为空,为空使用默认密码abc132456
	UserName    string `json:"username"`
	Major       string `json:"major"`       //专业
	PhoneNumber string `json:"phoneNumber"` //手机号
	UserType    string `json:"userType"`    //用户类型, 在转到server时映射为对应的PermGroupID
	Activation  *bool   `json:"activation"`  //账户激活状态(保留,用于验证邮箱是否存在)
}

func CreateUserUser(c *gin.Context) { //普通用户注册,强制绑定权限组PermGroupID=1
	CreateUser(c, false)
}

func CreateUserAdmin(c *gin.Context) { //管理员新增用户
	CreateUser(c, true)
}

func CreateUser(c *gin.Context, isAdmin bool) { //管理员新增用户
	var postForm UserInfo
	err := c.ShouldBindJSON(&postForm)
	if err != nil {
		c.Error(exception.ApiParamError)
		fmt.Println("参数错误0:", err)
		return
	}
	//初步处理
	if postForm.Password == "" {
		postForm.Password = "abc132456" //默认密码
	}
	if postForm.Activation == nil {
		postForm.Activation = new(bool)
		*postForm.Activation = true //默认激活
	}
	//判定数据合规性
	if len(postForm.Password) < 6 {
		c.Error(exception.UsrPasswordErr)
		return
	}
	emailRegex := `^(([^<>()[\]\\.,;:\s@\"]+(\.[^<>()[\]\\.,;:\s@\"]+)*)|(\".+\"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$`
	matched, _ := regexp.MatchString(emailRegex, postForm.Email)
	if !matched {
		c.Error(exception.UsrEmailErr)
		return
	}
	
	permGroupID := userService.MapPerm(c) //映射权限组id

	fmt.Println("注册信息:", postForm)
	user, err := userService.CreateUser(
		postForm.StudentID,
		postForm.Password,
		postForm.RealName,
		postForm.Email,
		postForm.UserName,
		postForm.Major,
		postForm.PhoneNumber,
		*postForm.Activation,
		permGroupID,
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
