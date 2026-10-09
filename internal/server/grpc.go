package server

import (
	"github.com/ez4bk/ezboot"
	"github.com/ez4bk/ezboot/bizerr"
	"github.com/ez4bk/ezboot/grpc"
	"github.com/ez4bk/ezboot/grpc/gateway"
	_grpc "google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"example.ezboot/myna/api"
	"example.ezboot/myna/internal/service"
)

// NewGRPCServer 创建GRPC服务
func NewGRPCServer(cfg *ezboot.Config, echo *service.EchoService) (*grpc.Server, error) {
	gh, err := gateway.New(
		gateway.WithErrorHandler(bizerr.GatewayErrorHandler()),
		gateway.WithHandlerOutgoingHeader("set-cookie"),
		gateway.WithHandlerBaseEndpoint(cfg.App.Middleware.GRPC.Server.Default.Addr),
		gateway.WithHandlerEndpointRegister(api.RegisterEchoServiceHandlerFromEndpoint),
		gateway.WithHandlerDialOptions(_grpc.WithTransportCredentials(insecure.NewCredentials())),
	)
	if err != nil {
		return nil, err
	}

	return grpc.New(cfg,
		grpc.WithGatewayHandler(gh),
		grpc.WithGRPCServerRegister(api.RegisterEchoServiceServer, echo),
	)
}
