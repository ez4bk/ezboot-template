package config

import (
	"github.com/google/wire"

	"github.com/ez4bk/ezboot"
)

// ProviderSet 提供配置相关的依赖
var ProviderSet = wire.NewSet(
	NewConfig,
	wire.FieldsOf(new(*Config), "World"),
)

// Config 当前应用的配置选项
type Config struct {
	Name  string `yaml:"name" json:"name"`
	World World  `yaml:"world" json:"world"`
}

// World 世界上总有一些奇怪的东西
type World struct {
	Mysterious string `yaml:"mysterious" json:"mysterious"`
}

// NewConfig 从配置文件加载服务配置
func NewConfig() (*Config, error) {
	cfg := new(Config)
	return cfg, ezboot.LoadUserConfig(cfg)
}
