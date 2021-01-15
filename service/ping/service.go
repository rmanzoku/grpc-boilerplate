package ping

import (
	"context"

	"github.com/rmanzoku/grpc-boilerplate/feature/ping"
)

type PingService struct {
	ping.UnimplementedPingServiceServer
}

func (s *PingService) Echo(ctx context.Context, in *ping.MessageInput) (*ping.MessageOutput, error) {
	return &ping.MessageOutput{Msg: in.Msg}, nil
}
