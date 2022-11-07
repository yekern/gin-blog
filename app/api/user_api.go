package api

import (
	"esc.show/blog/app/service"
	"esc.show/blog/model"
	"esc.show/blog/pkg/utils"
	"github.com/gin-gonic/gin"
)

type UserApi struct {
}

var userService service.UserService

// Create 创建用户
func (a *UserApi) Create(ctx *gin.Context) {
	password, err := utils.NewPassword().EncodePassword("123456")
	if err != nil {
		utils.NewResponse(ctx).Error("密码不符合规范:" + err.Error())
		return
	}
	user := &model.User{
		Nickname: "Admin",
		Username: "admin",
		Password: password,
		Status:   1,
	}
	user, err = userService.Create(user)
	if err != nil {
		utils.NewResponse(ctx).Error("用户创建失败:" + err.Error())
		return
	}
	utils.NewResponse(ctx).Data(user)
	return
}

// LoginForm 登录表单
type LoginForm struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Login 登录
func (a *UserApi) Login(ctx *gin.Context) {

	var form LoginForm

	if ctx.ShouldBindJSON(&form) != nil {
		utils.NewResponse(ctx).Error("无效的用户名或密码")
		return
	}
	user, err := userService.Login(form.Username, form.Password)
	if err != nil {
		utils.NewResponse(ctx).SetCode(200).Error(err.Error())
		return
	}
	utils.NewResponse(ctx).Data(utils.NewJWT().CreateToken(user.Id))
	return
}

// Profile 获取个人资料
func (u *UserApi) Profile(ctx *gin.Context) {
	userId := ctx.GetInt64("userId")
	profile, err := userService.Profile(userId)
	if err != nil {
		utils.NewResponse(ctx).Error(err.Error())
		return
	}
	utils.NewResponse(ctx).Data(profile)
	return
}
