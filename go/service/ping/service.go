package ping_service

import (
	"context"
	"time"

	"github.com/rmanzoku/grpc-boilerplate/go/feature/ping"
)

type PingServiceServer struct {
	ping.UnimplementedPingServiceServer
}

func (s *PingServiceServer) Echo(ctx context.Context, in *ping.MessageInput) (*ping.MessageOutput, error) {
	return &ping.MessageOutput{Msg: in.Msg}, nil
}

func (s *PingServiceServer) Now(ctx context.Context, in *ping.Empty) (*ping.Time, error) {
	return &ping.Time{T: uint64(time.Now().Unix())}, nil
}
