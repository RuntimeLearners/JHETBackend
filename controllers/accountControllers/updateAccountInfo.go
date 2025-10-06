package accountControllers

import (
	"JHETBackend/common/exception"
	"JHETBackend/services/userService"
	"errors"
	"fmt"
	"strconv"

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

type AdminUpdateUserInfo struct {
	Name        *string `json:"name"`        // 真实姓名
	UserName    *string `json:"username"`    // 用户名
	Email       *string `json:"email"`       // 邮箱
	Department  *string `json:"department"`  // 部门/院系
	StudentID   *string `json:"studentId"`   // 学号
	Major       *string `json:"major"`       // 专业
	PhoneNumber *string `json:"phoneNumber"` // 手机号
	Grade       *string `json:"grade"`       // 年级
	AvatarFile  *string `json:"avatar"`      // 头像文件
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
			//c.Set("data", postForm)
			e := exception.NewException(int(mysqlErr.Number), fmt.Sprintf("更新失败: %s", mysqlErr.Message))
			c.Error(e)
			//utils.JsonErrorResponse(c, int(mysqlErr.Number), fmt.Sprintf("更新失败: %s", mysqlErr.Message))
		}
	} else {
		//utils.JsonSuccessResponse(c, "更新成功", nil)
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
		return
	}
	accountID, err = strconv.ParseUint(accountIDStr, 10, 64) //将传入str id转换为uint, 小转换就不写service了
	if err != nil {
		fmt.Println("err1", err)
		c.Error(exception.ApiParamError)
		return
	}

	var postForm AdminUpdateUserInfo
	err = c.ShouldBindJSON(&postForm)
	if err != nil {
		c.Error(exception.ApiParamError)
		fmt.Println("参数错误0:", err)
		return
	}

	// 处理"del"字符串，将其转换为空值
	updateData := processDelValues(postForm)

	fmt.Println(accountID, "修改信息:", updateData)

	err = userService.UpdateUser(
		accountID,
		updateData,
	)
	if err != nil {
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr) && mysqlErr.Number == 1062 { // 1062 是 MySQL 的重复键错误代码
			c.Error(exception.UsrInfoAlreadyExisted)
			return
		} else {
			//1406 数据过长
			e := exception.NewException(int(mysqlErr.Number), fmt.Sprintf("更新失败: %s", mysqlErr.Message))
			c.Error(e)
			//utils.JsonErrorResponse(c, int(mysqlErr.Number), fmt.Sprintf("更新失败: %s", mysqlErr.Message))
		}
	} else {
		return
	}
}

// 你没看错,就是这么朴素
func processDelValues(form AdminUpdateUserInfo) map[string]interface{} {
	updateData := make(map[string]interface{})

	//真名也不可删除
	if form.Name != nil {
		updateData["realname"] = *form.Name
	}

	if form.UserName != nil {
		if *form.UserName == "del" {
			updateData["username"] = ""
		} else {
			updateData["username"] = *form.UserName
		}
	}

	//邮箱不可删除
	if form.Email != nil {
		updateData["email"] = *form.Email
	}

	if form.Department != nil {
		if *form.Department == "del" {
			updateData["department"] = ""
		} else {
			updateData["department"] = *form.Department
		}
	}

	//用户编号不可删除
	if form.StudentID != nil {
		updateData["student_id"] = *form.StudentID
	}

	if form.Major != nil {
		if *form.Major == "del" {
			updateData["major"] = ""
		} else {
			updateData["major"] = *form.Major
		}
	}

	if form.PhoneNumber != nil {
		if *form.PhoneNumber == "del" {
			updateData["phone_number"] = ""
		} else {
			updateData["phone_number"] = *form.PhoneNumber
		}
	}

	if form.Grade != nil {
		if *form.Grade == "del" {
			updateData["grade"] = ""
		} else {
			updateData["grade"] = *form.Grade
		}
	}

	if form.AvatarFile != nil {
		if *form.AvatarFile == "del" {
			updateData["avatar_fileuuid"] = ""
		} else {
			updateData["avatar_fileuuid"] = *form.AvatarFile
		}
	}

	return updateData
}
