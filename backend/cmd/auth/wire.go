//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/wzyjerry/auth/internal/biz"
	"github.com/wzyjerry/auth/internal/biz/third_party"
	"github.com/wzyjerry/auth/internal/conf"
	"github.com/wzyjerry/auth/internal/data"
	"github.com/wzyjerry/auth/internal/server"
	"github.com/wzyjerry/auth/internal/service"
	"github.com/wzyjerry/auth/internal/util"
)

// wireApp init kratos application.
func wireApp(
	*conf.Server,
	*conf.Data,
	*conf.Security,
	log.Logger,
) (*kratos.App, func(), error) {
	panic(wire.Build(third_party.ProviderSet, server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, util.ProviderSet, newApp))
}
