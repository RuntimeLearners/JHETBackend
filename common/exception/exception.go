package exception

type Exception struct {
	Code int
	Msg  string
	Data interface{}
}

func (e *Exception) Error() string {
	return e.Msg
}

func NewException(code int, msg string) *Exception {
	return &Exception{
		Code: code,
		Msg:  msg,
	}
}

func (e *Exception) NewWithData(code int, msg string, data interface{}) *Exception {
	return &Exception{
		Code: code,
		Msg:  msg,
		Data: data,
	}
}

//食用方式

// 动态创建异常
// err := exception.NewException(9999, "动态生成的错误")

// 动态创建带数据的异常
// err := (&exception.Exception{}).NewWithData(8888, "带数据的错误", map[string]interface{}{
//     "key": "value",
//     "info": "额外的错误信息",
// })

// 历史遗留代码 舍不得删

// func FindMsgByCode(errCode int) string {
// 	initReadCodesOnce.Do(initReadCodes) //懒加载错误码配置
// 	if msg := errCodeMap[errCode]; msg != "" {
// 		return msg
// 	}
// 	log.Printf("[WARN][ErrMidware] 尝试调用不存在的错误码：%d", errCode)
// 	return "Unknown Error"
// }

// var initReadCodesOnce sync.Once

// // 错误码对应表
// var errCodeMap map[int]string

// // 读入errCodes配置
// func initReadCodes() {
// 	log.Println("[INFO][ErrMidware] 载入错误码配置")
// 	v := viper.New()
// 	v.AddConfigPath(errCodeCfgPath)
// 	v.SetConfigName(errCodeCfgName) // 不带扩展名
// 	v.SetConfigType("yaml")

// 	if err := v.ReadInConfig(); err != nil {
// 		log.Fatalf("[FATAL][ErrMidware] 无法加载错误码配置 错误: %v", err)
// 	}

// 	raw := v.AllSettings()
// 	errCodeMap = make(map[int]string, len(raw))
// 	for k, val := range raw {
// 		code, err := strconv.Atoi(k)
// 		if err != nil {
// 			log.Println("[WARN][ErrMidware] 错误码配置中含有非法键，跳过。")
// 			continue //忽略非数字键
// 		}
// 		if msg, ok := val.(string); ok {
// 			errCodeMap[code] = msg
// 		} else {
// 			log.Println("[WARN][ErrMidware] 错误码配置中含有非法值，跳过。")
// 			continue
// 		}
// 	}
// }
