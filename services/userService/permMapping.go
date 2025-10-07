package userService

import (
	"JHETBackend/common/exception"
	"fmt"
	"github.com/gin-gonic/gin"
)

func MapPerm(c *gin.Context) uint32 {
	//映射权限组id
	var permGroupID uint32
	userType := c.PostForm("userType")
	if userType == "" {
		permGroupID = 1
	} else {
		switch userType {
		case "user":
			permGroupID = 37
		case "admin":
			permGroupID = 38
		case "superadmin":
			permGroupID = 35
		default:
			c.Error(exception.ApiParamError)
			fmt.Println("用户类型错误:", userType)
			permGroupID = 1
		}
	}
	return permGroupID
}
