package server

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/validate"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/go-kratos/swagger-api/openapiv2"
	"github.com/gorilla/handlers"
	v1 "github.com/wzyjerry/auth/api/user/v1"
	"github.com/wzyjerry/auth/internal/conf"
	"github.com/wzyjerry/auth/internal/middleware"
	"github.com/wzyjerry/auth/internal/service"
)

// NewHTTPServer new a HTTP server.
func NewHTTPServer(
	c *conf.Server,
	logger log.Logger,
	register *service.RegisterService,
	login *service.LoginService,
) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
			middleware.Metadata,
			validate.Validator(),
		),
		http.Filter(handlers.CORS(
			handlers.AllowedOrigins([]string{"*"}),
			handlers.AllowedHeaders([]string{"Content-Type", "x-md-global-ip"}),
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
	v1.RegisterRegisterHTTPServer(srv, register)
	v1.RegisterLoginHTTPServer(srv, login)
	return srv
}
