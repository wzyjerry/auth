package server

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/validate"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	v1 "github.com/wzyjerry/auth/api/user/v1"
	"github.com/wzyjerry/auth/internal/conf"
	"github.com/wzyjerry/auth/internal/middleware"
	"github.com/wzyjerry/auth/internal/service"
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(
	c *conf.Server,
	logger log.Logger,
	register *service.RegisterService,
	login *service.LoginService,
) *grpc.Server {
	var opts = []grpc.ServerOption{
		grpc.Middleware(
			recovery.Recovery(),
			middleware.Metadata,
			validate.Validator(),
		),
	}
	if c.Grpc.Addr != "" {
		opts = append(opts, grpc.Address(c.Grpc.Addr))
	}
	if c.Grpc.Timeout != nil {
		opts = append(opts, grpc.Timeout(c.Grpc.Timeout.AsDuration()))
	}
	srv := grpc.NewServer(opts...)
	v1.RegisterRegisterServer(srv, register)
	v1.RegisterLoginServer(srv, login)
	return srv
}
