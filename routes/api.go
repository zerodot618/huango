// Package routes 注册路由
package routes

import (
	"net/http"
	controllers "zerodot618/huango/app/http/controllers/api/v1"
	"zerodot618/huango/app/http/controllers/api/v1/auth"
	"zerodot618/huango/app/http/middlewares"
	"zerodot618/huango/pkg/config"

	"github.com/gin-gonic/gin"
)

// RegisterAPIRoutes 注册网页相关路由
func RegisterAPIRoutes(r *gin.Engine) {

	// v1 版本路由组，我们所有的 v1 版本的路由都将存放到这里
	var v1 *gin.RouterGroup
	if len(config.Get("app.api_domain")) == 0 {
		v1 = r.Group("/api/v1")
	} else {
		v1 = r.Group("/v1")
	}

	// 全局限流中间件：每小时限流。这里是所有 API （根据 IP）请求加起来。
	// 作为参考 Github API 每小时最多 60 个请求（根据 IP）。
	// 测试时，可以调高一点。
	v1.Use(middlewares.LimitIP("200-H"))

	{
		// 注册一个路由
		v1.GET("/", func(c *gin.Context) {
			// 以 JSON 格式响应
			c.JSON(http.StatusOK, gin.H{
				"Hello": "World!",
			})
		})

		authGroup := v1.Group("/auth")
		// 限流中间件：每小时限流，作为参考 Github API 每小时最多 60 个请求（根据 IP）
		// 测试时，可以调高一点
		authGroup.Use(middlewares.LimitIP("1000-H"))
		{
			// 登录
			lgc := new(auth.LoginController)
			// 使用手机号，短信验证码进行登录
			authGroup.POST("/login/using-phone", middlewares.GuestJWT(), lgc.LoginByPhone)
			// 支持手机号，Email 和 用户名
			authGroup.POST("/login/using-password", middlewares.GuestJWT(), lgc.LoginByPassword)
			// 刷新 Access Token
			authGroup.POST("/login/refresh-token", middlewares.AuthJWT(), lgc.RefreshToken)

			// 重置密码
			pwc := new(auth.PasswordController)
			// 通过手机找回密码
			authGroup.POST("/password-reset/using-phone", middlewares.GuestJWT(), pwc.ResetByPhone)
			// 通过 Email 找回密码
			authGroup.POST("/password-reset/using-email", middlewares.GuestJWT(), pwc.ResetByEmail)

			// 注册
			suc := new(auth.SignupController)
			// 手机号注册
			authGroup.POST("/signup/using-phone", middlewares.GuestJWT(), suc.SignupUsingPhone)
			// 邮箱注册
			authGroup.POST("/signup/using-email", middlewares.GuestJWT(), suc.SignupUsingEmail)
			// 判断手机是否已注册
			authGroup.POST("/signup/phone/exist", middlewares.GuestJWT(), middlewares.LimitPerRoute("60-H"), suc.IsPhoneExist)
			// 判断 Email 是否已注册
			authGroup.POST("/signup/email/exist", middlewares.GuestJWT(), middlewares.LimitPerRoute("60-H"), suc.IsEmailExist)

			// 发送验证码
			vcc := new(auth.VerifyCodeController)
			// 图片验证码，需要加限流
			authGroup.POST("/verify-codes/captcha", middlewares.LimitPerRoute("50-H"), vcc.ShowCaptcha)
			authGroup.POST("/verify-codes/phone", middlewares.LimitPerRoute("20-H"), vcc.SendUsingPhone)
			authGroup.POST("/verify-codes/email", middlewares.LimitPerRoute("20-H"), vcc.SendUsingEmail)

		}

		// 用户
		uc := new(controllers.UsersController)
		// 获取当前用户
		v1.GET("/user", middlewares.AuthJWT(), uc.CurrentUser)
		usersGroup := v1.Group("/users")
		{
			usersGroup.GET("/", uc.Index)
			usersGroup.PUT("/", middlewares.AuthJWT(), uc.UpdateProfile)
			usersGroup.PUT("/email", middlewares.AuthJWT(), uc.UpdateEmail)
			usersGroup.PUT("/phone", middlewares.AuthJWT(), uc.UpdatePhone)
			usersGroup.PUT("/password", middlewares.AuthJWT(), uc.UpdatePassword)
			usersGroup.PUT("/avatar", middlewares.AuthJWT(), uc.UpdateAvatar)
		}

		// 分类
		cgc := new(controllers.CategoriesController)
		cgcGroup := v1.Group("/categories")
		{
			cgcGroup.GET("/", cgc.Index)
			cgcGroup.POST("/", middlewares.AuthJWT(), cgc.Store)
			cgcGroup.PUT("/:id", middlewares.AuthJWT(), cgc.Update)
			cgcGroup.DELETE("/:id", middlewares.AuthJWT(), cgc.Delete)
		}

		// 话题
		tpc := new(controllers.TopicsController)
		tpcGroup := v1.Group("/topics")
		{
			tpcGroup.GET("/", tpc.Index)
			tpcGroup.POST("/", middlewares.AuthJWT(), tpc.Store)
			tpcGroup.PUT("/:id", middlewares.AuthJWT(), tpc.Update)
			tpcGroup.DELETE("/:id", middlewares.AuthJWT(), tpc.Delete)
			tpcGroup.GET("/:id", tpc.Show)
		}

		// 友情链接
		lsc := new(controllers.LinksController)
		linksGroup := v1.Group("/links")
		{
			linksGroup.GET("", lsc.Index)
		}

		// 短地址
		stc := new(controllers.ShortenerController)
		shortenerGroup := v1.Group("/shortener")
		{
			shortenerGroup.POST("/encode", middlewares.AuthJWT(), stc.CreateShortURL)
			shortenerGroup.GET("/decode/:short-url", stc.ReturnLongURL)
		}
	}
}
