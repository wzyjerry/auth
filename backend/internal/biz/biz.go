package biz

import (
	"github.com/google/wire"
	"github.com/wzyjerry/auth/internal/biz/applicationBiz"
	"github.com/wzyjerry/auth/internal/biz/oauth2Biz"
	"github.com/wzyjerry/auth/internal/biz/userBiz"
)

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(
	userBiz.NewRegisterUsecase,
	userBiz.NewLoginUsecase,
	userBiz.NewProfileUsecase,
	applicationBiz.NewApplicationUsecase,
	oauth2Biz.NewOAuth2Usecase,
)
