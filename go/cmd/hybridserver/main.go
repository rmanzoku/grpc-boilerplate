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
	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"github.com/rmanzoku/grpc-boilerplate/go/feature/ping"
	"github.com/rmanzoku/grpc-boilerplate/go/interceptor"
	ping_service "github.com/rmanzoku/grpc-boilerplate/go/service/ping"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

var httpServer = &http.Server{}
var grpcUnixSocketPath = "/tmp/grpc.sock"

type Handler struct {
	httpHandler    http.Handler
	grpcWebHandler *grpcweb.WrappedGrpcServer
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	if h.grpcWebHandler.IsGrpcWebRequest(req) {
		h.grpcWebHandler.ServeHTTP(w, req)
		return
	}

	h.httpHandler.ServeHTTP(w, req)
}

func registerServices(s *grpc.Server) {
	ping.RegisterPingServiceServer(s, &ping_service.PingServiceServer{})
}

func main() {
	// Logging
	logger := zap.NewExample()
	grpc_logger.ReplaceGrpcLoggerV2(logger)

	// Base grpc server
	grpcServer := grpc.NewServer(grpc_middleware.WithUnaryServerChain(
		grpc_ctxtags.UnaryServerInterceptor(),
		interceptor.UnaryServerAccessLogInterceptor(),
		interceptor.UnaryServerContextInterceptor(),
		grpc_recovery.UnaryServerInterceptor(),
	))
	registerServices(grpcServer)

	go func() {
		grpcListenPort, err := net.Listen("unix", grpcUnixSocketPath)
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
		err = grpcServer.Serve(grpcListenPort)
		if err != nil {
			log.Fatal(err)
		}
	}()

	// Wrapping grpc-web server
	grpcWebServer := grpcweb.WrapServer(grpcServer)
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	grpcGateway := runtime.NewServeMux()
	err := ping.RegisterPingServiceHandlerFromEndpoint(ctx, grpcGateway, "unix:"+grpcUnixSocketPath, []grpc.DialOption{grpc.WithInsecure()})
	if err != nil {
		panic(err)
	}

	handler := &Handler{
		httpHandler:    grpcGateway,
		grpcWebHandler: grpcWebServer,
	}

	httpServer.Addr = "0.0.0.0:8080"
	httpServer.Handler = handler
	err = httpServer.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
