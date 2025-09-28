package AccountController

import (
	"JHETBackend/common/exception"
	"JHETBackend/models"
	"JHETBackend/services/userService"

	"github.com/gin-gonic/gin"
)

// 获取当前用户的所有信息
func GetAccountInfo(c *gin.Context) models.AccountInfo {
	accountID, err := GetAccountIDFromContext(c)
	if err != nil {
		c.Error(err)
		return models.AccountInfo{}
	}
	accountInfo, err := userService.GetAccountInfoByUID(accountID)
	if err != nil {
		c.Error(err)
		return models.AccountInfo{}
	}
	return *accountInfo
}

// 从 gin 的上下文中取 AccountID
// 其它 Controller 也会用到这个函数
func GetAccountIDFromContext(c *gin.Context) (uint64, error) {
	accountIDObj, ok := c.Get("AccountID")
	if !ok { // 用户id不存在，视为未登录
		return 0, exception.UsrNotLogin
	}
	accountID, ok := accountIDObj.(uint64)
	if !ok { // 用户id不合法，视为未登录
		return 0, exception.UsrNotLogin
	}
	return accountID, nil
}
