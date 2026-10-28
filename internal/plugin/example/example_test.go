package example

import (
	"context"
	"testing"
	"time"

	"github.com/ez4bk/ezboot"
)

// TestPluginLifecycle 演示如何在不启动整个应用的情况下测试一个插件：
// 插件只依赖 ctx、配置和 logger，所以可以直接驱动它的四个阶段。
func TestPluginLifecycle(t *testing.T) {
	p := New(&Config{Interval: 5 * time.Millisecond}, ezboot.DiscardLogger)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	done := make(chan error, 1)
	go func() { done <- p.Start(ctx) }()

	deadline := time.Now().Add(time.Second)
	for {
		if err := p.Ready(context.Background()); err == nil {
			break
		}
		if time.Now().After(deadline) {
			t.Fatal("plugin never became ready")
		}
		time.Sleep(time.Millisecond)
	}

	drainCtx, cancelDrain := context.WithTimeout(context.Background(), time.Second)
	defer cancelDrain()

	if err := p.Drain(drainCtx); err != nil {
		t.Fatalf("Drain: %v", err)
	}

	cancel()

	select {
	case err := <-done:
		if err != nil {
			t.Fatalf("Start: %v", err)
		}
	case <-time.After(time.Second):
		t.Fatal("Start did not return after the context was cancelled")
	}
}
