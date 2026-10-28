//go:build wireinject

package main

import (
	"github.com/google/wire"

	"example.ezboot/myna/config"
	"example.ezboot/myna/internal/cli"
	"example.ezboot/myna/internal/plugin/example"
	"example.ezboot/myna/internal/server"
	"example.ezboot/myna/internal/service"
	"example.ezboot/myna/internal/world"

	"github.com/ez4bk/ezboot"
	"github.com/ez4bk/ezboot/infra"
)

// InitApp 初始化项目依赖并创建相应的应用
func InitApp() (*ezboot.Application, error) {
	panic(wire.Build(
		NewApp,
		cli.ProviderSet,
		ezboot.ProviderSet,
		example.ProviderSet,
		infra.NewMySQL,
		config.ProviderSet,
		server.ProviderSet,
		service.ProviderSet,
		world.ProviderSet,
	))
}
