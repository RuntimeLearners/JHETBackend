package exception

var (
	VeryGood              = NewException(0000, "啥事都煤油花生!")
	TestIntendedException = NewException(9000, "测试错误")

	UsrNotLogin       = NewException(1001, "用户未登录")
	UsrNotPermitted   = NewException(1002, "用户无此权限")
	UsrNotExisted     = NewException(1003, "用户不存在")
	UsrAlreadyExisted = NewException(1004, "用户已存在") //邮箱或者人员编号(学生ID)重复
	UsrPasswordErr    = NewException(1005, "用户密码错误")
	UsrLoginInvalid   = NewException(1006, "用户登录无效")

	ApiNoFormFile       = NewException(4001, "无文件字段")
	ApiFileTooLarge     = NewException(4002, "上传文件过大")
	ApiFileNotSupported = NewException(4003, "拒绝上传此类型文件类型")
	ApiParamError       = NewException(4004, "参数错误")
	ApiFileCannotOpen   = NewException(4005, "无法打开上传的文件")
	ApiFileNotSaved     = NewException(4006, "无法保存上传的文件")

	SysUknExc              = NewException(5000, "未知错误")
	SysCannotLoadFromDB    = NewException(5001, "内部异常: 加载数据库时出错")
	SysCannotLoadPermGroup = NewException(5002, "内部异常: 无法从数据库读取权限表")
	SysPwdHashFailed       = NewException(5003, "内部异常: 密码加密失败")  //暂且留着
	SysCannotUpdate        = NewException(5004, "内部异常: 无法更新数据库") //暂且留着
	SysCannotReadDB        = NewException(5005, "内部异常: 无法读取数据库")

	FileCannotSaveUploaded = NewException(6001, "文件系统错误: 无法保存上传的文件")
)
