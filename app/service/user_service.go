package service

import (
	"errors"
	"esc.show/blog/app/resource"
	"esc.show/blog/model"
	"esc.show/blog/pkg/db"
	"esc.show/blog/pkg/utils"
)

type UserService struct {
	db *db.Model
}

// UsersList 分页结构体
type UsersList struct {
	Items []model.User `json:"items"`
	Total int64        `json:"total"`
}

//UserFrom  新增用户表单
type UserFrom struct {
	Nickname string `json:"nickname" validate:"-"`
	Username string `json:"username" validate:"required" label:"用户名"`
	Password string `json:"password" validate:"required" label:"密码"`
	Status   int64  `json:"status" validate:"-"`
}

// Create 创建用户
func (s *UserService) Create(userForm *UserFrom) error {
	user := &model.User{
		Nickname: userForm.Nickname,
		Username: userForm.Username,
		Password: userForm.Password,
		Status:   userForm.Status,
	}
	result := s.db.Query().Create(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// Login 登录
func (s *UserService) Login(username, password string) (*model.User, error) {
	var user model.User
	result := s.db.Query().Where("username = ? ", username).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	if utils.NewPassword().DeCodePassword(password, user.Password) {
		if user.Status == 0 {
			return nil, errors.New("账户已停用")
		}
		return &user, nil
	} else {
		return nil, errors.New("用户名密码错误")
	}
}

func (s *UserService) Profile(userId int64) (*resource.Profile, error) {
	var user model.User
	result := s.db.Query().Where("id =?", userId).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &resource.Profile{
		Username: user.Username,
		Nickname: user.Nickname,
		Avatar:   user.Avatar,
	}, nil
}

func (s UserService) List(page, pageSize int) *UsersList {

	var users []model.User
	var total int64
	s.db.Query().Model(&model.User{}).Count(&total)
	s.db.Paginate(page, pageSize).Find(&users)
	return &UsersList{
		Items: users,
		Total: total,
	}
}
