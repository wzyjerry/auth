package server

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/validate"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/go-kratos/swagger-api/openapiv2"
	"github.com/gorilla/handlers"
	applicationV1 "github.com/wzyjerry/auth/api/application/v1"
	oauth2V1 "github.com/wzyjerry/auth/api/oauth2/v1"
	userV1 "github.com/wzyjerry/auth/api/user/v1"
	"github.com/wzyjerry/auth/internal/conf"
	"github.com/wzyjerry/auth/internal/middleware"
	"github.com/wzyjerry/auth/internal/service/applicationService"
	"github.com/wzyjerry/auth/internal/service/oauth2Service"
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
	oauth2 *oauth2Service.OAuth2Service,
) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
			middleware.Metadata,
			validate.Validator(),
		),
		http.Filter(handlers.CORS(
			handlers.AllowedMethods([]string{"POST", "PUT", "GET", "DELETE"}),
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
	userV1.RegisterRegisterServiceHTTPServer(srv, register)
	userV1.RegisterLoginServiceHTTPServer(srv, login)
	userV1.RegisterProfileServiceHTTPServer(srv, profile)
	applicationV1.RegisterApplicationServiceHTTPServer(srv, application)
	oauth2V1.RegisterOAuth2ServiceHTTPServer(srv, oauth2)
	return srv
}
