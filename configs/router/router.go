package router

import (
	"JHETBackend/common/permission"
	middleware "JHETBackend/middlewares"
	"fmt"

	"JHETBackend/controllers/accountControllers"
	feedbackcontrollers "JHETBackend/controllers/feedbackControllers"
	filecontroller "JHETBackend/controllers/fileController"
	"JHETBackend/controllers/loginControllers"
	"JHETBackend/controllers/registerControllers"

	"github.com/gin-gonic/gin"
	//"github.com/silenceper/wechat/v2/openplatform/account"
)

func SayHello(c *gin.Context) {
	// 200 表示 HTTP 响应状态码（<=> http.StatusOK）
	// 使用 Context 的 String 函数将 "Hello 精弘!" 这句话以纯文本（字符串）的形式返回给前端
	// 实际上是对返回响应的封装
	c.String(200, "Hello go!")
}

func InitEngine() *gin.Engine {
	ginEngine := gin.Default()

	fmt.Println(gin.Context{})
	// // 添加中间件处理字符编码
	// ginEngine.Use(func(c *gin.Context) {
	// 	c.Header("Content-Type", "application/json; charset=utf-8")
	// 	c.Next()
	// })

	ginEngine.GET("/test", middleware.UnifiedErrorHandler(),
		middleware.Auth,
		middleware.NeedPerm(
			permission.Perm_ForTestOnly1,
			permission.Perm_ForTestOnly2), SayHello)

	//登录注册这一块
	ginEngine.POST("/api/auth/login/combo", middleware.UnifiedErrorHandler(), loginControllers.AuthByCombo)
	ginEngine.GET("/api/auth/login/combo", middleware.UnifiedErrorHandler(), SayHello)

	ginEngine.POST("/api/auth/register", middleware.UnifiedErrorHandler(), registerControllers.CreateUserUser)

	ginEngine.POST("/api/user/avatar", middleware.UnifiedErrorHandler(), middleware.Auth, filecontroller.UpdateAvatar)

	//帖子这一块
	//上传图片这一块
	ginEngine.POST("/api/upload/image", middleware.UnifiedErrorHandler(), middleware.Auth,
		filecontroller.UpdateAvatar)
	//发帖这一块
	ginEngine.POST("/api/feedback", middleware.UnifiedErrorHandler(),
		middleware.Auth,
		feedbackcontrollers.CreateFeedbackPost)
	//查询帖子这一块
	ginEngine.GET("/api/feedback/:id", middleware.UnifiedErrorHandler(),
		middleware.Auth,
		feedbackcontrollers.GetAllFeedbackPosts) //由于给管理员配置函数麻烦,故先写死

	//用户信息这一块
	//无需权限 测试用
	ginEngine.GET("/api/user/info/", middleware.UnifiedErrorHandler(),
		middleware.Auth,
		accountControllers.GetAccountInfoUser)
	ginEngine.GET("/api/admin/users/", middleware.UnifiedErrorHandler(),
		middleware.Auth,
		accountControllers.GetAccountInfoAdmin)
	// 修改用户信息
	ginEngine.PUT("/api/user/info/", middleware.UnifiedErrorHandler(),
		middleware.Auth,
		accountControllers.UpdateAccountInfoUser)
	// 修改密码
	ginEngine.PUT("/api/user/pwd/", middleware.UnifiedErrorHandler(),
		middleware.Auth,
		accountControllers.ChangePassword)

	//超管
	// 新增用户
	ginEngine.POST("/api/admin/users", middleware.UnifiedErrorHandler(),
		middleware.Auth,
		registerControllers.CreateUserAdmin)
	//删除用户
	ginEngine.DELETE("/api/admin/users/:id", middleware.UnifiedErrorHandler(),
		middleware.Auth,
		accountControllers.DeleteAccount)
	//修改用户信息
	ginEngine.PUT("/api/admin/users/:id", middleware.UnifiedErrorHandler(),
		middleware.Auth,
		accountControllers.UpdateAccountInfoAdmin)

	//通用路由
	// // 修改密码
	// ginEngine.PUT("/api/user/pwd/", middleware.UnifiedErrorHandler(),
	// 	middleware.Auth,
	// 	accountControllers.ChangePassword)
	// // 更改头像
	// // ginEngine.POST("/api/user/avatar", middleware.UnifiedErrorHandler(), middleware.Auth,
	// middleware.NeedPerm(permission.Perm_UpdateAvatar),
	// filecontroller.UpdateAvatar)

	// //帖子这一块
	// //上传图片这一块
	// ginEngine.POST("/api/upload/image", middleware.UnifiedErrorHandler(), middleware.Auth,
	// 	middleware.NeedPerm(permission.Perm_UploadImage),
	// 	filecontroller.UpdateAvatar)
	// //发帖这一块
	// ginEngine.POST("/api/feedback", middleware.UnifiedErrorHandler(),
	// 	middleware.Auth,
	// 	middleware.NeedPerm(permission.Perm_SubmitFeedback),
	// 	feedbackcontrollers.CreateFeedbackPost)
	// //查询帖子这一块
	// ginEngine.GET("/api/feedback/:id", middleware.UnifiedErrorHandler(),
	// 	middleware.Auth,
	// 	middleware.NeedPerm(permission.Perm_QueryFeedbackLog),
	// 	feedbackcontrollers.GetAllFeedbackPosts) //由于给管理员配置函数麻烦,故先写死

	// //普通用户
	// // 获取用户信息
	// ginEngine.GET("/api/user/info/", middleware.UnifiedErrorHandler(),
	// 	middleware.Auth,
	// 	middleware.NeedPerm(permission.Perm_GetProfile),
	// 	accountControllers.GetAccountInfoUser)

	// // 修改用户信息
	// ginEngine.PUT("/api/user/info/", middleware.UnifiedErrorHandler(),
	// 	middleware.Auth,
	// 	middleware.NeedPerm(permission.Perm_UpdateProfile),
	// 	accountControllers.UpdateAccountInfoUser)

	// //管理员

	// //超级管理员
	// //获取用户信息
	// ginEngine.GET("/api/admin/users/", middleware.UnifiedErrorHandler(),
	// 	middleware.Auth,
	// 	middleware.NeedPerm(permission.Perm_GetAnyProfile),
	// 	accountControllers.GetAccountInfoAdmin)
	// // 新增用户
	// ginEngine.POST("/api/admin/users", middleware.UnifiedErrorHandler(),
	// 	middleware.Auth,
	// 	middleware.NeedPerm(permission.Perm_AddUser),
	// 	registerControllers.CreateUserAdmin)
	// //删除用户
	// ginEngine.DELETE("/api/admin/users/:id", middleware.UnifiedErrorHandler(),
	// 	middleware.Auth,
	// 	middleware.NeedPerm(permission.Perm_DeleteUser),
	// 	accountControllers.DeleteAccount)
	// //修改用户信息
	// ginEngine.PUT("/api/admin/users/:id", middleware.UnifiedErrorHandler(),
	// 	middleware.Auth,
	// 	middleware.NeedPerm(permission.Perm_EditAnyProfile),
	// 	accountControllers.UpdateAccountInfoAdmin)

	return ginEngine

}
