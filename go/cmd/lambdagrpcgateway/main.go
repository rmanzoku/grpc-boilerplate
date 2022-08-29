package main

// grpcとgrpc gatewayを同時に起動する
//
// |             |    |                                                    |
// | API Gateway | -> |                    AWS Lambda                      |
// |             |    | HandlerFuncAdapter -> grpcGatewayMux -> grpcServer |
//

import (
	"context"
	"log"
	"net"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/handlerfunc"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_logger "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/rmanzoku/grpc-boilerplate/go/feature/ping"
	"github.com/rmanzoku/grpc-boilerplate/go/interceptor"
	ping_service "github.com/rmanzoku/grpc-boilerplate/go/service/ping"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const grpcAddress = "127.0.0.1:3003"

var adapter *handlerfunc.HandlerFuncAdapterV2

func registerServices(s *grpc.Server, m *runtime.ServeMux) (err error) {

	ctx := context.Background()
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	// Serviceはgrpc serverとgateway向けにRegisterが必要
	// PingService
	ping.RegisterPingServiceServer(s, &ping_service.PingServiceServer{})
	err = ping.RegisterPingServiceHandlerFromEndpoint(ctx, m, grpcAddress, opts)
	if err != nil {
		return err
	}

	return
}

func init() {

	logger := zap.NewExample()
	grpc_logger.ReplaceGrpcLoggerV2(logger)

	// Base grpc server
	grpcServer := grpc.NewServer(grpc_middleware.WithUnaryServerChain(
		grpc_ctxtags.UnaryServerInterceptor(),
		interceptor.UnaryServerAccessLogInterceptor(),
		//		interceptor.UnaryServerContextInterceptor(),
		grpc_recovery.UnaryServerInterceptor(),
	))

	grpcGatewayMux := runtime.NewServeMux()
	err := registerServices(grpcServer, grpcGatewayMux)
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		grpcListenPort, err := net.Listen("tcp", grpcAddress)
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
		err = grpcServer.Serve(grpcListenPort)
		if err != nil {
			log.Fatal(err)
		}
	}()

	adapter = handlerfunc.NewV2(grpcGatewayMux.ServeHTTP)
}

func Handler(ctx context.Context, req events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
	return adapter.ProxyWithContext(ctx, req)
}

func main() {
	lambda.Start(Handler)
}
