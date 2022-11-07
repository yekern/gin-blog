package middleware

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Demo() gin.HandlerFunc {
	return func(context *gin.Context) {
		zap.L().Debug("this is demo middleware")
		context.Next()
	}
}
