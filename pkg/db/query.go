package db

import "gorm.io/gorm"

// NewQuery 返回ORM查询
func NewQuery() *gorm.DB {
	return db
}
