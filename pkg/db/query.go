package db

import "gorm.io/gorm"

// Query 查询实例
func (m *Model) Query() *gorm.DB {
	return db
}
