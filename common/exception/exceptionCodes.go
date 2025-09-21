package exception

var (
	TestIntendedException = NewException(9000, "测试错误")
	UsrNotLogin           = NewException(1001, "用户未登录")
	UsrNotPermitted       = NewException(1002, "用户无此权限")

	ApiNoFormFile       = NewException(4001, "无文件字段")
	ApiFileTooLarge     = NewException(4002, "上传文件过大")
	ApiFileNotSupported = NewException(4003, "拒绝上传此类型文件类型")

	SysUknExc              = NewException(5000, "未知错误")
	SysCannotLoadPermGroup = NewException(5001, "内部异常: 无法从数据库读取权限表")

	FileCannotSaveUploaded = NewException(6001, "文件系统错误: 无法保存上传的文件")
)
