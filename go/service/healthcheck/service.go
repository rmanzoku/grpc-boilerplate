package healthcheck_service

import (
	"context"

	"github.com/rmanzoku/grpc-boilerplate/go/feature/healthcheck"
)

type HealthcheckServiceServer struct {
	healthcheck.UnimplementedHealthcheckServiceServer
}

func (s *HealthcheckServiceServer) Root(ctx context.Context, in *healthcheck.Empty) (*healthcheck.MessageOutput, error) {
	return &healthcheck.MessageOutput{Msg: "ok"}, nil
}

func (s *HealthcheckServiceServer) Healthcheck(ctx context.Context, in *healthcheck.Empty) (*healthcheck.MessageOutput, error) {
	return &healthcheck.MessageOutput{Msg: "ok"}, nil
}
