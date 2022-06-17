package userBiz

import (
	"context"
	"net/http"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/wzyjerry/auth/internal/ent"
	"github.com/wzyjerry/auth/internal/ent/schema/authenticatorNested"
)

var (
	ErrNotFound              = errors.New(http.StatusNotFound, "NOT_FOUND", "not found")
	ErrUnknownKind           = errors.New(http.StatusInternalServerError, "UNKNOWN_KIND", "unknown kind")
	ErrInvalidPassword       = errors.New(http.StatusBadRequest, "INVALID_PASSWORD", "invalid password")
	ErrUniqueRequired        = errors.New(http.StatusBadRequest, "UNIQUE_REQUIRED", "unique required")
	ErrCodeMismatch          = errors.New(http.StatusBadRequest, "CODE_MISMATCH", "code mismatch")
	ErrAuthenticatorConflict = errors.New(http.StatusConflict, "AUTHENTICATOR_CONFLICT", "authenticator conflict")
	ErrPasswordLogin         = errors.Newf(http.StatusBadRequest, "PASSWORD_LOGIN", "invalid unique token or password")
)

type UserRepo interface {
	// redis部分
	CacheRegisterEmail(ctx context.Context, email string, code string) error
	CacheRegisterPhone(ctx context.Context, phone string, code string) error
	CacheLoginEmail(ctx context.Context, email string, code string) error
	CacheLoginPhone(ctx context.Context, phone string, code string) error
	VerifyRegisterEmailCode(ctx context.Context, email string, code string) (bool, error)
	VerifyRegisterPhoneCode(ctx context.Context, phone string, code string) (bool, error)
	VerifyLoginEmailCode(ctx context.Context, email string, code string) (bool, error)
	VerifyLoginPhoneCode(ctx context.Context, phone string, code string) (bool, error)
	// db部分
	GetAuthenticator(ctx context.Context, kind authenticatorNested.Kind, unique interface{}) (*ent.Authenticator, error)
	CreateUser(ctx context.Context, kind int32, unique *authenticatorNested.Unique, password *string, nickname string, ip string) (string, error)
	CreateAvatar(ctx context.Context, id string, base64 string)
	GetUser(ctx context.Context, id string) (*ent.User, error)
	GetAvatar(ctx context.Context, id string) (*string, error)
}
