package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"goex/pkg/app"
	"goex/pkg/limiter"
	"goex/pkg/logger"
	"goex/pkg/response"
	"net/http"
)

// LimitIP global current limiting middleware, for IP current limiting
func LimitIP(limit string) gin.HandlerFunc {
	if app.IsTesting() {
		limit = "1000000-H"
	}

	return func(c *gin.Context) {
		key := limiter.GetKeyIP(c)
		if ok := limitHandler(c, key, limit); !ok {
			return
		}
		c.Next()
	}
}

func LimitPerRoute(limit string) gin.HandlerFunc {
	if app.IsTesting() {
		limit = "1000000-H"
	}
	return func(c *gin.Context) {

		c.Set("limiter-once", false)

		key := limiter.GetKeyRouteWithIP(c)
		if ok := limitHandler(c, key, limit); !ok {
			return
		}
		c.Next()
	}
}

func limitHandler(c *gin.Context, key string, limit string) bool {
	rate, err := limiter.CheckRate(c, key, limit)
	if err != nil {
		logger.LogIf(err)
		response.Abort500(c)
		return false
	}

	// ---- SET HEAD INFORMATION ----
	// X-RateLimit-Limit :10000 maximum number of visits
	// X-RateLimit-Remaining :9993 visits remaining
	// X-RateLimit-Reset :1513784506 at that point, the number of visits is reset to X-RateLimit-Limit
	c.Header("X-RateLimit-Limit", cast.ToString(rate.Limit))
	c.Header("X-RateLimit-Remaining", cast.ToString(rate.Remaining))
	c.Header("X-RateLimitReset-", cast.ToString(rate.Reset))

	if rate.Reached {
		c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
			"message": "Interface requests are too frequent",
		})
		return false
	}

	return true
}
