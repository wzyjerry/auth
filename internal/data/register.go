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

type registerRepo struct {
	data *Data
	conf *conf.Security
	log  *log.Helper
}

const (
	AuthPreEmail = "Auth:Email:Pre:"
	AuthPrePhone = "Auth:Phone:Pre:"
)

func (r *registerRepo) cacheCode(ctx context.Context, key string, code string) error {
	if err := r.data.redis.Set(ctx, key, code, r.conf.Expiration.Code.AsDuration()).Err(); err != nil {
		return errors.New(http.StatusInternalServerError, "ERR_CACHE_CODE", err.Error())
	}
	return nil
}

func (r *registerRepo) CachePreEmail(ctx context.Context, email string, code string) error {
	return r.cacheCode(ctx, AuthPreEmail+email, code)
}

func (r *registerRepo) CachePrePhone(ctx context.Context, phone string, code string) error {
	return r.cacheCode(ctx, AuthPrePhone+phone, code)
}

func (r *registerRepo) verifyCode(ctx context.Context, key string, code string) (bool, error) {
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

func (r *registerRepo) VerifyPreEmailCode(ctx context.Context, email string, code string) (bool, error) {
	return r.verifyCode(ctx, AuthPreEmail+email, code)
}

func (r *registerRepo) VerifyPrePhoneCode(ctx context.Context, phone string, code string) (bool, error) {
	return r.verifyCode(ctx, AuthPrePhone+phone, code)
}

func (r *registerRepo) GetAuthenticator(ctx context.Context, kind authenticatorNested.Kind, unique string) (*ent.Authenticator, error) {
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

func (r *registerRepo) CreateUser(ctx context.Context, kind int32, anchor *authenticatorNested.Anchor, password *string, nickname string, ip string, avatar *string) (string, error) {
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

func NewRegisterRepo(
	data *Data,
	conf *conf.Security,
	logger log.Logger) biz.RegisterRepo {
	return &registerRepo{
		data: data,
		conf: conf,
		log:  log.NewHelper(logger),
	}
}
