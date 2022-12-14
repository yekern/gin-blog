package app

import (
	"esc.show/blog/kernel"
	"esc.show/blog/pkg/config"
	"esc.show/blog/pkg/db"
	"esc.show/blog/pkg/logger"
	"esc.show/blog/pkg/server"
)

func NewApp() {
	// 初始化配置
	config.InitConfig()
	// 初始化日志
	logger.InitLog()
	// 初始化数据库连接
	db.InitDatabase()
	// 注册其他扩展信息
	kernel.ExtensionExec()
	// 创建Web服务
	server.NewServer()
}
