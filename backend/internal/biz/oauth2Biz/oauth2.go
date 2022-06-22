package oauth2Biz

import (
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/go-kratos/kratos/v2/errors"
	v1 "github.com/wzyjerry/auth/api/oauth2/v1"
	"github.com/wzyjerry/auth/internal/conf"
	"github.com/wzyjerry/auth/internal/ent"
	"github.com/wzyjerry/auth/internal/ent/application"
	"github.com/wzyjerry/auth/internal/ent/schema/applicationNested"
	"github.com/wzyjerry/auth/internal/util"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type OAuth2Repo interface {
	// redis部分
	CacheCode(ctx context.Context, code string, id string, clientId string, scope string, nonce *string, expiration time.Duration) error
	CacheRefreshToken(ctx context.Context, refreshToken string, id string, clientId string, scope string, nonce *string, expiration time.Duration) error
	GetCodeInfoOnce(ctx context.Context, code string) (map[string]string, error)
	GetRefreshTokenInfoOnce(ctx context.Context, refreshToken string) (map[string]string, error)
	// db部分
	GetUserInfo(ctx context.Context, id string) (*ent.User, error)
	AvatarExist(ctx context.Context, id string) (bool, error)
	GetApplication(ctx context.Context, clientId string, fields ...string) (*ent.Application, error)
	UpdateClientSecrets(ctx context.Context, id string, clientSecrets []*applicationNested.ClientSecret) error
}

type OAuth2Usecase struct {
	repo   OAuth2Repo
	conf   *conf.Security
	helper *util.TokenHelper
}

type ResponseType int

const (
	ResponseTypeCode ResponseType = 1 << iota
	ResponseTypeIdToken
)

const (
	bearerWord = "Bearer"
)

type Scope string

const (
	scopeUser          Scope = "user"
	scopeOfflineAccess Scope = "offline_access"
	scopeOpenid        Scope = "openid"
)

var (
	ErrRedirectUriNotMatch  = errors.New(http.StatusBadRequest, "REDIRECT_URI_NOT_MATCH", "redirect uri not match")
	ErrInvalidScope         = errors.New(http.StatusBadRequest, "OAUTH2_INVALID_REQUEST", "invalid scope")
	ErrInvalidRedirectURI   = errors.New(http.StatusBadRequest, "OAUTH2_INVALID_REQUEST", "invalid redirect uri")
	ErrOAuth2InvalidRequest = errors.New(http.StatusBadRequest, "OAUTH2_INVALID_REQUEST", "invalid request")
	ErrOAuth2InvalidGrant   = errors.New(http.StatusBadRequest, "OAUTH2_INVALID_GRANT", "invalid grant")
	ErrInvalidGrantType     = errors.New(http.StatusBadRequest, "INVALID_GRANT_TYPE", "unknown grant type")
	ErrInvalidResponseType  = errors.New(http.StatusBadRequest, "INVALID_RESPONSE_TYPE", "unsupported response type")
)

func (uc *OAuth2Usecase) generateToken(ctx context.Context, clientId string, clientSecret string, info map[string]string) (*v1.TokenReply, error) {
	item, err := uc.repo.GetApplication(ctx, clientId, application.FieldID, application.FieldClientSecrets)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, ErrOAuth2InvalidRequest
		}
		return nil, err
	}
	verified := false
	for _, secret := range item.ClientSecrets {
		if *secret.Secret == clientSecret {
			secret.LastUsed = timestamppb.Now()
			go uc.repo.UpdateClientSecrets(ctx, item.ID, item.ClientSecrets)
			verified = true
			break
		}
	}
	if !verified {
		return nil, ErrOAuth2InvalidRequest
	}
	infoId, infoClientId, infoScope, infoNonce := info["id"], info["clientId"], info["scope"], info["nonce"]
	if infoClientId != clientId {
		return nil, ErrOAuth2InvalidGrant
	}
	token := uc.helper.GenerateBasicToken(clientId, infoId)
	scopes := strings.Fields(infoScope)
	token = uc.helper.GenerateAccessToken(token, scopes)
	accessToken, err := uc.helper.SignJWT(token)
	if err != nil {
		return nil, ErrOAuth2InvalidGrant
	}
	reply := &v1.TokenReply{
		TokenType:   bearerWord,
		AccessToken: accessToken,
		ExpiresIn:   uc.conf.Expiration.AccessToken.GetSeconds(),
		Scope:       infoScope,
	}

	for _, scope := range scopes {
		// 作用域要求颁发刷新令牌
		if Scope(scope) == scopeOfflineAccess {
			refreshToken := util.NewUUID()
			if err := uc.repo.CacheRefreshToken(ctx, refreshToken, infoId, infoClientId, infoScope, &infoNonce, uc.conf.Expiration.RefreshToken.AsDuration()+uc.conf.Expiration.TokenExtend.AsDuration()); err != nil {
				return nil, ErrOAuth2InvalidGrant
			}
			reply.RefreshToken = &refreshToken
			break
		}
	}

	if infoNonce != "" {
		// OIDC请求nonce一定非空
		accessTokenHash := uc.helper.GenerateAccessTokenHash(accessToken)
		idToken, err := uc.idToken(ctx, infoId, infoClientId, infoNonce, nil, &accessTokenHash)
		if err != nil {
			return nil, ErrOAuth2InvalidGrant
		}
		reply.IdToken = &idToken
	}

	return reply, nil
}

