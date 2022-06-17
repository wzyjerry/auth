package userBiz

import (
	"context"

	"github.com/wzyjerry/auth/internal/ent"
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
	avatar, err := uc.repo.GetAvatar(ctx, id)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, nil
		} else {
			return nil, err
		}
	}
	return avatar, nil
}
