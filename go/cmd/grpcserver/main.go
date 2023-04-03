package main

import (
	"net"
	"os"
	"os/signal"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"github.com/rmanzoku/grpc-boilerplate/go/feature/ping"
	ping_service "github.com/rmanzoku/grpc-boilerplate/go/service/ping"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func registerServices(s *grpc.Server) {
	ping.RegisterPingServiceServer(s, &ping_service.PingServiceServer{})
}

func init() {
}

func StartServer() int {
	logger := zap.NewExample()
	grpc_zap.ReplaceGrpcLoggerV2(logger)
	grpcServer := grpc.NewServer(grpc_middleware.WithUnaryServerChain(
		grpc_ctxtags.UnaryServerInterceptor(),
		grpc_zap.UnaryServerInterceptor(logger),
		grpc_recovery.UnaryServerInterceptor(),
	))
	registerServices(grpcServer)

	port := os.Getenv("GRPC_PORT")
	l, err := net.Listen("tcp", ":"+port)
	if err != nil {
		logger.Error(err.Error())
		return 1
	}

	go func() {
		logger.Info("start grpc server port: " + port)
		if err := grpcServer.Serve(l); err != nil {
			logger.Error(err.Error())
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	logger.Info("stopping grpc server...")
	grpcServer.GracefulStop()
	return 0
}

func main() {
	os.Exit(StartServer())
}
