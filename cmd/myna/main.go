package main

import (
	"context"

	"example.ezboot/myna/config"
	"example.ezboot/myna/internal/cli"

	"github.com/ez4bk/ezboot"
	"github.com/ez4bk/ezboot/grpc"
	"github.com/ez4bk/ezboot/http"
	"github.com/ez4bk/ezboot/signal"
)

// NewApp 创建应用服务并初始化
func NewApp(cfg *config.Config, gs *grpc.Server, hs *http.Server, commands cli.Commands) (*ezboot.Application, error) {
	return ezboot.NewApp(
		ezboot.WithAppName(cfg.Name),
		ezboot.WithGrpcServer(gs),
		ezboot.WithHttpServer(hs),
		ezboot.WithCmdline(commands...),
	)
}

func main() {
	app, err := InitApp()
	if err != nil {
		panic(err)
	}

	ctx, cancel := signal.WithExit(context.Background())
	defer cancel()

	if err = app.Start(ctx); err != nil {
		panic(err)
	}
}
