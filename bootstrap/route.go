package bootstrap

import (
	"github.com/gin-gonic/gin"
	"goex/app/http/middlewares"
	"goex/routes"
	"net/http"
	"strings"
)

// SetupRoute route initialization
func SetupRoute(r *gin.Engine) {

	registerGlobalMiddleWare(r)

	routes.RegisterAPIRoutes(r)

	setup404Handler(r)
}

func registerGlobalMiddleWare(r *gin.Engine) {
	r.Use(
		middlewares.Logger(),
		middlewares.Recovery(),
	)
}

func setup404Handler(r *gin.Engine) {

	r.NoRoute(func(c *gin.Context) {
		acceptString := c.Request.Header.Get("Accept")
		if strings.Contains(acceptString, "text/html") {
			// HTML
			c.String(http.StatusNotFound, "The page returns 404")
		} else {
			// JSON
			c.JSON(http.StatusNotFound, gin.H{
				"error_code":    404,
				"error_message": "The route is not defined, please confirm whether the url and request method are correct.",
			})
		}
	})
}
