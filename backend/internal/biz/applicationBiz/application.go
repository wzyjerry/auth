package applicationBiz

import (
	"context"
	"net/http"

	"github.com/go-kratos/kratos/v2/errors"
	v1 "github.com/wzyjerry/auth/api/application/v1"
	"github.com/wzyjerry/auth/internal/ent"
	"github.com/wzyjerry/auth/internal/ent/schema/applicationNested"
)

type ApplicationRepo interface {
	// db部分
	Create(ctx context.Context, admin string, name string, homepage string, description *string, callback string) (string, error)
	Retrieve(ctx context.Context, admin string, id string) (*ent.Application, error)
	GetAvatar(ctx context.Context, id string) (string, error)
	GenerateClientSecret(ctx context.Context, admin string, id string, description string) (*applicationNested.ClientSecret, error)
	RevokeClientSecret(ctx context.Context, admin string, id string, secretId string) error
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
	avatar, err := uc.repo.GetAvatar(ctx, id)
	if err == nil {
		reply.Avatar = &avatar
	} else if !ent.IsNotFound(err) {
		return nil, err
	}
	for _, item := range application.ClientSecrets {
		secret := &v1.Secret{
			Id:           *item.Id,
			LastUsed:     item.LastUsed,
			Description:  *item.Description,
			MaskedSecret: "*****" + (*item.Secret)[len(*item.Secret)-5:],
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
			Id:           *secret.Id,
			LastUsed:     secret.LastUsed,
			Description:  *secret.Description,
			MaskedSecret: *secret.Secret,
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
