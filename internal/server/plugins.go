package server

import (
	"github.com/ez4bk/ezboot"
	"github.com/ez4bk/ezboot/grpc"
	grpcclient "github.com/ez4bk/ezboot/grpc/client"
	"github.com/ez4bk/ezboot/http"
	"github.com/ez4bk/ezboot/infra"
	"github.com/ez4bk/ezboot/tracing"

	"example.ezboot/myna/internal/plugin/example"
)

// NewPlugins 按生命周期顺序收口插件。
//
// 顺序即 Close 的逆序依据：otel 放最前，最后关闭时先 flush span；
// 基础设施放在入口之前，保证服务器排空时依赖还活着。
func NewPlugins(
	otel *tracing.OTEL,
	db *infra.MySQL,
	gc *grpcclient.GRPCClient,
	ex *example.Plugin,
	gs *grpc.Server,
	hs *http.Server,
) []ezboot.Plugin {
	return []ezboot.Plugin{otel, db, gc, ex, gs, hs}
}
