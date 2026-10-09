package world

import "github.com/google/wire"

// ProviderSet 提供世界的外部依赖
var ProviderSet = wire.NewSet(
	NewLane,
)
