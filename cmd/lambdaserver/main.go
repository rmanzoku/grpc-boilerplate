package main

import (
	"context"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/rmanzoku/grpc-boilerplate/cmd/lambdaserver/grpcgwmux"
	"github.com/rmanzoku/grpc-boilerplate/feature/ping"
)

var adapter *grpcgwmux.GrpcGWMuxLambda

func registerServices(mux *runtime.ServeMux) (err error) {
	ctx := context.Background()
	err = ping.RegisterPingServiceHandlerServer(ctx, mux, &ping.UnimplementedPingServiceServer{})
	if err != nil {
		return
	}
	return
}

func init() {
	mux := runtime.NewServeMux()
	err := registerServices(mux)
	if err != nil {
		log.Fatal(err)
	}
	adapter = grpcgwmux.New(mux)
}

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return adapter.ProxyWithContext(ctx, req)
}

func main() {
	lambda.Start(Handler)
}
