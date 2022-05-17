package data

import (
	"context"
	"net/http"
	"strings"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqljson"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/wzyjerry/auth/internal/biz"
	"github.com/wzyjerry/auth/internal/conf"
	"github.com/wzyjerry/auth/internal/ent"
	"github.com/wzyjerry/auth/internal/ent/authenticator"
	"github.com/wzyjerry/auth/internal/ent/schema/authenticatorNested"
)

type userRepo struct {
	data *Data
	conf *conf.Security
	log  *log.Helper
}

const (
	AuthRegisterEmail = "Auth:Register:Email:"
	AuthRegisterPhone = "Auth:Register:Phone:"
	AuthLoginEmail    = "Auth:Login:Email:"
	AuthLoginPhone    = "Auth:Login:Phone:"
)

func (r *userRepo) cacheCode(ctx context.Context, key string, code string) error {
	if err := r.data.redis.Set(ctx, key, code, r.conf.Expiration.Code.AsDuration()).Err(); err != nil {
		return errors.New(http.StatusInternalServerError, "ERR_CACHE_CODE", err.Error())
	}
	return nil
}

func (r *userRepo) GetUser(ctx context.Context, id string) (*ent.User, error) {
	return r.data.db.User.Get(ctx, id)
}

func (r *userRepo) CacheRegisterEmail(ctx context.Context, email string, code string) error {
	return r.cacheCode(ctx, AuthRegisterEmail+email, code)
}

func (r *userRepo) CacheRegisterPhone(ctx context.Context, phone string, code string) error {
	return r.cacheCode(ctx, AuthRegisterPhone+phone, code)
}

func (r *userRepo) CacheLoginEmail(ctx context.Context, email string, code string) error {
	return r.cacheCode(ctx, AuthLoginEmail+email, code)
}

func (r *userRepo) CacheLoginPhone(ctx context.Context, phone string, code string) error {
	return r.cacheCode(ctx, AuthLoginPhone+phone, code)
}

func (r *userRepo) verifyCode(ctx context.Context, key string, code string) (bool, error) {
	result := r.data.redis.Get(ctx, key)
	if result.Err() != nil {
		return false, result.Err()
	}
	if result.Val() != code {
		return false, nil
	}
	r.data.redis.Unlink(ctx, key)
	return true, nil
}

func (r *userRepo) VerifyRegisterEmailCode(ctx context.Context, email string, code string) (bool, error) {
	return r.verifyCode(ctx, AuthRegisterEmail+email, code)
}

func (r *userRepo) VerifyRegisterPhoneCode(ctx context.Context, phone string, code string) (bool, error) {
	return r.verifyCode(ctx, AuthRegisterPhone+phone, code)
}

func (r *userRepo) VerifyLoginEmailCode(ctx context.Context, email string, code string) (bool, error) {
	return r.verifyCode(ctx, AuthLoginEmail+email, code)
}

func (r *userRepo) VerifyLoginPhoneCode(ctx context.Context, phone string, code string) (bool, error) {
	return r.verifyCode(ctx, AuthLoginPhone+phone, code)
}

func (r *userRepo) GetAuthenticator(ctx context.Context, kind authenticatorNested.Kind, unique string) (*ent.Authenticator, error) {
	switch kind {
	case authenticatorNested.Kind_KIND_ACCOUNT:
		return r.data.db.Authenticator.
			Query().
			Where(func(s *sql.Selector) {
				s.Where(
					sql.And(
						sql.EQ(authenticator.FieldKind, int32(kind)),
						sqljson.ValueEQ(authenticator.FieldAnchor, strings.ToLower(unique), sqljson.Path("account", "username")),
					),
				)
			}).
			Only(ctx)
	case authenticatorNested.Kind_KIND_EMAIL:
		return r.data.db.Authenticator.
			Query().
			Where(func(s *sql.Selector) {
				s.Where(
					sql.And(
						sql.EQ(authenticator.FieldKind, int32(kind)),
						sqljson.ValueEQ(authenticator.FieldAnchor, strings.ToLower(unique), sqljson.Path("email", "email")),
					),
				)
			}).
			Only(ctx)
	case authenticatorNested.Kind_KIND_PHONE:
		return r.data.db.Authenticator.
			Query().
			Where(func(s *sql.Selector) {
				s.Where(
					sql.And(
						sql.EQ(authenticator.FieldKind, int32(kind)),
						sqljson.ValueEQ(authenticator.FieldAnchor, unique, sqljson.Path("phone", "phone")),
					),
				)
			}).
			Only(ctx)
	default:
		return nil, biz.ErrUnknownKind
	}
}

func (r *userRepo) CreateUser(ctx context.Context, kind int32, anchor *authenticatorNested.Anchor, password *string, nickname string, ip string, avatar *string) (string, error) {
	id := r.data.db.GenerateId()
	if err := r.data.db.WithTx(ctx, func(tx *ent.Tx) error {
		if err := tx.User.Create().
			SetID(id).
			SetAncestorID(id).
			SetNickname(nickname).
			SetIP(ip).
			SetNillableAvatar(avatar).
			SetNillablePassword(password).
			Exec(ctx); err != nil {
			return err
		}
		if err := tx.Authenticator.Create().
			SetID(r.data.db.GenerateId()).
			SetUserID(id).
			SetKind(kind).
			SetAnchor(anchor).
			Exec(ctx); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return "", err
	}
	return id, nil
}

func NewUserRepo(
	data *Data,
	conf *conf.Security,
	logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		conf: conf,
		log:  log.NewHelper(logger),
	}
}
