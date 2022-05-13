package middleware

import (
	"context"
	"net/http"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/metadata"
	"github.com/go-kratos/kratos/v2/middleware"
)

type (
	x_md_global_ip struct{}
)

var (
	ErrIpNotFound = errors.New(http.StatusBadGateway, "IP_NOT_FOUND", "ip not found")
)

func Metadata(handler middleware.Handler) middleware.Handler {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		if md, ok := metadata.FromServerContext(ctx); ok {
			ip := md.Get("x-md-global-ip")
			if ip == "" {
				return nil, ErrIpNotFound
			}
			ctx = context.WithValue(ctx, x_md_global_ip{}, ip)
		}
		return handler(ctx, req)

	}
}

func GetIp(ctx context.Context) (string, error) {
	val := ctx.Value(x_md_global_ip{})
	if val == nil {
		return "", ErrIpNotFound
	}
	return val.(string), nil
}
