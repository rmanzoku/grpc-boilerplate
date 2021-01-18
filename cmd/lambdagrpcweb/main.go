package main

import (
	"context"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/httpadapter"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_logger "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"github.com/rmanzoku/grpc-boilerplate/feature/ping"
	ping_service "github.com/rmanzoku/grpc-boilerplate/service/ping"
	"go.uber.org/zap"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/grpc"
)

var adapter *httpadapter.HandlerAdapter

func registerServices(s *grpc.Server) {
	ping.RegisterPingServiceServer(s, &ping_service.PingServiceServer{})
}

func init() {
	// Logging
	logger := zap.NewExample()
	grpc_logger.ReplaceGrpcLoggerV2(logger)

	// Base grpc server
	grpcServer := grpc.NewServer(grpc_middleware.WithUnaryServerChain(
		grpc_ctxtags.UnaryServerInterceptor(),
		grpc_logger.UnaryServerInterceptor(logger),
		grpc_recovery.UnaryServerInterceptor(),
	))
	registerServices(grpcServer)

	// Wrapping grpc-web server
	grpcWebServer := grpcweb.WrapServer(grpcServer,
		grpcweb.WithOriginFunc(func(origin string) bool { return true }),
	)

	// Wrapping http2 for standalone
	httpServer := &http.Server{
		Handler: h2c.NewHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			grpcWebServer.ServeHTTP(w, r)

			// for cors preflight
			// grpcweb package use https://github.com/rs/cors/blob/master/cors.go#L202
			// cors package doesn't exec WriteHeader, it is needed lambda proxy
			if grpcWebServer.IsAcceptableGrpcCorsRequest(r) {
				w.WriteHeader(http.StatusNoContent)
			}
		}), &http2.Server{}),
	}

	// Lambda adapter
	adapter = httpadapter.New(httpServer.Handler)
}

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return adapter.ProxyWithContext(ctx, req)
}

func main() {
	lambda.Start(Handler)
}
