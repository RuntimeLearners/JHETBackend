package router

import (
	"JHETBackend/common/permission"
	middleware "JHETBackend/middlewares"
	"fmt"

	"JHETBackend/controllers/accountControllers"
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

	ginEngine.POST("/api/auth/register", middleware.UnifiedErrorHandler(), registerControllers.CreateStudentUser)

	//上传图片这一块
	//ginEngine.POST("/api/upload/image", middleware.UnifiedErrorHandler(), middleware.Auth, fileControllers.UploadImage)

	//用户信息这一块
	// 无需权限 测试用
	// ginEngine.GET("/api/user/info/", middleware.UnifiedErrorHandler(),
	// 	middleware.Auth,
	// 	accountControllers.GetAccountInfoUser)
	// ginEngine.GET("/api/admin/users/", middleware.UnifiedErrorHandler(),
	// 	middleware.Auth,
	// 	accountControllers.GetAccountInfoAdmin)

	//普通用户获取用户信息
	ginEngine.GET("/api/user/info/", middleware.UnifiedErrorHandler(),
		middleware.Auth,
		middleware.NeedPerm(permission.Perm_GetProfile),
		accountControllers.GetAccountInfoUser)

	//管理员获取用户信息
	ginEngine.GET("/api/admin/users/", middleware.UnifiedErrorHandler(),
		middleware.Auth,
		middleware.NeedPerm(permission.Perm_GetAnyProfile),
		accountControllers.GetAccountInfoAdmin)

	return ginEngine

}
