package ping_service

import (
	"context"

	"github.com/rmanzoku/grpc-boilerplate/go/feature/ping"
)

type PingServiceServer struct {
	ping.UnimplementedPingServiceServer
}

func (s *PingServiceServer) Echo(ctx context.Context, in *ping.MessageInput) (*ping.MessageOutput, error) {
	return &ping.MessageOutput{Msg: in.Msg}, nil
}
