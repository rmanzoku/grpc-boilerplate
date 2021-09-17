package grpcutil

import (
	"context"
	"net/http"
	"strings"

	"google.golang.org/grpc/metadata"
)

func ExtractHeaders(ctx context.Context) metadata.MD {
	headers, _ := metadata.FromIncomingContext(ctx)
	return headers
}

func ExtractCookies(ctx context.Context) []*http.Cookie {
	return ParseCookie(ExtractHeaders(ctx))
}

func ParseCookie(headers metadata.MD) []*http.Cookie {
	cookieList := append(headers["grpcgateway-cookie"], headers["grpc-cookie"]...)
	cookieList = append(cookieList, headers["cookie"]...)
	rawCookies := strings.Join(cookieList, ",")
	header := http.Header{}
	header.Add("Cookie", rawCookies)
	request := http.Request{Header: header}
	return request.Cookies()
}
