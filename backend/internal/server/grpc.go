package server

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/validate"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	applicationV1 "github.com/wzyjerry/auth/api/application/v1"
	userV1 "github.com/wzyjerry/auth/api/user/v1"
	"github.com/wzyjerry/auth/internal/conf"
	"github.com/wzyjerry/auth/internal/middleware"
	"github.com/wzyjerry/auth/internal/service/applicationService"
	"github.com/wzyjerry/auth/internal/service/userService"
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(
	c *conf.Server,
	logger log.Logger,
	register *userService.RegisterService,
	login *userService.LoginService,
	profile *userService.ProfileService,
	application *applicationService.ApplicationService,
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
	userV1.RegisterRegisterServiceServer(srv, register)
	userV1.RegisterLoginServiceServer(srv, login)
	userV1.RegisterProfileServiceServer(srv, profile)
	applicationV1.RegisterApplicationServiceServer(srv, application)
	return srv
}
