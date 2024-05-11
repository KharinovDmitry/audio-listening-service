package middleware

import (
	"auth-service/internal/domain/service"
	"github.com/gin-gonic/gin"
)

func NewLogMiddleware(log service.LoggerService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		path := ctx.Request.URL.Path
		method := ctx.Request.Method
		log.Debug("Received %s request for path: %s", method, path)
		ctx.Next()
	}
}
