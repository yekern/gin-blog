package api

import (
	"esc.show/blog/app/service"
	"esc.show/blog/model"
	"esc.show/blog/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/gookit/validate"
	"strconv"
)

type UserApi struct {
}

var userService service.UserService

func (a *UserApi) Index(ctx *gin.Context) {
	page := ctx.DefaultQuery("page", "1")
	pageSize := ctx.DefaultQuery("page_size", "10")

	pageInt, _ := strconv.Atoi(page)
	pageSizeInt, _ := strconv.Atoi(pageSize)

	list := userService.List(pageInt, pageSizeInt)

	utils.NewResponse(ctx).Data(list)
}

// Create 创建用户
func (a *UserApi) Create(ctx *gin.Context) {

	var user model.User
	err := ctx.ShouldBind(&user)
	if err != nil {
		utils.NewResponse(ctx).Error(err.Error())
		return
	}
	v := validate.Struct(&user)
	if !v.Validate() {
		utils.NewResponse(ctx).Error(v.Errors.One())
		return
	}
	//json["password"], _ = utils.NewPassword().EncodePassword(string(json["password"].([]byte)))
	err = userService.Create(&user)
	if err != nil {
		utils.NewResponse(ctx).Error(err.Error())
		return
	}
	utils.NewResponse(ctx).SuccessOk()
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
}

// Profile 获取个人资料
func (a *UserApi) Profile(ctx *gin.Context) {
	userId := ctx.GetInt64("userId")
	profile, err := userService.Profile(userId)
	if err != nil {
		utils.NewResponse(ctx).Error(err.Error())
		return
	}
	utils.NewResponse(ctx).Data(profile)
}
