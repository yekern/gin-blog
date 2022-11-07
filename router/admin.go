package router

import (
	"esc.show/blog/app/api"
	"github.com/gin-gonic/gin"
)

func RegisterAdminRoutes(router *gin.Engine) {
	router.GET("/", api.UserApi{}.List)
	router.GET("/password", api.UserApi{}.Password)
}
