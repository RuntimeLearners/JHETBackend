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
			permGroupID = 1
		case "admin":
			permGroupID = 2
		case "superadmin":
			permGroupID = 3
		default:
			c.Error(exception.ApiParamError)
			fmt.Println("用户类型错误:", userType)
			permGroupID = 1
		}
	}
	return permGroupID
}
