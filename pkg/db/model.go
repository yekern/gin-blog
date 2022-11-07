package db

import "gorm.io/gorm"

type Model struct {
}

func (m *Model) Query() *gorm.DB {
	return db
}
