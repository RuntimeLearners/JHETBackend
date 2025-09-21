package middleware

import (
	"JHETBackend/common/exception"
	"JHETBackend/common/permission"

	"github.com/gin-gonic/gin"
)

func NeedPerm(needed ...permission.PermissionID) gin.HandlerFunc {
	return func(c *gin.Context) {
		pgid := uint32(c.GetUint("PermissionGroupID"))
		if pgid == 0 {
			c.Error(exception.UsrNotLogin)
			return
		}
		if !permission.IsPermSatisfied(pgid, needed...) {
			c.Error(exception.UsrNotPermitted)
			c.Abort()
			return
		}
		c.Next()
	}
}
