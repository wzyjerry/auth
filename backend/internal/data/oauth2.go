package data

import (
	"context"
	"time"

	"github.com/wzyjerry/auth/internal/biz/oauth2Biz"
	"github.com/wzyjerry/auth/internal/ent"
	"github.com/wzyjerry/auth/internal/ent/application"
	"github.com/wzyjerry/auth/internal/ent/avatar"
	"github.com/wzyjerry/auth/internal/ent/schema/applicationNested"
	"github.com/wzyjerry/auth/internal/ent/schema/avatarNested"
	"github.com/wzyjerry/auth/internal/ent/user"
)

const (
	OAuth2Code         = "OAuth2:Code:"
	OAuth2RefreshToken = "OAuth2:RefreshToken:"
)

type oauth2Repo struct {
	data *Data
}

func (r *oauth2Repo) AvatarExist(ctx context.Context, id string) (bool, error) {
	return r.data.db.Avatar.Query().Where(avatar.KindEQ(int32(avatarNested.Kind_KIND_USER)), avatar.RelIDEQ(id)).Select(avatar.FieldID).Exist(ctx)
}

func (r *oauth2Repo) GetUserInfo(ctx context.Context, id string) (*ent.User, error) {
	return r.data.db.User.Query().Where(user.IDEQ(id)).Select(user.FieldIP, user.FieldNickname).Only(ctx)
}

func (r *oauth2Repo) UpdateClientSecrets(ctx context.Context, id string, clientSecrets []*applicationNested.ClientSecret) error {
	return r.data.db.Application.UpdateOneID(id).SetClientSecrets(clientSecrets).Exec(ctx)
}

func (r *oauth2Repo) cacheKey(ctx context.Context, key string, id string, clientId string, scope string, nonce *string, expiration time.Duration) error {
	var nonceVal string
	if nonce != nil {
		nonceVal = *nonce
	}
	if err := r.data.redis.HSet(ctx, key, "id", id, "clientId", clientId, "scope", scope, "nonce", nonceVal).Err(); err != nil {
		return err
	}
	return r.data.redis.Expire(ctx, key, expiration).Err()
}

func (r *oauth2Repo) getKeyInfoOnce(ctx context.Context, key string) (map[string]string, error) {
	mp, err := r.data.redis.HGetAll(ctx, key).Result()
	if err != nil {
		return nil, err
	}
	r.data.redis.Unlink(ctx, key)
	return mp, nil
}

func (r *oauth2Repo) CacheCode(ctx context.Context, code string, id string, clientId string, scope string, nonce *string, expiration time.Duration) error {
	key := OAuth2Code + code
	return r.cacheKey(ctx, key, id, clientId, scope, nonce, expiration)
}

func (r *oauth2Repo) GetCodeInfoOnce(ctx context.Context, code string) (map[string]string, error) {
	key := OAuth2Code + code
	return r.getKeyInfoOnce(ctx, key)
}

func (r *oauth2Repo) CacheRefreshToken(ctx context.Context, refreshToken string, id string, clientId string, scope string, nonce *string, expiration time.Duration) error {
	key := OAuth2RefreshToken + refreshToken
	return r.cacheKey(ctx, key, id, clientId, scope, nonce, expiration)
}

func (r *oauth2Repo) GetRefreshTokenInfoOnce(ctx context.Context, refreshToken string) (map[string]string, error) {
	key := OAuth2RefreshToken + refreshToken
	return r.getKeyInfoOnce(ctx, key)
}

func (r *oauth2Repo) GetApplication(ctx context.Context, clientId string, fields ...string) (*ent.Application, error) {
	return r.data.db.Application.Query().Where(application.ClientIDEQ(clientId)).Select(fields...).Only(ctx)
}

func NewOAuth2Repo(
	data *Data,
) oauth2Biz.OAuth2Repo {
	return &oauth2Repo{
		data: data,
	}
}
