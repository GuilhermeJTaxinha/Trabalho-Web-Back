package middlewares

import "github.com/gin-gonic/gin"

func MiddlewaresGet() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		origin := ctx.Request.Header.Get("Origin")
		ctx.Writer.Header().Set("Access-Control-Allow-Origin", origin)

		if ctx.Request.Method == "OPTIONS" {
			ctx.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
			ctx.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
			ctx.AbortWithStatus(204)
			return
		}
		ctx.Next()
	}
}
