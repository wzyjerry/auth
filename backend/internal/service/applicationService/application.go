package applicationService

import (
	"context"

	v1 "github.com/wzyjerry/auth/api/application/v1"
	"github.com/wzyjerry/auth/internal/biz/applicationBiz"
	"github.com/wzyjerry/auth/internal/conf"
	"github.com/wzyjerry/auth/internal/middleware"
	"github.com/wzyjerry/auth/internal/util"
)

type ApplicationService struct {
	v1.UnimplementedApplicationServer

	uc     *applicationBiz.ApplicationUsecase
	conf   *conf.Security
	helper *util.TokenHelper
}

func NewApplicationService(
	uc *applicationBiz.ApplicationUsecase,
	conf *conf.Security,
	helper *util.TokenHelper,
) *ApplicationService {
	return &ApplicationService{
		uc:     uc,
		conf:   conf,
		helper: helper,
	}
}

func (s *ApplicationService) Create(ctx context.Context, in *v1.CreateRequest) (*v1.CreateReply, error) {
	token, err := middleware.Validator(ctx, s.helper, middleware.UserToken)
	if err != nil {
		return nil, err
	}
	id, err := s.uc.Create(ctx, (*token).Subject(), in.Name, in.Homepage, in.Description, in.Callback)
	if err != nil {
		return nil, err
	}
	return &v1.CreateReply{
		Id: id,
	}, nil
}
