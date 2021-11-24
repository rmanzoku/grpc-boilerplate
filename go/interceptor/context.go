package interceptor

import (
	"context"

	"github.com/rmanzoku/grpc-boilerplate/go/utility/ctxutil"
	"google.golang.org/grpc"
)

func UnaryServerContextInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {

		ctx = ctxutil.Setup(ctx)
		reply, hErr := handler(ctx, req)
		if hErr != nil {
			err := ctxutil.RollbackWithErr(ctx, hErr)
			return reply, err
		}

		err := ctxutil.Commit(ctx)
		if err != nil {
			return reply, err
		}

		return reply, hErr
	}
}
