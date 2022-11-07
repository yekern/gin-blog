package router

import (
	"esc.show/blog/app/api"
	"esc.show/blog/pkg/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterAdminRoutes(router *gin.Engine) {
	var userApi api.UserApi
	router.GET("/", userApi.Create)
	router.POST("/login", userApi.Login)
	router.Use(middleware.JWT())
	router.GET("/profile", userApi.Profile)
}
