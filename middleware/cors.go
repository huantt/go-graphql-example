package middleware

import (
	"github.com/gin-gonic/gin"
)

var allowedOriginsMap map[string]bool

func CORS(allowedOrigins []string) func(c *gin.Context) {
	allowedOriginsMap = make(map[string]bool)
	for _, origin := range allowedOrigins {
		allowedOriginsMap[origin] = true
	}
	return func(c *gin.Context) {
		if len(allowedOrigins) == 0 {
			c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		} else if allowedOriginsMap[c.Request.Host] {
			c.Writer.Header().Set("Access-Control-Allow-Origin", c.Request.Host)
		} else {
			c.Writer.Header().Set("Access-Control-Allow-Origin", allowedOrigins[0])
		}
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With, login-type")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")
		if c.Request.Method == "OPTIONS" {
			c.Writer.Header().Set("Cache-Control", "public, immutable, no-transform, s-maxage=2592000, max-age=2592000")
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	}
}
