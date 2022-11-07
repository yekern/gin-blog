package kernel

import (
	"esc.show/blog/middleware"
	"esc.show/blog/router"
	"github.com/gin-gonic/gin"
)

// RegisterMiddleware 注册中间件
func RegisterMiddleware(engine *gin.Engine) {
	engine.Use(middleware.Demo())
}

// RegisterRouter 注册路由服务
func RegisterRouter(engine *gin.Engine) {
	router.RegisterAdminRoutes(engine)
}
