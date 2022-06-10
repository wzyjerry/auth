package server

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/validate"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/go-kratos/swagger-api/openapiv2"
	"github.com/gorilla/handlers"
	applicationV1 "github.com/wzyjerry/auth/api/application/v1"
	userV1 "github.com/wzyjerry/auth/api/user/v1"
	"github.com/wzyjerry/auth/internal/conf"
	"github.com/wzyjerry/auth/internal/middleware"
	"github.com/wzyjerry/auth/internal/service/applicationService"
	"github.com/wzyjerry/auth/internal/service/userService"
)

// NewHTTPServer new a HTTP server.
func NewHTTPServer(
	c *conf.Server,
	logger log.Logger,
	register *userService.RegisterService,
	login *userService.LoginService,
	profile *userService.ProfileService,
	application *applicationService.ApplicationService,
) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
			middleware.Metadata,
			validate.Validator(),
		),
		http.Filter(handlers.CORS(
			handlers.AllowedOrigins([]string{"*"}),
			handlers.AllowedHeaders([]string{"Content-Type", "x-md-global-ip", "Authorization"}),
		)),
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)
	openAPIhandler := openapiv2.NewHandler()
	srv.HandlePrefix("/q/", openAPIhandler)
	userV1.RegisterRegisterHTTPServer(srv, register)
	userV1.RegisterLoginHTTPServer(srv, login)
	userV1.RegisterProfileHTTPServer(srv, profile)
	applicationV1.RegisterApplicationHTTPServer(srv, application)
	return srv
}
