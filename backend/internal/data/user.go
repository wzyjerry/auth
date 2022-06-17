package data

import (
	"context"
	"strings"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqljson"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/wzyjerry/auth/internal/biz/userBiz"
	"github.com/wzyjerry/auth/internal/conf"
	"github.com/wzyjerry/auth/internal/ent"
	"github.com/wzyjerry/auth/internal/ent/authenticator"
	"github.com/wzyjerry/auth/internal/ent/avatar"
	"github.com/wzyjerry/auth/internal/ent/schema/authenticatorNested"
	"github.com/wzyjerry/auth/internal/ent/schema/avatarNested"
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
	return r.data.redis.Set(ctx, key, code, r.conf.Expiration.Code.AsDuration()).Err()
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

func (r *userRepo) GetAuthenticator(ctx context.Context, kind authenticatorNested.Kind, unique interface{}) (*ent.Authenticator, error) {
	switch kind {
	case authenticatorNested.Kind_KIND_ACCOUNT:
		return r.data.db.Authenticator.
			Query().
			Where(func(s *sql.Selector) {
				s.Where(
					sql.And(
						sql.EQ(authenticator.FieldKind, int32(kind)),
						sqljson.ValueEQ(authenticator.FieldUnique, strings.ToLower(unique.(string)), sqljson.Path("account")),
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
						sqljson.ValueEQ(authenticator.FieldUnique, strings.ToLower(unique.(string)), sqljson.Path("email")),
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
						sqljson.ValueEQ(authenticator.FieldUnique, unique, sqljson.Path("phone")),
					),
				)
			}).
			Only(ctx)
	case authenticatorNested.Kind_KIND_GITHUB:
		return r.data.db.Authenticator.
			Query().
			Where(func(s *sql.Selector) {
				s.Where(
					sql.And(
						sql.EQ(authenticator.FieldKind, int32(kind)),
						sqljson.ValueEQ(authenticator.FieldUnique, unique, sqljson.Path("github")),
					),
				)
			}).
			Only(ctx)
	case authenticatorNested.Kind_KIND_MICROSOFT:
		return r.data.db.Authenticator.
			Query().
			Where(func(s *sql.Selector) {
				s.Where(
					sql.And(
						sql.EQ(authenticator.FieldKind, int32(kind)),
						sqljson.ValueEQ(authenticator.FieldUnique, unique, sqljson.Path("microsoft")),
					),
				)
			}).
			Only(ctx)
	default:
		return nil, userBiz.ErrUnknownKind
	}
}

func (r *userRepo) CreateUser(ctx context.Context, kind int32, unique *authenticatorNested.Unique, password *string, nickname string, ip string) (string, error) {
	id := r.data.db.GenerateId()
	if err := r.data.db.WithTx(ctx, func(tx *ent.Tx) error {
		if err := tx.User.Create().
			SetID(id).
			SetAncestorID(id).
			SetNickname(nickname).
			SetIP(ip).
			SetNillablePassword(password).
			Exec(ctx); err != nil {
			return err
		}
		if err := tx.Authenticator.Create().
			SetID(r.data.db.GenerateId()).
			SetUserID(id).
			SetKind(kind).
			SetUnique(unique).
			Exec(ctx); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return "", err
	}
	return id, nil
}

func (r *userRepo) CreateAvatar(ctx context.Context, id string, base64 string) {
	r.data.db.Avatar.Create().SetID(r.data.db.GenerateId()).SetKind(int32(avatarNested.Kind_KIND_USER)).SetRelID(id).SetAvatar(base64).Exec(ctx)
}

func (r *userRepo) GetAvatar(ctx context.Context, id string) (*string, error) {
	avatar, err := r.data.db.Avatar.Query().Where(avatar.KindEQ(int32(avatarNested.Kind_KIND_USER)), avatar.RelIDEQ(id)).Only(ctx)
	if err != nil {
		return nil, err
	}
	return avatar.Avatar, nil
}

func NewUserRepo(
	data *Data,
	conf *conf.Security,
	logger log.Logger) userBiz.UserRepo {
	return &userRepo{
		data: data,
		conf: conf,
		log:  log.NewHelper(logger),
	}
}
