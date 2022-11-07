package server

import (
	"context"
	"errors"
	"esc.show/blog/kernel"
	system_middleware "esc.show/blog/pkg/middleware"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// Server 服务结构体
type Server struct {
	Engine *gin.Engine
}

var server *Server

func init() {
	debug := viper.GetBool("server.debug")
	if debug == true {
		gin.SetMode("release")
	}
	server = &Server{Engine: gin.New()}
}

// registerStaticRoot 注册静态资源目录
func (s *Server) registerStaticRoot() {
	staticPrefix := fmt.Sprintf("/%s/", viper.GetString("server.static_prefix"))
	rootPath := viper.GetString("server.static_root")
	s.Engine.Static(staticPrefix, rootPath)
}

// registerRouter 注册路由
func (s *Server) registerRouter() {
	kernel.RegisterRouter(s.Engine)
}

// registerMiddleware 注册中间件
func (s *Server) registerMiddleware() {
	// 注册系统中间件
	s.Engine.Use(
		system_middleware.ApiLogger(),
		system_middleware.ApiRecover(viper.GetBool("log.stack")),
	)
	kernel.RegisterMiddleware(s.Engine)
}

// start 启动服务
func (s *Server) start() {

	serv := &http.Server{
		Addr:    viper.GetString("server.port"),
		Handler: s.Engine,
	}

	go func() {
		if err := serv.ListenAndServe(); err != nil && errors.Is(err, http.ErrServerClosed) {
			log.Printf("%s...\n", err.Error())
		}
	}()

	quitChan := make(chan os.Signal)
	signal.Notify(quitChan, syscall.SIGINT, syscall.SIGTERM)
	<-quitChan
	log.Println("正在关闭服务...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := serv.Shutdown(ctx); err != nil {
		log.Fatalln("服务器被迫关闭:", err)
	}
	log.Println("服务关闭成功...")
}

// NewServer 初始化Web服务
func NewServer() {
	server.registerStaticRoot()
	server.registerMiddleware()
	server.registerRouter()
	server.start()
}

// GetServer 获取sever实例
func GetServer() *Server {
	return server
}
