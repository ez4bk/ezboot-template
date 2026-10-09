package service

import "github.com/google/wire"

// ProviderSet 提供服务实现相关的依赖
var ProviderSet = wire.NewSet(
	NewEchoService,
)
