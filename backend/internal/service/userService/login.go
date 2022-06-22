package userService

import (
	"context"
	"os"

	v1 "github.com/wzyjerry/auth/api/user/v1"
	"github.com/wzyjerry/auth/internal/biz/userBiz"
	"github.com/wzyjerry/auth/internal/conf"
	"github.com/wzyjerry/auth/internal/util"
	"google.golang.org/protobuf/types/known/emptypb"
)

type LoginService struct {
	v1.UnimplementedLoginServiceServer

	uc     *userBiz.LoginUsecase
	conf   *conf.Security
	helper *util.TokenHelper
}

func NewLoginService(
	uc *userBiz.LoginUsecase,
	conf *conf.Security,
	helper *util.TokenHelper,
) *LoginService {
	return &LoginService{
		uc:     uc,
		conf:   conf,
		helper: helper,
	}
}

func (s *LoginService) PrePhone(ctx context.Context, in *v1.LoginPrePhoneRequest) (*emptypb.Empty, error) {
	return nil, s.uc.PrePhone(ctx, in.Phone)
}

func (s *LoginService) PreEmail(ctx context.Context, in *v1.LoginPreEmailRequest) (*emptypb.Empty, error) {
	return nil, s.uc.PreEmail(ctx, in.Email)
}

func (s *LoginService) Login(ctx context.Context, in *v1.LoginRequest) (*v1.LoginReply, error) {
	user, err := s.uc.Login(ctx, in.Type, in.Method, in.Unique, in.Secret)
	if err != nil {
		return nil, err
	}
	token, err := s.uc.GenerateAccessToken(os.Getenv(s.conf.ClientId), user)
	if err != nil {
		return nil, err
	}
	return &v1.LoginReply{
		Token: token,
	}, nil
}
