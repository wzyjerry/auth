package applicationBiz

import "context"

type ApplicationRepo interface {
	// db部分
	Create(ctx context.Context, admin string, name string, homepage string, description *string, callback string) (string, error)
}

type ApplicationUsecase struct {
	repo ApplicationRepo
}

func NewApplicationUsecase(
	repo ApplicationRepo,
) *ApplicationUsecase {
	return &ApplicationUsecase{
		repo: repo,
	}
}

func (uc *ApplicationUsecase) Create(ctx context.Context, admin string, name string, homepage string, description *string, callback string) (string, error) {
	return uc.repo.Create(ctx, admin, name, homepage, description, callback)
}
