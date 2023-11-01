package middleware

import (
	"backstage/common/http/cors"
	"github.com/gin-gonic/gin"
)

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set(cors.AccessControlAllowOrigin, "*")
		c.Writer.Header().Set(cors.AccessControlAllowMethods, "POST, GET, OPTIONS, PUT, DELETE")
		c.Writer.Header().Set(cors.AccessControlAllowHeaders, "Accept, Content-Type, Content-Length, Accept-Encoding,Authorization")

		if c.Request.Method == "OPTIONS" {
			c.Writer.WriteHeader(200)
			return
		}
		c.Next()
	}
}
