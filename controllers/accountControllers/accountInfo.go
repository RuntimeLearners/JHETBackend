package accountControllers

import (
	"JHETBackend/common/exception"
	"JHETBackend/models"
	"JHETBackend/services/userService"
	"JHETBackend/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 管理员获取用户信息
func GetAccountInfoAdmin(c *gin.Context) {
	confidentiality := false //保密信息--否
	if c.Query("id") == "" { //如果没有传入id参数
		c.Error(exception.ApiParamError)
		return models.AccountInfo{}
	}
	accountInfo := GetAccountInfo(c, c.Query("id"), confidentiality)
	utils.JsonSuccessResponse(c,"查询成功",accountInfo)
}

// 普通用户获取信息
func GetAccountInfoUser(c *gin.Context){
	var accountID uint64
	var err error
	accountID, err = GetAccountIDFromContext(c)
	if err != nil {
		c.Error(err)
		return models.AccountInfo{}
	}
	accountIDStr := strconv.FormatUint(accountID, 10)
	if c.Query("id") == "" || c.Query("id") == accountIDStr { //如果没有传入id参数或传入自己的id
		confidentiality := false //保密信息--否
		accountInfo := GetAccountInfo(c, accountIDStr, confidentiality)
	}
	//传入了别人的id
	confidentiality := true //保密信息--是
	accountInfo := GetAccountInfo(c, c.Query("id"), confidentiality)
	utils.JsonSuccessResponse(c,"查询成功",accountInfo)
}

// 获取用户的所有信息
func GetAccountInfo(c *gin.Context, userIDStr string, confidentiality bool) models.AccountInfo {
	var err error
	userID, err := strconv.ParseUint(userIDStr, 10, 64) //将传入str id转换为uint
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

	result := *accountInfo

	//勿cue 有空再整理成函数
	//始终屏蔽的信息
	if result.PasswordHash != "" {
		result.PasswordHash = "have"
	} else {
		result.PasswordHash = "don't have"
	}
	if result.TwoFactorAuth != "" {
		result.TwoFactorAuth = "have"
	} else {
		result.TwoFactorAuth = "don't have"
	}
	if result.WechatOpenID != "" {
		result.WechatOpenID = "have"
	} else {
		result.WechatOpenID = "don't have"
	}

	if confidentiality { //保密信息
		result.ID = 0
		result.RealName = "保密"
		result.PermGroupID = 0
		result.Email = "保密"
		result.PhoneNumber = "保密"
		result.Major = "保密"
		result.PhoneNumber = "保密"
		result.StudentID = "保密"
		result.Department = "保密"
		result.Grade = "保密"
		result.CreatedAt = result.CreatedAt.Truncate(0)
		result.UpdatedAt = result.UpdatedAt.Truncate(0)
	}

	return result
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
