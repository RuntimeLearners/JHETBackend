package loginControllers

import (
	"JHETBackend/common/exception"
	"JHETBackend/models"
	"JHETBackend/services/userService"
	"errors"
	"regexp"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type passwordLoginForm struct {
	Account  string `json:"account" binding:"required"` //返回的姓名或id
	Password string `json:"password" binding:"required"`
	Remenber bool   `json:"rememberMe" binding:"required"` //记住我
}

// AuthByPassword 通过密码认证
func AuthByCombo(c *gin.Context) {
	var postForm passwordLoginForm
	err := c.ShouldBindJSON(&postForm) //验证数据完整性
	if err != nil {
		c.Error(exception.ApiParamError)
		return
	}

	//用正则判断是id登录还是姓名登录
	var user interface{}
	var userErr error
	matched, _ := regexp.MatchString(`^\d+$`, postForm.Account)
	if matched {
		// Convert id string to uint64
		var userID uint64
		userID, err := strconv.ParseUint(postForm.Account, 10, 64)
		if err != nil {
			c.Error(exception.ApiParamError)
			return
		}
		user, userErr = userService.GetUserByID(userID) //从数据库获取用户信息,判断用户存在
	} else {
		//fmt.Println("姓名登录:", postForm.Account)
		user, userErr = userService.GetUserByName(postForm.Account) //从数据库获取用户信息,判断用户存在
	}
	if errors.Is(userErr, gorm.ErrRecordNotFound) {
		c.Error(exception.UsrNotExisted)
		return
	}
	if err != nil {
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
			//exception.NewException(c, apiErr, err)
		} else {
			//apiException.AbortWithException(c, apiException.ServerError, err)
		}
		return
	}
}
