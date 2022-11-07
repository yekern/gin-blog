package api

import (
	"esc.show/blog/app/service"
	"github.com/gin-gonic/gin"
)

type UserApi struct {
}

func (a UserApi) List(ctx *gin.Context) {
	var userService service.UserService
	ctx.JSON(200, gin.H{
		"code":    200,
		"message": "ok",
		"data":    userService.List(),
	})
}
