package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/httpadapter"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/rmanzoku/grpc-boilerplate/feature/ping"
)

var adapter *httpadapter.HandlerAdapter

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
	adapter = httpadapter.New(accessLogMiddleware(mux))
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
