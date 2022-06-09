package service

import (
	"context"
	"os"

	"github.com/lestrrat-go/jwx/v2/jwt"
	v1 "github.com/wzyjerry/auth/api/user/v1"
	"github.com/wzyjerry/auth/internal/biz"
	"github.com/wzyjerry/auth/internal/conf"
	"github.com/wzyjerry/auth/internal/middleware"
	"github.com/wzyjerry/auth/internal/util"
	"google.golang.org/protobuf/types/known/emptypb"
)

type LoginService struct {
	v1.UnimplementedLoginServer

	uc     *biz.LoginUsecase
	conf   *conf.Security
	helper *util.TokenHelper
}

func NewLoginService(
	uc *biz.LoginUsecase,
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

func (s *LoginService) Trash(ctx context.Context, _ *emptypb.Empty) (*v1.TrashReply, error) {
	var token *jwt.Token
	var err error
	var sub string
	token, err = middleware.Validator(ctx, s.helper, middleware.AuthToken)
	authToken := (err == nil)
	if sub == "" && authToken {
		sub = (*token).Subject()
	}
	token, err = middleware.Validator(ctx, s.helper, middleware.UserToken)
	userToken := (err == nil)
	if sub == "" && userToken {
		sub = (*token).Subject()
	}
	token, err = middleware.Validator(ctx, s.helper, middleware.ClientToken)
	clientToken := (err == nil)
	if sub == "" && clientToken {
		sub = (*token).Subject()
	}
	return &v1.TrashReply{
		AuthToken:   authToken,
		UserToken:   userToken,
		ClientToken: clientToken,
		Sub:         sub,
	}, nil

}
