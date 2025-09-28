package loginControllers

import (
	"JHETBackend/common/exception"
	"JHETBackend/common/webtoken"
	"JHETBackend/models"
	"JHETBackend/services/userService"
	"JHETBackend/utils"
	"errors"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type passwordLoginForm struct {
	Account  string `json:"account" binding:"required"` //返回的姓名或id
	Password string `json:"password" binding:"required"`
	UserType string `json:"userType"`                      //用户类型 0学生 1教师 2管理员
	Remember bool   `json:"rememberMe" binding:"required"` //记住我
}

// AuthByPassword 通过密码认证
func AuthByCombo(c *gin.Context) {
	var postForm passwordLoginForm
	err := c.ShouldBindJSON(&postForm) //验证数据完整性
	if err != nil {
		c.Error(exception.ApiParamError)
		return
	}
	fmt.Println("登录信息:", postForm)

	//用数据类型转换判断是编号登录还是邮箱登录
	var user interface{}
	var userErr error
	//matched, _ := regexp.MatchString(`^\d+$`, postForm.Account) 正则,已弃用
	_, err = strconv.ParseUint(postForm.Account, 10, 64)
	if err != nil {
		// Convert id string to uint64

		// if err != nil {
		// 	c.Error(exception.ApiParamError)
		// 	return
		// }

		// WARN(MUCHEXD) 取消姓名登录
		// 显而易见，姓名说可以重复的  改成邮箱登录

		fmt.Println("邮箱登录:", postForm.Account)
		user, userErr = userService.GetUserByEmail(postForm.Account) //从数据库获取用户信息,判断用户存在
	} else {

		// WARN(MUCHEXD) 命名不合理
		// 根据你的代码，"编号"指的是student_id，请在代码里使用更明确的命名

		fmt.Println("编号登录:", postForm.Account)
		user, userErr = userService.GetUserByNum(postForm.Account) //从数据库获取用户信息,判断用户存在
	}
	if errors.Is(userErr, gorm.ErrRecordNotFound) {
		c.Error(exception.UsrNotExisted)
		return
	}
	if userErr != nil {
		c.Error(exception.SysUknExc)
		return
	}

	accountInfo, ok := user.(*models.AccountInfo)
	if !ok {
		c.Error(exception.SysUknExc)
		return
	}

	if err := userService.VerifyPwd(accountInfo, postForm.Password); err != nil { //验证密码
		var apiErr *exception.Exception
		if errors.As(err, &apiErr) {
			fmt.Println("密码错误0:", err)
			c.Error(exception.UsrPasswordErr)
		} else {
			fmt.Println("密码错误1:", err)
			c.Error(exception.SysCannotLoadFromDB)
		}
		return
	}
	//TODO:解决秘钥签名错误的问题
	utils.JsonSuccessResponse(c, "登录成功", map[string]interface{}{
		"token":    webtoken.GenerateWt(accountInfo.ID, accountInfo.PermGroupID, 100000000),
		"userID":   accountInfo.ID,
		"username": accountInfo.UserName,
		"userType": strconv.FormatUint(uint64(accountInfo.PermGroupID), 10),
	})
}
