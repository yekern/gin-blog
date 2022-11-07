package service

import (
	"esc.show/blog/model"
	"esc.show/blog/pkg/db"
)

type UserService struct {
	db *db.Model
}

type UsersList struct {
	Items []model.User `json:"items"`
	Total int64        `json:"total"`
}

func (s UserService) List() *UsersList {

	var users []model.User
	var total int64
	s.db.Query().Count(&total)
	s.db.Paginate(1, 10).Find(&users)
	return &UsersList{
		Items: users,
		Total: total,
	}
}
