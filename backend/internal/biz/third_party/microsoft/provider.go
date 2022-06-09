package microsoft

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/wzyjerry/auth/internal/biz"
	"github.com/wzyjerry/auth/internal/conf"
	"github.com/wzyjerry/auth/internal/util"
)

type Info struct {
	Id          string `json:"id"`
	DisplayName string `json:"displayName"`
	accessToken string
}

func (i *Info) GetUnique() string {
	return i.Id
}

func (i *Info) GetNickname() string {
	return i.DisplayName
}

func (i *Info) GetAvatar() func(ctx context.Context) (string, error) {
	return func(ctx context.Context) (string, error) {
		req, err := http.NewRequest(http.MethodGet, `https://graph.microsoft.com/v1.0/me/photo/$value`, nil)
		req.Header.Set("Authorization", "Bearer "+i.accessToken)
		if err != nil {
			return "", err
		}
		resp, err := http.DefaultClient.Do(req.WithContext(ctx))
		if err != nil {
			return "", err
		}
		defer resp.Body.Close()
		if resp.StatusCode != http.StatusOK {
			return "", ErrNetworkError
		}
		img, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return "", err
		}
		return util.ImgBase64(img), nil
	}
}

type Provider struct {
	conf *conf.ThirdParty
}

var (
	ErrCodeMismatch = errors.New(http.StatusBadRequest, "CODE_MISMATCH", "code mismatch")
	ErrNetworkError = errors.Newf(http.StatusInternalServerError, "NETWORK_ERROR", "http functional error")
)

func New(
	conf *conf.ThirdParty,
) biz.MicrosoftProvider {
	return &Provider{
		conf: conf,
	}
}

func (p *Provider) acquirAccessToken(ctx context.Context, code string) (string, error) {
	body := url.Values{}
	body.Set("grant_type", "authorization_code")
	body.Set("code", code)
	body.Set("client_id", os.Getenv(p.conf.Microsoft.ClientId))
	body.Set("client_secret", os.Getenv(p.conf.Microsoft.ClientSecret))
	body.Set("redirect_uri", p.conf.Microsoft.RedirectUri)
	req, err := http.NewRequest(http.MethodPost, "https://login.microsoftonline.com/common/oauth2/v2.0/token", strings.NewReader(body.Encode()))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Accept", "application/json")
	resp, err := http.DefaultClient.Do(req.WithContext(ctx))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return "", ErrNetworkError
	}
	reply := make(map[string]interface{})
	if err := json.NewDecoder(resp.Body).Decode(&reply); err != nil {
		return "", err
	}
	token, ok := reply["access_token"].(string)
	if !ok {
		return "", ErrCodeMismatch
	}
	return token, nil
}

func (p *Provider) acquirUser(ctx context.Context, accessToken string) (*Info, error) {
	req, err := http.NewRequest(http.MethodGet, "https://graph.microsoft.com/v1.0/me", nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+accessToken)
	resp, err := http.DefaultClient.Do(req.WithContext(ctx))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, err
	}
	var reply Info
	if err := json.NewDecoder(resp.Body).Decode(&reply); err != nil {
		return nil, err
	}
	reply.accessToken = accessToken
	return &reply, nil
}

func (p *Provider) Login(ctx context.Context, code string) (biz.BasicInfo[string], error) {
	accessToken, err := p.acquirAccessToken(ctx, code)
	if err != nil {
		return nil, err
	}
	return p.acquirUser(ctx, accessToken)
}
