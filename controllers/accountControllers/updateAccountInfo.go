package accountControllers

import (
	"JHETBackend/common/exception"
	"JHETBackend/services/userService"
	"JHETBackend/utils"
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
)

type UserInfo struct {
	Email       string `json:"email"`
	UserName    string `json:"username" gorm:"column:username"`
	Major       string `json:"major"`       //专业
	PhoneNumber string `json:"phoneNumber"` //手机号
	Department  string `json:"department"`  //部门/院系 学生和管理员均有此项
	Grade       string `json:"grade"`       //年级 F:2025
}

func UpdateAccountInfoUser(c *gin.Context) {
	// 普通用户更改自己的信息
	var accountID uint64
	var err error
	accountID, err = GetAccountIDFromContext(c)
	if err != nil {
		c.Error(exception.ApiParamError)
		return
	}

	var postForm UserInfo
	err = c.ShouldBindJSON(&postForm)
	if err != nil {
		c.Error(exception.ApiParamError)
		fmt.Println("参数错误0:", err)
		return
	}
	fmt.Println(accountID, "修改信息:", postForm)

	err = userService.UpdateUser(
		accountID,
		postForm,
	)
	if err != nil {
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr) && mysqlErr.Number == 1062 { // 1062 是 MySQL 的重复键错误代码
			//fmt.Println("信息存在:", err)
			c.Error(exception.UsrInfoAlreadyExisted)
			//utils.JsonSuccessResponse(c, "更新失败", nil)
			return
		} else {
			//1406 数据过长
			utils.JsonErrorResponse(c, int(mysqlErr.Number), fmt.Sprintf("更新失败: %v", mysqlErr.Message))
		}
	} else {
		utils.JsonSuccessResponse(c, "更新成功", nil)
		return
	}
}
