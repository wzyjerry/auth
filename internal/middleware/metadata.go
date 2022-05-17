package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/lestrrat-go/jwx/v2/jwt"
	"github.com/wzyjerry/auth/internal/util"
)

type (
	x_md_global_ip struct{}
	auth           struct{}
	TokenType      int
)

const (
	AuthToken TokenType = iota
	UserToken
	ClientToken
)

const (
	bearerWord       = "Bearer"
	ipKey            = "x-md-global-ip"
	authorizationKey = "Authorization"
)

var (
	ErrIpNotFound   = errors.New(http.StatusBadGateway, "IP_NOT_FOUND", "ip not found")
	ErrUnauthorized = errors.New(http.StatusUnauthorized, "UNAUTHORIZED", "need authorization header")
	ErrForbidden    = errors.New(http.StatusForbidden, "FORBIDDEN", "token invalid")
)

func Metadata(handler middleware.Handler) middleware.Handler {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		if tr, ok := transport.FromServerContext(ctx); ok {
			md := tr.RequestHeader()
			ip := md.Get(ipKey)
			if ip == "" {
				return nil, ErrIpNotFound
			}
			ctx = context.WithValue(ctx, x_md_global_ip{}, ip)
			authorization := md.Get(authorizationKey)
			if authorization != "" {
				ctx = context.WithValue(ctx, auth{}, authorization)
			}
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

func Validator(ctx context.Context, helper *util.TokenHelper, tokenType TokenType) (*jwt.Token, error) {
	val := ctx.Value(auth{})
	if val == nil {
		return nil, ErrUnauthorized
	}
	split := strings.SplitN(val.(string), " ", 2)
	if len(split) != 2 || strings.Compare(split[0], bearerWord) != 0 {
		return nil, ErrForbidden
	}
	token, err := helper.ParseJWT(split[1])
	if err != nil || helper.IsIdToken(token) {
		return nil, ErrForbidden
	}
	switch tokenType {
	case AuthToken:
		if helper.IsAuthToken(token) {
			return token, nil
		}
	case UserToken:
		if helper.IsUserToken(token) {
			return token, nil
		}
	case ClientToken:
		if helper.IsClientToken(token) {
			return token, nil
		}
	}
	return nil, ErrForbidden
}
