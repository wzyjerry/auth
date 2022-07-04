package oauth2Service

import (
	"context"
	"strings"

	v1 "github.com/wzyjerry/auth/api/oauth2/v1"
	"github.com/wzyjerry/auth/internal/biz/oauth2Biz"
	"github.com/wzyjerry/auth/internal/conf"
	"github.com/wzyjerry/auth/internal/middleware"
	"github.com/wzyjerry/auth/internal/util"
)

type OAuth2Service struct {
	v1.UnimplementedOAuth2ServiceServer

	uc     *oauth2Biz.OAuth2Usecase
	conf   *conf.Security
	helper *util.TokenHelper
}

func NewOAuth2Service(
	uc *oauth2Biz.OAuth2Usecase,
	conf *conf.Security,
	helper *util.TokenHelper,
) *OAuth2Service {
	return &OAuth2Service{
		uc:     uc,
		conf:   conf,
		helper: helper,
	}
}

const (
	responseTypeCode           = "code"
	responseTypeIdToken        = "id_token"
	grantTypeAuthorizationCode = "authorization_code"
	grantTypeClientCredentials = "client_credentials"
	grantTypeRefreshToken      = "refresh_token"
)

func (s *OAuth2Service) PreAuthorize(ctx context.Context, in *v1.PreAuthorizeRequest) (*v1.PreAuthorizeReply, error) {
	token, err := middleware.Validator(ctx, s.helper, middleware.AuthToken)
	if err != nil {
		return nil, err
	}
	var responseType oauth2Biz.ResponseType
	for _, item := range strings.Fields(in.ResponseType) {
		switch item {
		case responseTypeCode:
			responseType |= oauth2Biz.ResponseTypeCode
		case responseTypeIdToken:
			responseType |= oauth2Biz.ResponseTypeIdToken
		default:
			return nil, oauth2Biz.ErrInvalidResponseType
		}
	}
	return s.uc.PreAuthorize(ctx, responseType, token.Subject(), in.ClientId, in.RedirectUri, in.Scope)
}

func (s *OAuth2Service) Authorize(ctx context.Context, in *v1.AuthorizeRequest) (*v1.AuthorizeReply, error) {
	token, err := middleware.Validator(ctx, s.helper, middleware.AuthToken)
	if err != nil {
		return nil, err
	}
	var responseType oauth2Biz.ResponseType
	for _, item := range strings.Fields(in.ResponseType) {
		switch item {
		case responseTypeCode:
			responseType |= oauth2Biz.ResponseTypeCode
		case responseTypeIdToken:
			responseType |= oauth2Biz.ResponseTypeIdToken
		default:
			return nil, oauth2Biz.ErrInvalidResponseType
		}
	}
	return s.uc.Authorize(ctx, responseType, token.Subject(), in.ClientId, in.RedirectUri, in.Scope, in.Nonce)
}

func (s *OAuth2Service) Token(ctx context.Context, in *v1.TokenRequest) (*v1.TokenReply, error) {
	switch in.GrantType {
	case grantTypeAuthorizationCode:
		if in.Code == nil {
			return nil, oauth2Biz.ErrOAuth2InvalidRequest
		}
		return s.uc.AuthorizationCode(ctx, in.ClientId, in.ClientSecret, *in.Code)
	case grantTypeClientCredentials:
		return s.uc.ClientCredentials(ctx, in.ClientId, in.ClientSecret)
	case grantTypeRefreshToken:
		if in.RefreshToken == nil {
			return nil, oauth2Biz.ErrOAuth2InvalidRequest
		}
		return s.uc.RefreshToken(ctx, in.ClientId, in.ClientSecret, *in.RefreshToken)
	default:
		return nil, oauth2Biz.ErrInvalidGrantType
	}
}
