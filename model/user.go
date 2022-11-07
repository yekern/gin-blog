package model

import (
	"esc.show/blog/pkg/db"
	"esc.show/blog/pkg/db/datatype"
)

type User struct {
	db.Model
	Id        int64               `json:"id"`
	Nickname  string              `json:"nickname"`
	Username  string              `json:"username"`
	Password  string              `json:"password"`
	Status    int64               `json:"status"`
	Avatar    string              `json:"avatar"`
	LoginAt   *datatype.LocalTime `json:"login_at"`
	LoginIP   string              `json:"login_ip"`
	CreatedAt *datatype.LocalTime `json:"created_at"`
	UpdatedAt *datatype.LocalTime `json:"updated_at"`
}
