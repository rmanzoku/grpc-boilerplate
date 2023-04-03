package ping_service

import (
	"context"

	"github.com/rmanzoku/grpc-boilerplate/go/feature/ping"
	"github.com/rmanzoku/grpc-boilerplate/go/utility/ctxutil"
)

type PingServiceServer struct {
	ping.UnimplementedPingServiceServer
}

func (s *PingServiceServer) Echo(ctx context.Context, in *ping.MessageInput) (*ping.MessageOutput, error) {
	return &ping.MessageOutput{Msg: in.Msg}, nil
}

func (s *PingServiceServer) Now(ctx context.Context, in *ping.Empty) (*ping.Time, error) {
	now := ctxutil.ExtractTime(ctx)
	return &ping.Time{T: uint64(now)}, nil
}
