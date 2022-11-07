package db

import (
	"go.uber.org/zap"
)

type GormLogger struct {
}

func (l GormLogger) Printf(format string, v ...interface{}) {
	zap.L().Error("SQL", zap.Any("sql_raw", v))
}
