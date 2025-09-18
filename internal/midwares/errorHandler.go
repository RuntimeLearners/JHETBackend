package middleware

import (
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 业务错误类型
type ErrorInfo struct {
	Code int
	Msg  string
}

func (e *ErrorInfo) Error() string {
	return e.Msg
}

// 统一响应结构
type Responce struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

// 单中间件：既收 c.Error 也收 panic
func UnifiedErrorMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		// 捕获并处理panic错误
		defer panicErrHandler(c)

		//执行后续中间件与业务代码
		c.Next()

		// 处理 c.Error 收集到的错误，如果panic此处会跳过
		if len(c.Errors) > 0 {
			err := c.Errors.Last().Err
			if err != nil {
				var errInfo *ErrorInfo
				//预定义的错误
				if ok := errors.As(err, &errInfo); ok {
					c.JSON(http.StatusOK, ErrorInfo{Code: errInfo.Code, Msg: errInfo.Msg})
				} else { //未知错误
					log.Printf("[ERROR][ErrMidware]Unknow error occurred.")
					c.JSON(http.StatusOK, ErrorInfo{Code: 50001, Msg: "未知错误发生"}) //TODO: 50001 错误码不规范
				}
			}
			c.Abort()
		}
	}
}

func panicErrHandler(c *gin.Context) {
	if rec := recover(); rec != nil {
		// 1.1 如果是业务主动 panic 的 BizError
		if biz, ok := rec.(*ErrorInfo); ok {
			c.JSON(http.StatusOK, Responce{Code: biz.Code, Msg: biz.Msg})
			c.Abort()
			return
		}
		// 1.2 其它未知 panic
		log.Printf("[FATAL] %v", rec)
		c.JSON(http.StatusOK, Responce{Code: 50001, Msg: "未知崩溃，联系管理员"}) //TODO: 50001 错误码不规范
		c.Abort()
	}
}
