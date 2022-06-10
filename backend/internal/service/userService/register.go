package userService

import (
	"context"

	v1 "github.com/wzyjerry/auth/api/user/v1"
	"github.com/wzyjerry/auth/internal/biz/userBiz"
	"google.golang.org/protobuf/types/known/emptypb"
)

type RegisterService struct {
	v1.UnimplementedRegisterServer

	uc *userBiz.RegisterUsecase
}

func NewRegisterService(uc *userBiz.RegisterUsecase) *RegisterService {
	return &RegisterService{
		uc: uc,
	}
}

func (s *RegisterService) Account(ctx context.Context, in *v1.RegisterAccountRequest) (*v1.RegisterReply, error) {
	id, err := s.uc.Account(ctx, in.Username, in.Password, in.Nickname)
	if err != nil {
		return nil, err
	}
	return &v1.RegisterReply{
		Id: id,
	}, nil
}

func (s *RegisterService) PreEmail(ctx context.Context, in *v1.RegisterPreEmailRequest) (*emptypb.Empty, error) {
	return nil, s.uc.PreEmail(ctx, in.Email)
}

func (s *RegisterService) Email(ctx context.Context, in *v1.RegisterEmailRequest) (*v1.RegisterReply, error) {
	id, err := s.uc.EmailPassword(ctx, in.Email, in.Code, in.Password, in.Nickname)
	if err != nil {
		return nil, err
	}
	return &v1.RegisterReply{
		Id: id,
	}, nil
}

func (s *RegisterService) PrePhone(ctx context.Context, in *v1.RegisterPrePhoneRequest) (*emptypb.Empty, error) {
	return nil, s.uc.PrePhone(ctx, in.Phone)
}

func (s *RegisterService) Phone(ctx context.Context, in *v1.RegisterPhoneRequest) (*v1.RegisterReply, error) {
	id, err := s.uc.PhonePassword(ctx, in.Phone, in.Code, in.Password, in.Nickname)
	if err != nil {
		return nil, err
	}
	return &v1.RegisterReply{
		Id: id,
	}, nil
}
