package middleware

import (
	configreader "JHETBackend/internal/configs/configReader"
	"errors"
	"log"
	"net/http"
	"strconv"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

//TODO: error codes and naming (MucheXD/09.18)

// #####CONST#####
// disable only in dev or doing a demo for ZHEYI
const enableFailFast = true

const errCodeCfgPath = "../configs"
const errCodeCfgName = "errorCodes"

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

func NewBizExc(errCode int) BusinessException {
	configreader.GetConfig()
	return BusinessException{errCode, findMsgByCode(errCode)}
}

// 统一错误处理中间件：既收 c.Error 也收 panic
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
					log.Printf("[INFO][ErrMidware] 发生预期业务的错误: %v", err)
					c.JSON(http.StatusOK, ExceptionResponce{
						Code: bizexc.Code, Msg: bizexc.Msg})
				} else { //未知错误
					log.Printf("[WARN][ErrMidware] 发生未知的业务错误: %v", err)
					c.JSON(http.StatusOK, ExceptionResponce{
						Code: 50001, Msg: "UknBizError occurred"}) //TODO: 50001 错误码不规范
				}
			}
			c.Abort()
		}
	}
}

//#####PRIVATE#####

var initReadCodesOnce sync.Once

// 错误码对应表
var errCodeMap map[int]string

// 读入errCodes配置
func initReadCodes() {
	log.Println("[INFO][ErrMidware] 载入错误码配置")
	v := viper.New()
	v.AddConfigPath(errCodeCfgPath)
	v.SetConfigName(errCodeCfgName) // 不带扩展名
	v.SetConfigType("yaml")

	if err := v.ReadInConfig(); err != nil {
		log.Fatalf("[FATAL][ErrMidware] 无法加载错误码配置 错误: %v", err)
	}

	raw := v.AllSettings()
	errCodeMap = make(map[int]string, len(raw))
	for k, val := range raw {
		code, err := strconv.Atoi(k)
		if err != nil {
			log.Println("[WARN][ErrMidware] 错误码配置中含有非法键，跳过。")
			continue //忽略非数字键
		}
		if msg, ok := val.(string); ok {
			errCodeMap[code] = msg
		} else {
			log.Println("[WARN][ErrMidware] 错误码配置中含有非法值，跳过。")
			continue
		}
	}
}

func findMsgByCode(errCode int) string {
	initReadCodesOnce.Do(initReadCodes) //懒加载错误码配置
	if msg := errCodeMap[errCode]; msg != "" {
		return msg
	}
	log.Printf("[WARN][ErrMidware] 尝试调用不存在的错误码：%d", errCode)
	return "Unknown Error"
}

func panicErrHandler(c *gin.Context) {
	var isUnexpectedPanic = false
	if rec := recover(); rec != nil {
		// 如果是业务主动 panic 的 bizexp
		if bizexc, ok := rec.(*BusinessException); ok {
			log.Printf("[ERROR][ErrMidware] 发生预期的异常: %v", rec)
			c.JSON(http.StatusOK, ExceptionResponce{
				Code: bizexc.Code, Msg: bizexc.Msg})
			c.Abort()
			return
		}
		// 其它未知 panic
		log.Printf("[FATAL][ErrMidware] 发生未知的异常: %v", rec)
		c.JSON(http.StatusOK, ExceptionResponce{
			Code: 50001, Msg: "UknExc happened with panic"}) //TODO: 50001 错误码不规范
		c.Abort()
		isUnexpectedPanic = true
	}
	if isUnexpectedPanic && enableFailFast {
		log.Panic("[PANIC] Panic found in errorHandler and Fail-Fast enabled.")
	}
}
