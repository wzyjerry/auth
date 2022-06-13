package data

import (
	"context"

	"github.com/wzyjerry/auth/internal/biz/applicationBiz"
	"github.com/wzyjerry/auth/internal/ent"
	"github.com/wzyjerry/auth/internal/ent/application"
	"github.com/wzyjerry/auth/internal/ent/avatar"
	"github.com/wzyjerry/auth/internal/ent/schema/applicationNested"
	"github.com/wzyjerry/auth/internal/ent/schema/avatarNested"
	"github.com/wzyjerry/auth/internal/util"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type applicationRepo struct {
	data *Data
}

func (r *applicationRepo) GenerateClientSecret(ctx context.Context, admin string, id string, description string) (*applicationNested.ClientSecret, error) {
	secret := &applicationNested.ClientSecret{
		Generated:   timestamppb.Now(),
		Description: &description,
		Secret:      util.P(util.NewUUID()),
	}
	if err := r.data.db.WithTx(ctx, func(tx *ent.Tx) error {
		application, err := tx.Application.Query().Where(application.AdminEQ(admin), application.IDEQ(id)).Select(application.FieldClientSecrets).Only(ctx)
		if err != nil {
			return err
		}
		application.ClientSecrets = append([]*applicationNested.ClientSecret{secret}, application.ClientSecrets...)
		return r.data.db.Application.UpdateOneID(id).SetClientSecrets(application.ClientSecrets).Exec(ctx)
	}); err != nil {
		return nil, err
	}
	return secret, nil
}

func (r *applicationRepo) GetAvatar(ctx context.Context, id string) (string, error) {
	avatar, err := r.data.db.Avatar.Query().Where(avatar.KindEQ(int32(avatarNested.Kind_KIND_APPLICATION)), avatar.RelIDEQ(id)).Only(ctx)
	if err != nil {
		return "", err
	}
	return *avatar.Avatar, nil
}

func (r *applicationRepo) Retrieve(ctx context.Context, admin string, id string) (*ent.Application, error) {
	return r.data.db.Application.Query().Where(application.AdminEQ(admin), application.IDEQ(id)).Only(ctx)
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
