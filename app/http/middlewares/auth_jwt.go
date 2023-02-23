package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goex/app/models/user"
	"goex/pkg/config"
	"goex/pkg/jwt"
	"goex/pkg/response"
)

func AuthJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, err := jwt.NewJWT().ParserToken(c)

		if err != nil {
			response.Unauthorized(c, fmt.Sprintf("Please see %v related interface certification documentation", config.GetString("app.name")))
			return
		}

		userModel := user.Get(claims.UserID)
		if userModel.ID == 0 {
			response.Unauthorized(c, "Can't find corresponding user")
			return
		}

		c.Set("current_user_id", userModel.GetStringID())
		c.Set("current_user_name", userModel.Name)
		c.Set("current_user", userModel)

		c.Next()
	}
}
