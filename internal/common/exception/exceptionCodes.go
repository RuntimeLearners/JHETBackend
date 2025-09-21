package exception

var (
	TestIntendedException = NewException(9000, "测试错误")
	UsrNotLogin           = NewException(1001, "用户未登录")
	UsrNotPermitted       = NewException(1002, "用户无此权限")

	SysUknExc              = NewException(5000, "未知错误")
	SysCannotLoadPermGroup = NewException(5001, "内部异常: 无法从数据库读取权限表")
)
