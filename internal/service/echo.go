package service

import (
	"context"
	"strings"

	"example.ezboot/myna/api"
	"example.ezboot/myna/internal/world"
)

// EchoService 回声服务的具体实现
type EchoService struct {
	l *world.Lane

	api.UnimplementedEchoServiceServer
}

// Say 响应用户的请求, 并在其中加上一点点佐料
func (s *EchoService) Say(_ context.Context, req *api.SayRequest) (*api.SayResponse, error) {
	if strings.Contains(req.Text, "abc") {
		return nil, api.ErrEchoBadText
	}
	return &api.SayResponse{Text: s.l.Via(req.Text)}, nil
}

// NewEchoService 创建一个回声服务
func NewEchoService(l *world.Lane) (*EchoService, error) {
	return &EchoService{l: l}, nil
}
