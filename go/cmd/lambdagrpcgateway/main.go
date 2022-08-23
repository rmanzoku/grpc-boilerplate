package main

import (
	"context"
	"log"
	"net"
	"net/http"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_logger "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/rmanzoku/grpc-boilerplate/go/feature/ping"
	interceptor "github.com/rmanzoku/grpc-boilerplate/go/middleware/grpc"
	ping_service "github.com/rmanzoku/grpc-boilerplate/go/service/ping"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const grpcAddress = "127.0.0.1:3003"
const grpcGwAddress = "127.0.0.1:3004"

var grpcServer = &grpc.Server{}
var grpcGatewayMux = &runtime.ServeMux{}

func registerServices(s *grpc.Server, m *runtime.ServeMux) (err error) {

	ctx := context.Background()
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

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
	grpcServer = grpc.NewServer(grpc_middleware.WithUnaryServerChain(
		grpc_ctxtags.UnaryServerInterceptor(),
		interceptor.UnaryServerAccessLogInterceptor(),
		//		interceptor.UnaryServerContextInterceptor(),
		grpc_recovery.UnaryServerInterceptor(),
	))

	// Grcp Gateway
	grpcGatewayMux = runtime.NewServeMux()

	err := registerServices(grpcServer, grpcGatewayMux)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	GrpcListenPort, err := net.Listen("tcp", grpcAddress)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	go func() {
		err = grpcServer.Serve(GrpcListenPort)
		if err != nil {
			log.Fatal(err)
		}
	}()

	if err := http.ListenAndServe(grpcGwAddress, grpcGatewayMux); err != nil {
		log.Fatal("err")
	}
}
