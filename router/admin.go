package router

import (
	"esc.show/blog/app/api"
	"esc.show/blog/pkg/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterAdminRoutes(router *gin.Engine) {
	router.GET("/", api.UserApi{}.Create)
	router.POST("/login", api.UserApi{}.Login)
	router.Use(middleware.JWT())
	router.GET("/profile", api.UserApi{}.Profile)
}
