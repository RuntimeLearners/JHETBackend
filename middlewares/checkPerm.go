package middleware

import (
	"JHETBackend/common/exception"
	"JHETBackend/common/permission"
	"log"

	"github.com/gin-gonic/gin"
)

func NeedPerm(needed ...permission.PermissionID) gin.HandlerFunc {
	return func(c *gin.Context) {
		var pgid = uint32(0)
		if rawPGID, ok := c.Get("PermissionGroupID"); ok {
			if pgid, ok = rawPGID.(uint32); !ok {
				log.Println("[ERROR][CHKPERM] PermissionGroupID 类型错误")
			}
		}
		if pgid == 0 {
			//c.Redirect(401, "/login") 感觉放进错误处理中间件?
			c.Error(exception.UsrNotLogin)
			c.AbortWithStatus(401) // 直接返回401
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
