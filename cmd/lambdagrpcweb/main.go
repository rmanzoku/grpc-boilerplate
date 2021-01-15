package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/httpadapter"
	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"github.com/rmanzoku/grpc-boilerplate/feature/ping"
	ping_service "github.com/rmanzoku/grpc-boilerplate/service/ping"
	"google.golang.org/grpc"
)

var adapter *httpadapter.HandlerAdapter

func registerServices(s *grpc.Server) {
	ping.RegisterPingServiceServer(s, &ping_service.PingServiceServer{})
}

func init() {
	s := grpc.NewServer()
	registerServices(s)
	wrappedGrpc := grpcweb.WrapServer(s)
	adapter = httpadapter.New(accessLogMiddleware(wrappedGrpc))
}

func accessLogMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		method := r.Method
		path := r.URL.Path
		fmt.Printf("method:%s\tpath:%s\n", method, path)
		h.ServeHTTP(w, r)
	})
}

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return adapter.ProxyWithContext(ctx, req)
}

func main() {
	lambda.Start(Handler)
}
