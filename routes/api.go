package routes

import (
	"github.com/gin-gonic/gin"
	"goex/app/http/controllers/api/v1/auth"
)

func RegisterAPIRoutes(r *gin.Engine) {

	v1 := r.Group("/v1")
	{
		authGroup := v1.Group("/auth")
		{
			suc := new(auth.SignupController)

			authGroup.POST("/signup/email/exist", suc.IsEmailExist)
			authGroup.POST("/signup/using-email", suc.SignupUsingEmail)

			vcc := new(auth.VerifyCodeController)

			authGroup.POST("/verify-codes/captcha", vcc.ShowCaptcha)
			authGroup.POST("/verify-codes/email", vcc.SendUsingEmail)

			lgc := new(auth.LoginController)

			authGroup.POST("/login/using-password", lgc.LoginByPassword)
		}
	}
}
