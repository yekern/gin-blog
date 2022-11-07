package model

import "esc.show/blog/pkg/db/datatype"

type Role struct {
	ID        int64               `json:"id"`
	Name      string              `json:"name"`
	Status    int64               `json:"status"`
	CreatedAt *datatype.LocalTime `json:"created_at"`
	UpdatedAt *datatype.LocalTime `json:"updated_at"`
}
