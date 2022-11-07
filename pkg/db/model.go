package db

import (
	"esc.show/blog/pkg/db/scopes"
	"gorm.io/gorm"
)

type Model struct {
}

// Query 查询实例
func (m *Model) Query() *gorm.DB {
	return db
}

// Paginate 分页器
func (m *Model) Paginate(page, pageSize int) *gorm.DB {
	return m.Query().Scopes(scopes.Paginate(page, pageSize))
}
