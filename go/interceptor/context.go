package interceptor

import (
	"context"

	"google.golang.org/grpc"
)

func UnaryServerContextInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {

		ctx = ctxutil.Setup(ctx)
		reply, hErr := handler(ctx, req)

		return reply, hErr
	}
}
