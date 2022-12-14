package middleware

import (
	"esc.show/blog/pkg/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func JWT() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		header := ctx.GetHeader("Authorization")
		if header == "" {
			utils.NewResponse(ctx).SetCode(http.StatusUnauthorized).Message("Unauthorized Miss Token", nil)
			ctx.Abort()
		}
		parts := strings.Split(header, " ")
		if len(parts) != 2 {
			utils.NewResponse(ctx).SetCode(http.StatusUnauthorized).Message("Unauthorized", nil)
			ctx.Abort()
		}
		if parts[0] != "Bearer" {
			utils.NewResponse(ctx).SetCode(http.StatusUnauthorized).Message("Unauthorized", nil)
			ctx.Abort()
		}
		if user, err := utils.NewJWT().Decode(parts[1]); err != nil {
			utils.NewResponse(ctx).SetCode(http.StatusUnauthorized).Message("Unauthorized", nil)
			ctx.Abort()
		} else {
			ctx.Set("userId", user.UserId)
		}
		ctx.Next()
	}
}
