// Packge echolambda add Echo support for the aws-severless-go-api library.
// Uses the core package behind the scenes and exposes the New method to
// get a new instance and Proxy method to send request to the echo.Echo
package grpcgwmux

import (
	"context"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/awslabs/aws-lambda-go-api-proxy/core"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
)

// EchoLambda makes it easy to send API Gateway proxy events to a echo.Echo.
// The library transforms the proxy event into an HTTP request and then
// creates a proxy response object from the http.ResponseWriter
type GrpcGWMuxLambda struct {
	core.RequestAccessor

	Mux *runtime.ServeMux
}

// New creates a new instance of the EchoLambda object.
// Receives an initialized *echo.Echo object - normally created with echo.New().
// It returns the initialized instance of the EchoLambda object.
func New(m *runtime.ServeMux) *GrpcGWMuxLambda {
	return &GrpcGWMuxLambda{Mux: m}
}

// Proxy receives an API Gateway proxy event, transforms it into an http.Request
// object, and sends it to the echo.Echo for routing.
// It returns a proxy response object generated from the http.ResponseWriter.
func (g *GrpcGWMuxLambda) Proxy(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	muxRequest, err := g.ProxyEventToHTTPRequest(req)
	return g.proxyInternal(muxRequest, err)
}

// ProxyWithContext receives context and an API Gateway proxy event,
// transforms them into an http.Request object, and sends it to the echo.Echo for routing.
// It returns a proxy response object generated from the http.ResponseWriter.
func (g *GrpcGWMuxLambda) ProxyWithContext(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	muxRequest, err := g.EventToRequestWithContext(ctx, req)
	return g.proxyInternal(muxRequest, err)
}

func (g *GrpcGWMuxLambda) proxyInternal(req *http.Request, err error) (events.APIGatewayProxyResponse, error) {

	if err != nil {
		return core.GatewayTimeout(), core.NewLoggedError("Could not convert proxy event to request: %v", err)
	}

	respWriter := core.NewProxyResponseWriter()
	g.Mux.ServeHTTP(http.ResponseWriter(respWriter), req)

	proxyResponse, err := respWriter.GetProxyResponse()
	if err != nil {
		return core.GatewayTimeout(), core.NewLoggedError("Error while generating proxy response: %v", err)
	}

	return proxyResponse, nil
}
