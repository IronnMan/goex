package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func main() {
	// New a gin engine instance
	r := gin.New()

	// Register a middleware
	r.Use(gin.Logger(), gin.Recovery())

	// Register a route
	r.GET("/", func(c *gin.Context) {

		// Response  in JSON format
		c.JSON(http.StatusOK, gin.H{
			"Hello": "World!",
		})
	})

	// Handling 404 requests
	r.NoRoute(func(c *gin.Context) {
		// Get the Accept information of the header information
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

	// Run service, ports is 8000
	r.Run(":8000")
}