func (uc *OAuth2Usecase) RefreshToken(ctx context.Context, clientId string, clientSecret string, refreshToken string) (*v1.TokenReply, error) {
	info, err := uc.repo.GetRefreshTokenInfoOnce(ctx, refreshToken)
	if err != nil {
		return nil, ErrOAuth2InvalidGrant
	}
	return uc.generateToken(ctx, clientId, clientSecret, info)
}

func (uc *OAuth2Usecase) AuthorizationCode(ctx context.Context, clientId string, clientSecret string, code string) (*v1.TokenReply, error) {
	info, err := uc.repo.GetCodeInfoOnce(ctx, code)
	if err != nil {
		return nil, ErrOAuth2InvalidGrant
	}
	return uc.generateToken(ctx, clientId, clientSecret, info)
}

func (uc *OAuth2Usecase) ClientCredentials(ctx context.Context, clientId string, clientSecret string) (*v1.TokenReply, error) {
	info := map[string]string{
		"id":       clientId,
		"clientId": clientId,
		"scope":    "",
	}
	return uc.generateToken(ctx, clientId, clientSecret, info)
}

func (uc *OAuth2Usecase) code(ctx context.Context, id string, clientId string, redirectUri string, scope string, nonce *string) (string, error) {
	code := util.NewUUID()
	if err := uc.repo.CacheCode(ctx, code, id, clientId, scope, nonce, uc.conf.Expiration.Code.AsDuration()); err != nil {
		return "", err
	}
	return code, nil
}

func (uc *OAuth2Usecase) idToken(ctx context.Context, id string, clientId string, nonce string, codeHash *string, accessTokenHash *string) (string, error) {
	token := uc.helper.GenerateBasicToken(clientId, id)
	user, err := uc.repo.GetUserInfo(ctx, id)
	if err != nil {
		return "", err
	}
	hasAvatar, err := uc.repo.AvatarExist(ctx, id)
	if err != nil {
		return "", err
	}
	token = uc.helper.GenerateIdToken(token, nonce, codeHash, accessTokenHash, *user.IP, *user.Nickname, hasAvatar)
	return uc.helper.SignJWT(token)
}

func (uc *OAuth2Usecase) Authorize(ctx context.Context, responseType ResponseType, id string, clientId string, redirectUri string, scope string, nonce *string) (*v1.AuthorizeReply, error) {
	item, err := uc.repo.GetApplication(ctx, clientId, application.FieldCallback)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, ErrInvalidRedirectURI
		}
		return nil, err
	}
	if ok := util.VerifyRedirectUri(*item.Callback, redirectUri); !ok {
		return nil, ErrInvalidRedirectURI
	}
	scopeInvalid, hasOpenid := false, false
	for _, item := range strings.Fields(scope) {
		switch Scope(item) {
		case scopeUser, scopeOfflineAccess:
		// do nothing
		case scopeOpenid:
			hasOpenid = true
		default:
			scopeInvalid = true
		}
	}
	if scopeInvalid {
		return nil, ErrInvalidScope
	}
	if responseType == ResponseTypeCode {
		code, err := uc.code(ctx, id, clientId, redirectUri, scope, nil)
		if err != nil {
			return nil, err
		}
		return &v1.AuthorizeReply{
			Code: code,
		}, nil
	}
	if nonce == nil || *nonce == "" || !hasOpenid {
		return nil, ErrOAuth2InvalidGrant
	}
	if responseType == ResponseTypeIdToken {
		idToken, err := uc.idToken(ctx, id, clientId, *nonce, nil, nil)
		if err != nil {
			return nil, err
		}
		return &v1.AuthorizeReply{
			IdToken: &idToken,
		}, nil
	}
	if responseType == ResponseTypeCode|ResponseTypeIdToken {
		reply := new(v1.AuthorizeReply)
		code, err := uc.code(ctx, id, clientId, redirectUri, scope, nonce)
		if err != nil {
			return nil, err
		}
		reply.Code = code
		codeHash := uc.helper.GenerateCodeHash(code)
		idToken, err := uc.idToken(ctx, id, clientId, *nonce, &codeHash, nil)
		if err != nil {
			return nil, err
		}
		reply.IdToken = &idToken
		return reply, nil
	}
	return nil, ErrInvalidResponseType
}

func NewOAuth2Usecase(
	repo OAuth2Repo,
	conf *conf.Security,
	helper *util.TokenHelper,
) *OAuth2Usecase {
	return &OAuth2Usecase{
		repo:   repo,
		conf:   conf,
		helper: helper,
	}
}
