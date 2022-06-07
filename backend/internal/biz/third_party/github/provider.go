package github

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"os"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/wzyjerry/auth/internal/biz"
	"github.com/wzyjerry/auth/internal/conf"
)

type Info struct {
	Id     int64  `json:"id"`
	Name   string `json:"name"`
	Avatar string `json:"avatar_url"`
}

func (i *Info) GetUnique() int64 {
	return i.Id
}

func (i *Info) GetNickname() string {
	return i.Name
}

func (i *Info) GetAvatar() *string {
	return &i.Name
}

type Provider struct {
	conf *conf.Security
}

var (
	ErrCodeMismatch = errors.New(http.StatusBadRequest, "CODE_MISMATCH", "code mismatch")
	ErrNetworkError = errors.Newf(http.StatusInternalServerError, "NETWORD_ERROR", "http functional error")
)

func New(
	conf *conf.Security,
) biz.GithubProvider {
	return &Provider{
		conf: conf,
	}
}

func (p *Provider) acquirAccessToken(ctx context.Context, code string) (string, error) {
	body := bytes.NewBuffer(nil)
	json.NewEncoder(body).Encode(map[string]any{
		"code":          code,
		"client_id":     os.Getenv(p.conf.Oauth.Github.ClientId),
		"client_secret": os.Getenv(p.conf.Oauth.Github.ClientSecret),
	})
	req, err := http.NewRequest(http.MethodPost, "https://github.com/login/oauth/access_token", body)
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	resp, err := http.DefaultClient.Do(req.WithContext(ctx))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		reply := make(map[string]interface{})
		if err := json.NewDecoder(resp.Body).Decode(&reply); err != nil {
			return "", err
		}
		token, ok := reply["access_token"].(string)
		if !ok {
			return "", ErrCodeMismatch
		}
		return token, nil
	} else {
		return "", ErrNetworkError
	}
}

func (p *Provider) acquirUser(ctx context.Context, accessToken string) (*Info, error) {
	req, err := http.NewRequest(http.MethodGet, "https://api.github.com/user", nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "token "+accessToken)
	resp, err := http.DefaultClient.Do(req.WithContext(ctx))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		var reply Info
		if err := json.NewDecoder(resp.Body).Decode(&reply); err != nil {
			return nil, err
		}
		return &reply, nil
	} else {
		return nil, err
	}
}

func (p *Provider) Login(ctx context.Context, code string) (biz.BasicInfo[int64], error) {
	accessToken, err := p.acquirAccessToken(ctx, code)
	if err != nil {
		return nil, err
	}
	return p.acquirUser(ctx, accessToken)
}
