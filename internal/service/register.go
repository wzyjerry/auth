package service

import (
	"context"

	v1 "github.com/wzyjerry/auth/api/register/v1"
	"github.com/wzyjerry/auth/internal/biz"
	"google.golang.org/protobuf/types/known/emptypb"
)

type RegisterService struct {
	v1.UnimplementedRegisterServer

	uc *biz.RegisterUsecase
}

func NewRegisterService(uc *biz.RegisterUsecase) *RegisterService {
	return &RegisterService{
		uc: uc,
	}
}

func (s *RegisterService) Account(ctx context.Context, in *v1.AccountRequest) (*v1.RegisterReply, error) {
	id, err := s.uc.Account(ctx, in.Username, in.Password, in.Nickname)
	if err != nil {
		return nil, err
	}
	return &v1.RegisterReply{
		Id: id,
	}, nil
}

func (s *RegisterService) PreEmail(ctx context.Context, in *v1.PreEmailRequest) (*emptypb.Empty, error) {
	if err := s.uc.PreEmail(ctx, in.Email); err != nil {
		return nil, err
	}
	return nil, nil
}

func (s *RegisterService) Email(ctx context.Context, in *v1.EmailRequest) (*v1.RegisterReply, error) {
	id, err := s.uc.Email(ctx, in.Email, in.Code, in.Password, in.Nickname)
	if err != nil {
		return nil, err
	}
	return &v1.RegisterReply{
		Id: id,
	}, nil
}

func (s *RegisterService) PrePhone(ctx context.Context, in *v1.PrePhoneRequest) (*emptypb.Empty, error) {
	if err := s.uc.PrePhone(ctx, in.Phone); err != nil {
		return nil, err
	}
	return nil, nil
}

func (s *RegisterService) Phone(ctx context.Context, in *v1.PhoneRequest) (*v1.RegisterReply, error) {
	id, err := s.uc.Phone(ctx, in.Phone, in.Code, in.Password, in.Nickname)
	if err != nil {
		return nil, err
	}
	return &v1.RegisterReply{
		Id: id,
	}, nil
}
