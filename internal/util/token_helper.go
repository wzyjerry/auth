package util

import (
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"os"
	"strings"
	"time"

	"github.com/lestrrat-go/jwx/v2/jwa"
	"github.com/lestrrat-go/jwx/v2/jwt"
	"github.com/wzyjerry/auth/internal/conf"
)

type TokenHelper struct {
	conf       *conf.Security
	privateKey *rsa.PrivateKey
}

const (
	issuerWord         = "oauth.windranger.tk"
	nonceKey           = "nonce"
	codeHashKey        = "c_hash"
	accessTokenHashKey = "at_hash"
	ipAddressKey       = "ipaddr"
	nicknameKey        = "nickname"
	avatarKey          = "avatar"
	idTypeKey          = "idtyp"
	appWord            = "app"
)

func NewTokenHelper(
	conf *conf.Security,
	privateKey *rsa.PrivateKey,
) *TokenHelper {
	return &TokenHelper{
		conf:       conf,
		privateKey: privateKey,
	}
}

func (h *TokenHelper) randAccessTokenExpiration() time.Duration {
	min := h.conf.Expiration.AccessTokenMin.AsDuration().Seconds()
	max := h.conf.Expiration.AccessTokenMax.AsDuration().Seconds()
	diff := max - min
	return time.Duration(min+(diff*RndFloat64())) * time.Second
}

func (h *TokenHelper) ParseJWT(
	signed string,
) (*jwt.Token, error) {
	token, err := jwt.ParseString(signed, jwt.WithKey(jwa.RS256, h.privateKey.PublicKey))
	if err != nil {
		return nil, err
	}
	return &token, nil
}

func (h *TokenHelper) IsIdToken(token *jwt.Token) bool {
	_, ok := (*token).Get(nonceKey)
	return ok
}

func (h *TokenHelper) IsAuthToken(token *jwt.Token) bool {
	audience := (*token).Audience()
	return len(audience) > 0 && strings.Compare(audience[0], os.Getenv(h.conf.Oauth.ClientId)) == 0
}

func (h *TokenHelper) IsClientToken(token *jwt.Token) bool {
	idType, ok := (*token).Get(idTypeKey)
	return ok && strings.Compare(idType.(string), appWord) == 0
}

func (h *TokenHelper) IsUserToken(token *jwt.Token) bool {
	return !h.IsClientToken(token) && !h.IsAuthToken(token)
}

func (h *TokenHelper) GenerateAccessToken(
	clientId string,
	subject string,
) jwt.Token {
	now := time.Now()
	accessTokenExpiration := h.randAccessTokenExpiration()
	token, _ := jwt.NewBuilder().
		Audience([]string{clientId}).
		Issuer(issuerWord).
		IssuedAt(now).
		NotBefore(now).
		Expiration(now.Add(accessTokenExpiration)).
		JwtID(NewUUID()).
		Subject(subject).
		Build()
	if clientId == subject {
		token.Set(idTypeKey, appWord)
	}
	return token
}

func (h *TokenHelper) GenerateIdToken(
	accessToken jwt.Token,
	nonce string,
	codeHash *string,
	accessTokenHash *string,
	ip string,
	nickname string,
	avatar *string,
) jwt.Token {
	token := accessToken
	token.Set(nonceKey, nonce)
	if codeHash != nil {
		token.Set(codeHashKey, *codeHash)
	}
	if accessTokenHash != nil {
		token.Set(accessTokenHashKey, *accessTokenHash)
	}
	token.Set(ipAddressKey, ip)
	token.Set(nicknameKey, nickname)
	if avatar != nil {
		token.Set(avatarKey, avatar)
	}
	return token
}

func (h *TokenHelper) SignJWT(token jwt.Token) (string, error) {
	signed, err := jwt.Sign(token, jwt.WithKey(jwa.RS256, h.privateKey))
	if err != nil {
		return "", err
	}
	return string(signed), nil
}

func (h *TokenHelper) generateHash(b []byte) string {
	sum := sha256.Sum256(b)
	return base64.StdEncoding.EncodeToString(sum[:len(sum)/2])
}

func (h *TokenHelper) GenerateCodeHash(code string) string {
	return h.generateHash([]byte(code))
}

func (h *TokenHelper) GenerateAccessTokenHash(accessToken string) string {
	return h.generateHash([]byte(accessToken))
}