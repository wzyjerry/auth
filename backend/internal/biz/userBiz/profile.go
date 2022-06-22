package userBiz

import (
	"context"
)

type ProfileUsecase struct {
	repo UserRepo
}

func NewProfileUsecase(
	repo UserRepo,
) *ProfileUsecase {
	return &ProfileUsecase{
		repo: repo,
	}
}

func (uc *ProfileUsecase) GetAvatar(ctx context.Context, id string) (*string, error) {
	return uc.repo.GetAvatar(ctx, id)
}
