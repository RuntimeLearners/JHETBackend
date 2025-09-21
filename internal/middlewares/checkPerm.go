package middleware

import (
	"JHETBackend/internal/common/permission"

	"github.com/gin-gonic/gin"
)

func NeedPerm(needed ...permission.PermissionID) gin.HandlerFunc {
	return func(c *gin.Context) {
		pgid := uint32(c.GetUint("PermissionGroupID"))
		if pgid == 0 {
			c.Error(NewBizExc(1001))
			return
		}
		if !permission.IsPermSatisfied(pgid, needed...) {
			c.Error(NewBizExc(1002))
			c.Abort()
			return
		}
		c.Next()
	}
}
