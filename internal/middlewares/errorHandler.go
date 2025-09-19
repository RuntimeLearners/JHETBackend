package middleware

import (
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

//TODO: error codes and naming (MucheXD/09.18)

// #####CONST#####
// disable only in dev or doing a demo for ZHEYI
const enableFailFast = true

//#####PUBLIC#####

// 业务预期错误类型(实现Error()string，是一个标准的错误)
type BusinessException struct {
	Code int
	Msg  string
}

func (e *BusinessException) Error() string {
	return e.Msg
}

// 统一响应结构
type ExceptionResponce struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

// 单中间件：既收 c.Error 也收 panic
func UnifiedErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {

		// 捕获并处理panic错误
		defer panicErrHandler(c)

		//执行后续中间件与业务代码
		c.Next()

		// 处理 c.Error 收集到的错误，如果panic此处会跳过
		if len(c.Errors) > 0 {
			err := c.Errors.Last().Err // TODO: 似乎只会弹出最后一个错误
			if err != nil {
				var bizexc *BusinessException
				//预定义的错误
				if ok := errors.As(err, &bizexc); ok { //预期错误
					log.Printf("[INFO][ErrMidware]BizErr happened but expected: %v", err)
					c.JSON(http.StatusOK, ExceptionResponce{
						Code: bizexc.Code, Msg: bizexc.Msg})
				} else { //未知错误
					log.Printf("[WARN][ErrMidware]UknBizError occurred: %v", err)
					c.JSON(http.StatusOK, ExceptionResponce{
						Code: 50001, Msg: "UknBizError occurred"}) //TODO: 50001 错误码不规范
				}
			}
			c.Abort()
		}
	}
}

//#####PRIVATE#####

func panicErrHandler(c *gin.Context) {
	var isUnexpectedPanic = false
	if rec := recover(); rec != nil {
		// 如果是业务主动 panic 的 bizexp
		if bizexc, ok := rec.(*BusinessException); ok {
			log.Printf("[ERROR][ErrMidware] BizExc happened with panic: %v", rec)
			c.JSON(http.StatusOK, ExceptionResponce{
				Code: bizexc.Code, Msg: bizexc.Msg})
			c.Abort()
			return
		}
		// 其它未知 panic
		log.Printf("[FATAL][ErrMidware] UknExc happened with panic: %v", rec)
		c.JSON(http.StatusOK, ExceptionResponce{
			Code: 50001, Msg: "UknExc happened with panic"}) //TODO: 50001 错误码不规范
		c.Abort()
		isUnexpectedPanic = true
	}
	if isUnexpectedPanic && enableFailFast {
		log.Panic("[PANIC] Panic found in errorHandler and Fail-Fast enabled.")
	}
}
