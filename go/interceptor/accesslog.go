package interceptor

import (
	"context"
	"time"

	"github.com/rmanzoku/grpc-boilerplate/go/utility/log"
	"google.golang.org/grpc"
)

func UnaryServerAccessLogInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		start := time.Now()
		res, err := handler(ctx, req)
		nano := time.Since(start).Nanoseconds()
		log.AccessLog(ctx, req, res, err, info, nano)
		return res, err
	}
}
