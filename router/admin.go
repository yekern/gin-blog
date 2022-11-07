package router

import (
	"esc.show/blog/app/api"
	"esc.show/blog/pkg/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterAdminRoutes(router *gin.Engine) {
	admin := router.Group("/admin")
	{
		var userApi api.UserApi
		admin.POST("/auth/login", userApi.Login)
		admin.Use(middleware.JWT())
		admin.GET("/profile", userApi.Profile)
		admin.GET("/users", userApi.Index)
		admin.PUT("/users", userApi.Create)
	}

}
