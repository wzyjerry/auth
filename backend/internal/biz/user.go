package biz

import (
	"context"
	"net/http"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/wzyjerry/auth/internal/ent"
	"github.com/wzyjerry/auth/internal/ent/schema/authenticatorNested"
)

var (
	ErrUnknownKind           = errors.New(http.StatusInternalServerError, "UNKNOWN_KIND", "unknown kind")
	ErrInvalidPassword       = errors.New(http.StatusBadRequest, "INVALID_PASSWORD", "invalid password")
	ErrCodeMismatch          = errors.New(http.StatusBadRequest, "CODE_MISMATCH", "code mismatch")
	ErrAuthenticatorConflict = errors.New(http.StatusConflict, "AUTHENTICATOR_CONFLICT", "authenticator conflict")
	ErrPasswordLogin         = errors.Newf(http.StatusBadRequest, "PASSWORD_LOGIN", "invalid unique token or password")
	ErrNetworkError          = errors.Newf(http.StatusInternalServerError, "NETWORD_ERROR", "http functional error")
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
	CreateUser(ctx context.Context, kind int32, unique *authenticatorNested.Unique, password *string, nickname string, ip string, avatar *string) (string, error)
	GetUser(ctx context.Context, id string) (*ent.User, error)
}
