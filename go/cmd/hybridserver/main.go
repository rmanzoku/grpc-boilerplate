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
	"github.com/rmanzoku/grpc-boilerplate/go/feature/healthcheck"
	"github.com/rmanzoku/grpc-boilerplate/go/feature/ping"
	"github.com/rmanzoku/grpc-boilerplate/go/interceptor"
	healthcheck_service "github.com/rmanzoku/grpc-boilerplate/go/service/healthcheck"
	ping_service "github.com/rmanzoku/grpc-boilerplate/go/service/ping"
	"github.com/rmanzoku/grpc-boilerplate/go/utility/env"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

var httpServer = &http.Server{}
var grpcUnixSocketPath = "/tmp/grpc.sock"

type Handler struct {
	grpcGatewayRuntime *runtime.ServeMux
	grpcWebHandler     *grpcweb.WrappedGrpcServer
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	// if req.URL.Path == "/" {
	// 	w.Write([]byte("ok"))
	// 	return
	// }

	if h.grpcWebHandler.IsGrpcWebRequest(req) {
		h.grpcWebHandler.ServeHTTP(w, req)
		return
	}
	h.grpcGatewayRuntime.ServeHTTP(w, req)
}

func registerServices(h *Handler, s *grpc.Server) {
	ctx := context.TODO()
	var err error

	ping.RegisterPingServiceServer(s, &ping_service.PingServiceServer{})
	err = ping.RegisterPingServiceHandlerFromEndpoint(ctx, h.grpcGatewayRuntime, "unix:"+grpcUnixSocketPath, []grpc.DialOption{grpc.WithInsecure()})
	if err != nil {
		panic(err)
	}

	healthcheck.RegisterHealthcheckServiceServer(s, &healthcheck_service.HealthcheckServiceServer{})
	err = healthcheck.RegisterHealthcheckServiceHandlerFromEndpoint(ctx, h.grpcGatewayRuntime, "unix:"+grpcUnixSocketPath, []grpc.DialOption{grpc.WithInsecure()})
	if err != nil {
		panic(err)
	}
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

	handler := &Handler{
		grpcGatewayRuntime: runtime.NewServeMux(),
		grpcWebHandler:     grpcweb.WrapServer(grpcServer),
	}

	registerServices(handler, grpcServer)

	go func() {
		grpcListenPort, err := net.Listen("unix", grpcUnixSocketPath)
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
		defer grpcListenPort.Close()

		err = grpcServer.Serve(grpcListenPort)
		if err != nil {
			log.Fatal(err)
		}

		defer grpcListenPort.Close()
	}()

	httpServer.Addr = "0.0.0.0:" + env.GetWithDefault("PORT", "8080")
	httpServer.Handler = handler
	err := httpServer.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
