package model

import (
	"esc.show/blog/pkg/db"
	"esc.show/blog/pkg/db/datatype"
	"esc.show/blog/pkg/utils"
	"gorm.io/gorm"
)

type User struct {
	db.Model
	Id        int64               `json:"id"`
	Nickname  string              `json:"nickname"`
	Username  string              `json:"username" validate:"required" label:"用户名"`
	Password  string              `json:"password" validate:"required" label:"密码"`
	Status    int64               `json:"status" validate:"-" label:"状态"`
	Avatar    string              `json:"avatar"`
	LoginAt   *datatype.LocalTime `json:"login_at"`
	LoginIP   string              `json:"login_ip"`
	CreatedAt *datatype.LocalTime `json:"-"`
	UpdatedAt *datatype.LocalTime `json:"-"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.Password, err = utils.NewPassword().EncodePassword(u.Password)
	return
}
