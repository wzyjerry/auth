package service

import (
	"github.com/google/wire"
	"github.com/wzyjerry/auth/internal/service/applicationService"
	"github.com/wzyjerry/auth/internal/service/userService"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(
	userService.NewRegisterService,
	userService.NewLoginService,
	userService.NewProfileService,
	applicationService.NewApplicationService,
)
