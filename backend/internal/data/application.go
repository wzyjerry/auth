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
)

type applicationRepo struct {
	data *Data
}

func (r *applicationRepo) GetAll(ctx context.Context, admin string) ([]*ent.Application, error) {
	return r.data.db.Application.Query().Where(application.AdminEQ(admin)).Select(
		application.FieldID,
		application.FieldName,
		application.FieldDescription,
	).Order(ent.Desc(application.FieldID)).All(ctx)
}

func (r *applicationRepo) GetLogoMap(ctx context.Context, ids []string) (map[string]string, error) {
	result := make(map[string]string)
	avatar, err := r.data.db.Avatar.Query().Where(avatar.KindEQ(int32(avatarNested.Kind_KIND_APPLICATION)), avatar.RelIDIn(ids...)).All(ctx)
	if err != nil {
		return nil, err
	}
	for _, item := range avatar {
		result[*item.RelID] = *item.Avatar
	}
	return result, nil
}

func (r *applicationRepo) Delete(ctx context.Context, admin string, id string) error {
	_, err := r.data.db.Application.Query().Where(application.AdminEQ(admin), application.IDEQ(id)).Select(application.FieldID).Only(ctx)
	if err != nil {
		return err
	}
	go func() {
		_, _ = r.data.db.Avatar.Delete().Where(avatar.KindEQ(int32(avatarNested.Kind_KIND_APPLICATION)), avatar.RelIDEQ(id)).Exec(ctx)
	}()
	return r.data.db.Application.DeleteOneID(id).Exec(ctx)
}

func (r *applicationRepo) Update(ctx context.Context, admin string, id string, name string, homepage string, description *string, callback string) error {
	application, err := r.data.db.Application.Query().Where(application.AdminEQ(admin), application.IDEQ(id)).Select(application.FieldID).Only(ctx)
	if err != nil {
		return err
	}
	update := application.Update().SetName(name).SetHomepage(homepage).SetCallback(callback)
	if description == nil {
		update.ClearDescription()
	} else {
		update.SetNillableDescription(description)
	}
	return update.Exec(ctx)
}

func (r *applicationRepo) SetLogo(ctx context.Context, admin string, id string, logo string) error {
	exist, err := r.data.db.Application.Query().Where(application.AdminEQ(admin), application.IDEQ(id)).Select(application.FieldID).Exist(ctx)
	if err != nil {
		return err
	}
	if !exist {
		return applicationBiz.ErrApplicationNotFound
	}
	avatar, err := r.data.db.Avatar.Query().Where(avatar.KindEQ(int32(avatarNested.Kind_KIND_APPLICATION)), avatar.RelIDEQ(id)).Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return r.data.db.Avatar.Create().
				SetID(r.data.db.GenerateId()).
				SetKind(int32(avatarNested.Kind_KIND_APPLICATION)).
				SetRelID(id).
				SetAvatar(logo).Exec(ctx)
		}
		return err
	}
	return avatar.Update().SetAvatar(logo).Exec(ctx)
}

func (r *applicationRepo) RevokeClientSecret(ctx context.Context, admin string, id string, secretId string) error {
	application, err := r.data.db.Application.Query().Where(application.AdminEQ(admin), application.IDEQ(id)).Select(application.FieldClientSecrets).Only(ctx)
	if err != nil {
		return err
	}
	clientSecrets := make([]*applicationNested.ClientSecret, 0, len(application.ClientSecrets))
	for _, item := range application.ClientSecrets {
		if *item.Id != secretId {
			clientSecrets = append(clientSecrets, item)
		}
	}
	if len(clientSecrets) > 0 {
		return r.data.db.Application.UpdateOneID(id).SetClientSecrets(clientSecrets).Exec(ctx)
	}
	return applicationBiz.ErrRevokeBadRequest
}

func (r *applicationRepo) GenerateClientSecret(ctx context.Context, admin string, id string, description string) (*applicationNested.ClientSecret, error) {
	secret := &applicationNested.ClientSecret{
		Id:          util.P(r.data.db.GenerateId()),
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

func (r *applicationRepo) GetLogo(ctx context.Context, id string) (*string, error) {
	avatar, err := r.data.db.Avatar.Query().Where(avatar.KindEQ(int32(avatarNested.Kind_KIND_APPLICATION)), avatar.RelIDEQ(id)).Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, nil
		}
		return nil, err
	}
	return avatar.Avatar, nil
}

func (r *applicationRepo) Retrieve(ctx context.Context, admin string, id string) (*ent.Application, error) {
	return r.data.db.Application.Query().Where(application.AdminEQ(admin), application.IDEQ(id)).Only(ctx)
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

func NewApplicationRepo(
	data *Data,
) applicationBiz.ApplicationRepo {
	return &applicationRepo{
		data: data,
	}
}
