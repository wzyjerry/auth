package service

import (
	"github.com/google/wire"
	"github.com/wzyjerry/auth/internal/service/applicationService"
	"github.com/wzyjerry/auth/internal/service/oauth2Service"
	"github.com/wzyjerry/auth/internal/service/userService"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(
	userService.NewRegisterService,
	userService.NewLoginService,
	userService.NewProfileService,
	applicationService.NewApplicationService,
	oauth2Service.NewOAuth2Service,
)
