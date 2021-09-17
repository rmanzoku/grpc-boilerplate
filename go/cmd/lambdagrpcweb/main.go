package main

import (
	"context"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/handlerfunc"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_logger "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"github.com/rmanzoku/grpc-boilerplate/feature/ping"
	ping_service "github.com/rmanzoku/grpc-boilerplate/service/ping"
	"github.com/rs/cors"
	"go.uber.org/zap"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/grpc"
)

var adapter *handlerfunc.HandlerFuncAdapterV2

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
			if grpcWebServer.IsGrpcWebRequest(r) {
				grpcWebServer.ServeHTTP(w, r)
			}

			// CORSを通すおまじない
			if grpcWebServer.IsAcceptableGrpcCorsRequest(r) {
				cors.New(cors.Options{
					AllowOriginFunc:  func(origin string) bool { return true },
					AllowedMethods:   []string{"POST", "OPTIONS"},
					AllowedHeaders:   []string{"*"},
					AllowCredentials: true,
				}).ServeHTTP(w, r, nil)
			}
		}), &http2.Server{}),
	}

	// Lambda adapter
	adapter = handlerfunc.NewV2(httpServer.Handler.ServeHTTP)
}

func Handler(ctx context.Context, req events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
	return adapter.ProxyWithContext(ctx, req)
}

func main() {
	lambda.Start(Handler)
}
