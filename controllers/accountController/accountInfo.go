package AccountController

import (
	"JHETBackend/common/exception"
	"JHETBackend/models"
	"JHETBackend/services/userService"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 获取当前用户的所有信息
func GetAccountInfo(c *gin.Context) models.AccountInfo {
	var err error
	var accountID uint64
	accountID, err = GetAccountIDFromContext(c)
	if err != nil {
		c.Error(err)
		return models.AccountInfo{}
	}

	//这一步是留给查别人的信息用的
	userIDStr := c.DefaultQuery("id", string(accountID)) //如果不存在query, 返回自己的id
	userID, err := strconv.ParseUint(userIDStr, 10, 64)  //尝试转换字符串->uint64
	if err != nil {
		c.Error(err)
		c.Error(exception.ApiParamError)
		return models.AccountInfo{}
	}

	accountInfo, err := userService.GetAccountInfoByUID(userID)
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
