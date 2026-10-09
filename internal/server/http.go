package server

import (
	"errors"

	"github.com/ez4bk/ezboot"
	"github.com/ez4bk/ezboot/grpc"
	"github.com/ez4bk/ezboot/http"
	"github.com/gin-gonic/gin"
)

// NewHttpServer 创建HTTP服务
func NewHttpServer(cfg *ezboot.Config, gs *grpc.Server) (*http.Server, error) {
	return http.New(cfg,
		http.WithHealthy(),
		http.WithAdditionHandler(http.MethodAny, "/api/*any", gs),
		http.WithAdditionGinHandler(http.MethodAny, "/test/*any", func(c *gin.Context) {
			panic(errors.New(c.Request.URL.String()))
		}),
	)
}
