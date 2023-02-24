package routes

import (
	"github.com/gin-gonic/gin"
	"goex/app/http/controllers/api/v1/auth"
	"goex/app/http/middlewares"
)

func RegisterAPIRoutes(r *gin.Engine) {

	v1 := r.Group("/v1")

	v1.Use(middlewares.LimitIP("200-H"))
	{
		authGroup := v1.Group("/auth")

		authGroup.Use(middlewares.LimitIP("1000-H"))
		{
			suc := new(auth.SignupController)
			authGroup.POST("/signup/email/exist", middlewares.GuestJWT(), middlewares.LimitPerRoute("60-H"), suc.IsEmailExist)
			authGroup.POST("/signup/using-email", middlewares.GuestJWT(), suc.SignupUsingEmail)

			vcc := new(auth.VerifyCodeController)
			authGroup.POST("/verify-codes/captcha", middlewares.LimitPerRoute("50-H"), vcc.ShowCaptcha)
			authGroup.POST("/verify-codes/email", middlewares.LimitPerRoute("20-H"), vcc.SendUsingEmail)

			lgc := new(auth.LoginController)
			authGroup.POST("/login/using-password", middlewares.GuestJWT(), lgc.LoginByPassword)
			authGroup.POST("/login/refresh-token", middlewares.AuthJWT(), lgc.RefreshToken)

			pwc := new(auth.PasswordController)
			authGroup.POST("/password-reset/using-email", middlewares.GuestJWT(), pwc.ResetByEmail)
		}
	}
}
