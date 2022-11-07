package model

import (
	"esc.show/blog/pkg/db"
	"esc.show/blog/pkg/db/datatype"
	"fmt"
)

type User struct {
	db.Model
	Id        int64               `json:"id"`
	Nickname  string              `json:"nickname"`
	Username  string              `json:"username"`
	Password  string              `json:"password"`
	Status    int64               `json:"status"`
	CreatedAt *datatype.LocalTime `json:"created_at"`
	UpdatedAt *datatype.LocalTime `json:"updated_at"`
}

func (u *User) List() {
	var user User
	u.Query().First(&user)
	fmt.Println(user)
}

func (u *User) Create() {
	user := &User{
		Nickname:  "demo",
		Username:  "demo",
		Password:  "123456",
		Status:    0,
		CreatedAt: nil,
		UpdatedAt: nil,
	}
	result := u.Query().Create(user)
	fmt.Println(result.Error)
}

//
//func (u *User) Update() {
//	DB.Save(u)
//}
