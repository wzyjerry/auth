package applicationService

import (
	"context"

	v1 "github.com/wzyjerry/auth/api/application/v1"
	"github.com/wzyjerry/auth/internal/biz/applicationBiz"
	"github.com/wzyjerry/auth/internal/conf"
	"github.com/wzyjerry/auth/internal/middleware"
	"github.com/wzyjerry/auth/internal/util"
	"google.golang.org/protobuf/types/known/emptypb"
)

type ApplicationService struct {
	v1.UnimplementedApplicationServiceServer

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

func (s *ApplicationService) Retrieve(ctx context.Context, in *v1.RetrieveRequest) (*v1.RetrieveReply, error) {
	token, err := middleware.Validator(ctx, s.helper, middleware.UserToken)
	if err != nil {
		return nil, err
	}
	return s.uc.Retrieve(ctx, (*token).Subject(), in.Id)
}

func (s *ApplicationService) GenerateClientSecret(ctx context.Context, in *v1.GenerateClientSecretRequest) (*v1.GenerateClientSecretReply, error) {
	token, err := middleware.Validator(ctx, s.helper, middleware.UserToken)
	if err != nil {
		return nil, err
	}
	return s.uc.GenerateClientSecret(ctx, (*token).Subject(), in.Id, in.Description)
}

func (s *ApplicationService) RevokeClientSecret(ctx context.Context, in *v1.RevokeClientSecretRequest) (*emptypb.Empty, error) {
	token, err := middleware.Validator(ctx, s.helper, middleware.UserToken)
	if err != nil {
		return nil, err
	}
	return nil, s.uc.RevokeClientSecret(ctx, (*token).Subject(), in.Id, in.SecretId)
}

func (s *ApplicationService) UploadLogo(ctx context.Context, in *v1.UploadLogoRequest) (*emptypb.Empty, error) {
	token, err := middleware.Validator(ctx, s.helper, middleware.UserToken)
	if err != nil {
		return nil, err
	}
	return nil, s.uc.UploadLogo(ctx, (*token).Subject(), in.Id, in.Logo)
}

func (s *ApplicationService) Update(ctx context.Context, in *v1.UpdateRequest) (*emptypb.Empty, error) {
	token, err := middleware.Validator(ctx, s.helper, middleware.UserToken)
	if err != nil {
		return nil, err
	}
	return nil, s.uc.Update(ctx, (*token).Subject(), in.Id, in.Name, in.Homepage, in.Description, in.Callback)
}

func (s *ApplicationService) Delete(ctx context.Context, in *v1.DeleteRequest) (*emptypb.Empty, error) {
	token, err := middleware.Validator(ctx, s.helper, middleware.UserToken)
	if err != nil {
		return nil, err
	}
	return nil, s.uc.Delete(ctx, (*token).Subject(), in.Id)
}

func (s *ApplicationService) GetAll(ctx context.Context, _ *emptypb.Empty) (*v1.GetAllReply, error) {
	token, err := middleware.Validator(ctx, s.helper, middleware.UserToken)
	if err != nil {
		return nil, err
	}
	return s.uc.GetAll(ctx, (*token).Subject())
}
