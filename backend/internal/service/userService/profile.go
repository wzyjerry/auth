package userService

import (
	"context"

	v1 "github.com/wzyjerry/auth/api/user/v1"
	"github.com/wzyjerry/auth/internal/biz/userBiz"
	"github.com/wzyjerry/auth/internal/conf"
	"github.com/wzyjerry/auth/internal/middleware"
	"github.com/wzyjerry/auth/internal/util"
	"google.golang.org/protobuf/types/known/emptypb"
)

type ProfileService struct {
	v1.UnimplementedProfileServiceServer

	uc     *userBiz.ProfileUsecase
	conf   *conf.Security
	helper *util.TokenHelper
}

func NewProfileService(
	uc *userBiz.ProfileUsecase,
	conf *conf.Security,
	helper *util.TokenHelper,
) *ProfileService {
	return &ProfileService{
		uc:     uc,
		conf:   conf,
		helper: helper,
	}
}

func (s *ProfileService) GetAvatar(ctx context.Context, _ *emptypb.Empty) (*v1.GetAvatarReply, error) {
	token, err := middleware.Validator(ctx, s.helper, middleware.AuthToken)
	if err != nil {
		return nil, err
	}
	avatar, err := s.uc.GetAvatar(ctx, (*token).Subject())
	if err != nil {
		return nil, err
	}
	return &v1.GetAvatarReply{
		Avatar: avatar,
	}, nil
}
