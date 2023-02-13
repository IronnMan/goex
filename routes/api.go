package routes

import (
	"github.com/gin-gonic/gin"
	"goex/app/http/controllers/api/auth"
)

func RegisterAPIRoutes(r *gin.Engine) {

	v1 := r.Group("/v1")
	{
		authGroup := v1.Group("/auth")
		{
			suc := new(auth.SignupController)
			
			authGroup.POST("/signup/email/exist", suc.IsEmailExist)
		}
	}
}
