package applicationBiz

import (
	"context"
	"net/http"

	"github.com/go-kratos/kratos/v2/errors"
	v1 "github.com/wzyjerry/auth/api/application/v1"
	"github.com/wzyjerry/auth/internal/ent"
	"github.com/wzyjerry/auth/internal/ent/schema/applicationNested"
	"github.com/wzyjerry/auth/internal/util"
)

type ApplicationRepo interface {
	// db部分
	Create(ctx context.Context, admin string, name string, homepage string, description *string, callback string) (string, error)
	Retrieve(ctx context.Context, admin string, id string) (*ent.Application, error)
	GetLogo(ctx context.Context, id string) (*string, error)
	GenerateClientSecret(ctx context.Context, admin string, id string, description string) (*applicationNested.ClientSecret, error)
	RevokeClientSecret(ctx context.Context, admin string, id string, secretId string) error
	SetLogo(ctx context.Context, admin string, id string, logo string) error
	Update(ctx context.Context, admin string, id string, iname string, homepage string, description *string, callback string) error
	Delete(ctx context.Context, admin string, id string) error
	GetAll(ctx context.Context, admin string) ([]*ent.Application, error)
	GetLogoMap(ctx context.Context, ids []string) (map[string]string, error)
}

type ApplicationUsecase struct {
	repo ApplicationRepo
}

var (
	ErrApplicationNotFound = errors.New(http.StatusNotFound, "APPLICATION_NOT_FOUND", "application not found")
	ErrRevokeBadRequest    = errors.New(http.StatusBadRequest, "REVOKE_BAD_REQUEST", "can not revoke the last client secret")
)

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

func (uc *ApplicationUsecase) Retrieve(ctx context.Context, admin string, id string) (*v1.RetrieveReply, error) {
	application, err := uc.repo.Retrieve(ctx, admin, id)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, ErrApplicationNotFound
		}
		return nil, err
	}
	reply := &v1.RetrieveReply{
		Name:          *application.Name,
		ClientId:      *application.ClientID,
		ClientSecrets: make([]*v1.Secret, 0, len(application.ClientSecrets)),
		Homepage:      *application.Homepage,
		Description:   application.Description,
		Callback:      *application.Callback,
	}
	logo, err := uc.repo.GetLogo(ctx, id)
	if err != nil {
		return nil, err

	}
	reply.Logo = logo
	for _, item := range application.ClientSecrets {
		secret := &v1.Secret{
			Id:          *item.Id,
			LastUsed:    item.LastUsed,
			Description: *item.Description,
			Masked:      true,
			Secret:      "*****" + (*item.Secret)[len(*item.Secret)-5:],
		}
		reply.ClientSecrets = append(reply.ClientSecrets, secret)
	}
	return reply, nil
}

func (uc *ApplicationUsecase) GenerateClientSecret(ctx context.Context, admin string, id string, description string) (*v1.GenerateClientSecretReply, error) {
	secret, err := uc.repo.GenerateClientSecret(ctx, admin, id, description)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, ErrApplicationNotFound
		}
		return nil, err
	}
	return &v1.GenerateClientSecretReply{
		Secret: &v1.Secret{
			Id:          *secret.Id,
			LastUsed:    secret.LastUsed,
			Description: *secret.Description,
			Masked:      false,
			Secret:      *secret.Secret,
		},
	}, nil

}

func (uc *ApplicationUsecase) RevokeClientSecret(ctx context.Context, admin string, id string, secretId string) error {
	err := uc.repo.RevokeClientSecret(ctx, admin, id, secretId)
	if ent.IsNotFound(err) {
		return ErrApplicationNotFound
	}
	return err
}

func (uc *ApplicationUsecase) UploadLogo(ctx context.Context, admin string, id string, logo string) error {
	err := uc.repo.SetLogo(ctx, admin, id, logo)
	if ent.IsNotFound(err) {
		return ErrApplicationNotFound
	}
	return err
}

func (uc *ApplicationUsecase) Update(ctx context.Context, admin string, id string, name string, homepage string, description *string, callback string) error {
	return uc.repo.Update(ctx, admin, id, name, homepage, description, callback)
}

func (uc *ApplicationUsecase) Delete(ctx context.Context, admin string, id string) error {
	return uc.repo.Delete(ctx, admin, id)
}

func (uc *ApplicationUsecase) GetAll(ctx context.Context, admin string) (*v1.GetAllReply, error) {
	applications, err := uc.repo.GetAll(ctx, admin)
	if err != nil {
		return nil, err
	}
	reply := &v1.GetAllReply{
		ApplicationOverviews: make([]*v1.ApplicationOverview, 0, len(applications)),
	}
	ids := make([]string, 0, len(applications))
	for _, item := range applications {
		applicationOverview := &v1.ApplicationOverview{
			Id:   item.ID,
			Name: *item.Name,
		}
		if item.Description != nil {
			if len(*item.Description) > 37 {
				applicationOverview.MaskedDescription = util.P((*item.Description)[:37] + "...")
			} else {
				applicationOverview.MaskedDescription = item.Description
			}
		}
		ids = append(ids, item.ID)
		reply.ApplicationOverviews = append(reply.ApplicationOverviews, applicationOverview)
	}
	logoMap, err := uc.repo.GetLogoMap(ctx, ids)
	if err != nil {
		return nil, err
	}
	for _, item := range reply.ApplicationOverviews {
		if logo, ok := logoMap[item.Id]; ok {
			item.Logo = &logo
		}
	}
	return reply, nil
}
