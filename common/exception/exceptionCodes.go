package exception

var (
	VeryGood              = NewException(0000, "啥事都煤油花生!")
	TestIntendedException = NewException(9000, "测试错误")

	UsrNotLogin           = NewException(1001, "用户未登录")
	UsrNotPermitted       = NewException(1002, "用户无此权限")
	UsrNotExisted         = NewException(1003, "用户不存在")
	UsrAlreadyExisted     = NewException(1004, "用户已存在")   //邮箱或者人员编号(学生ID)重复
	UsrInfoAlreadyExisted = NewException(1005, "用户信息已存在") //更新信息时邮箱重复
	UsrPasswordErr        = NewException(1006, "用户密码错误")  //在注册(合规性验证),登录,修改密码 时均使用此错误码
	UsrLoginInvalid       = NewException(1007, "用户登录无效")
	UsrEmailErr           = NewException(1008, "用户邮箱格式错误")
	UsrEmailNotVerified   = NewException(1009, "用户邮箱未验证")
	Usr2FAErr             = NewException(1010, "用户2FA认证错误")

	FbPostDataInvalid       = NewException(2001, "传入的反馈帖数据无效")
	FbReplyPostNotFound     = NewException(2002, "回复指向的原帖无效")
	FbReplyNestTooDeep      = NewException(2003, "回复嵌套过深")
	FbPostAttachmentInvalid = NewException(2004, "反馈帖附件无效")
	FbPostNotFound          = NewException(2005, "反馈帖不存在")
	FbPostDeleted           = NewException(2006, "反馈帖已被删除")
	FbPostUpdateFailed      = NewException(2007, "无法更新反馈帖")
	FbNotSpamDontChecked    = NewException(2008, "反馈帖未被标记为垃圾信息，无需审核")
	FbRatingUnslovedPost    = NewException(2009, "无法评价未解决的反馈帖")

	ApiNoFormFile         = NewException(4001, "无文件字段")
	ApiFileTooLarge       = NewException(4002, "上传文件过大")
	ApiFileNotSupported   = NewException(4003, "拒绝上传此类型文件类型")
	ApiParamError         = NewException(4004, "参数错误")
	ApiFileCannotOpen     = NewException(4005, "无法打开上传的文件")
	ApiFileNotSaved       = NewException(4006, "无法保存上传的文件")
	ApiFeedbackNotCreated = NewException(2002, "无法创建回复")

	SysUknExc              = NewException(5000, "未知错误")
	SysCannotLoadFromDB    = NewException(5001, "内部异常: 加载数据库时出错")
	SysCannotLoadPermGroup = NewException(5002, "内部异常: 无法从数据库读取权限表")
	SysPwdHashFailed       = NewException(5003, "内部异常: 密码加密失败") //暂且留着
	SysCannotUpdate        = NewException(5004, "内部异常: 无法更新数据库")
	SysCannotReadDB        = NewException(5005, "内部异常: 无法读取数据库")

	FileCannotSaveUploaded = NewException(6001, "文件系统错误: 无法保存上传的文件")

	// "code":    200,	"message": "success"
	// Code: 9001, Msg: "UknBizError occurred"
)
