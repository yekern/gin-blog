package api

import (
	"esc.show/blog/pkg/utils"
	"fmt"
	"github.com/gin-gonic/gin"
)

type UserApi struct {
}

func (a UserApi) List(ctx *gin.Context) {
	utils.NewResponse(ctx).SuccessOk()
}

func (a UserApi) Password(ctx *gin.Context) {
	password, err := utils.NewPassword().EncodePassword("123456")
	if err != nil {
		fmt.Println("密码加密失败")
	}
	ctx.JSON(200, gin.H{
		"code":    200,
		"message": "ok",
		"data":    password,
	})
}
