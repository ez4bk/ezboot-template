// Package example 是一个自定义插件样例，演示 ezboot.Plugin 的各个阶段，
// 以及插件配置从 <env>_custom.yaml 读取的方式：
//
//	example:
//	  interval: 1s
package example

import (
	"context"
	"errors"
	"sync"
	"sync/atomic"
	"time"

	"github.com/google/wire"

	"github.com/ez4bk/ezboot"
)

// Config 是插件的配置，只存在于 <env>_custom.yaml 中。
type Config struct {
	Interval time.Duration `yaml:"interval"`
}

// ProviderSet 提供示例插件。
var ProviderSet = wire.NewSet(NewConfig, New)

// NewConfig 从 custom 配置文件读取配置，缺省 1s。
func NewConfig() (*Config, error) {
	var root struct {
		Example Config `yaml:"example"`
	}

	if err := ezboot.LoadUserConfig(&root); err != nil {
		return nil, err
	}

	cfg := root.Example
	if cfg.Interval <= 0 {
		cfg.Interval = time.Second
	}

	return &cfg, nil
}

// Plugin 每隔 interval 输出一条调试日志，用来演示自定义插件的生命周期。
//
// Close 没有覆盖：它由 ezboot.PluginBase 提供空实现。
type Plugin struct {
	ezboot.PluginBase

	logger   *ezboot.Logger
	interval time.Duration

	ticks atomic.Int64

	stopOnce sync.Once
	stop     chan struct{}
	done     chan struct{}
}

// New 创建插件。
func New(cfg *Config, logger *ezboot.Logger) *Plugin {
	interval := cfg.Interval
	if interval <= 0 {
		interval = time.Second
	}

	return &Plugin{
		logger:   logger,
		interval: interval,
		stop:     make(chan struct{}),
		done:     make(chan struct{}),
	}
}

// Name 返回插件名，只用于日志与错误信息。
func (p *Plugin) Name() string { return "example" }

// Start 启动心跳循环，并阻塞到 ctx 结束。
//
// 阻塞是契约的一部分：Start 返回就代表插件已经停止。
func (p *Plugin) Start(ctx context.Context) error {
	ticker := time.NewTicker(p.interval)
	defer ticker.Stop()

	go func() {
		defer close(p.done)

		for {
			select {
			case <-p.stop:
				return
			case <-ctx.Done():
				return
			case <-ticker.C:
				n := p.ticks.Add(1)
				p.logger.Debugw("example", "ticks", n)
			}
		}
	}()

	<-ctx.Done()
	return nil
}

// Ready 在第一次心跳之前返回错误，lifecycle 会轮询到成功为止。
//
// 它只在启动期被调用；不要在这里做重试或长时间阻塞。
func (p *Plugin) Ready(context.Context) error {
	if p.ticks.Load() == 0 {
		return errors.New("example: waiting for the first tick")
	}
	return nil
}

// Drain 停止心跳并等待循环退出。ctx 是 lifecycle 新建的、带超时的。
func (p *Plugin) Drain(ctx context.Context) error {
	p.stopOnce.Do(func() { close(p.stop) })

	select {
	case <-p.done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}
