package data

import (
	"context"

	"github.com/wzyjerry/auth/internal/biz/applicationBiz"
	"github.com/wzyjerry/auth/internal/util"
)

type applicationRepo struct {
	data *Data
}

func NewApplicationRepo(
	data *Data,
) applicationBiz.ApplicationRepo {
	return &applicationRepo{
		data: data,
	}
}

func (r *applicationRepo) Create(ctx context.Context, admin string, name string, homepage string, description *string, callback string) (string, error) {
	application, err := r.data.db.Application.Create().
		SetID(r.data.db.GenerateId()).
		SetAdmin(admin).
		SetName(name).
		SetHomepage(homepage).
		SetNillableDescription(description).
		SetClientID(util.NewUUID()).
		SetCallback(callback).
		Save(ctx)
	if err != nil {
		return "", err
	}
	return application.ID, nil
}
