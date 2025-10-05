package accountControllers

import (
	"JHETBackend/common/exception"
	"JHETBackend/services/userService"
	"JHETBackend/models"
	"JHETBackend/utils"
	"errors"
	"strconv"
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
			utils.JsonErrorResponse(c, int(mysqlErr.Number), fmt.Sprintf("更新失败: %s", mysqlErr.Message))
		}
	} else {
		utils.JsonSuccessResponse(c, "更新成功", nil)
		return
	}
}

func UpdateAccountInfoAdmin(c *gin.Context) {
	// 管理员更改用户信息
	var accountID uint64
	var err error
	accountIDStr := c.Param("id")
	if accountIDStr == "" { //如果没有传入id参数, 用qurey时有用, 用path时空直接返回404
		c.Error(exception.ApiParamError)
		fmt.Println("err1")
		//utils.JsonSuccessResponse(c, "查询失败", models.AccountInfo{})
		return
	}
	accountID, err = strconv.ParseUint(accountIDStr, 10, 64) //将传入str id转换为uint, 小转换就不写service了
	if err != nil {
		fmt.Println("err1", err)
		c.Error(exception.ApiParamError)
		return
	}

	var postForm *models.AccountInfo
	// err = c.ShouldBindJSON(&postForm)
	// if err != nil {
	// 	c.Error(exception.ApiParamError)
	// 	fmt.Println("参数错误0:", err)
	// 	return
	// }
	fmt.Println(accountID, "修改信息:", postForm)
	//updateForm.permGroupID := userService.MapPerm(c) //映射权限组id

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
			//c.Error(exception.ApiParamError)
			utils.JsonErrorResponse(c, int(mysqlErr.Number), fmt.Sprintf("更新失败: %s", mysqlErr.Message))
		}
	} else {
		utils.JsonSuccessResponse(c, "更新成功", nil)
		return
	}
}