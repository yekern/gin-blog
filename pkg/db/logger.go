package db

import (
	"go.uber.org/zap"
)

type GormLogger struct {
}

func (l GormLogger) Printf(format string, v ...interface{}) {
	zap.L().Debug("SQL ERROR",
		zap.Any("error", v[1]),
		zap.Any("sql_raw", v[4]),
		zap.Any("time", v[2]),
	)
}
