package middlewares

import (
	"github.com/gin-gonic/gin"
	"goex/pkg/jwt"
	"goex/pkg/response"
)

func GuestJWT() gin.HandlerFunc {
	return func(c *gin.Context) {

		if len(c.GetHeader("Authorization")) > 0 {

			// parse the token success
			_, err := jwt.NewJWT().ParserToken(c)
			if err == nil {
				response.Unauthorized(c, "Please access as a guest")
				c.Abort()
				return
			}
		}

		c.Next()
	}
}
