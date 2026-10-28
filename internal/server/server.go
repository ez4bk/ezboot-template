package server

import "github.com/google/wire"

// ProviderSet 提供服务相关的依赖
var ProviderSet = wire.NewSet(
	NewGRPCServer,
	NewHttpServer,
	NewPlugins,
)
