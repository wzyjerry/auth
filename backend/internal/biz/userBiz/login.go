package userBiz

import (
	"context"
	"strings"

	v1 "github.com/wzyjerry/auth/api/user/v1"
	"github.com/wzyjerry/auth/internal/conf"
	"github.com/wzyjerry/auth/internal/ent"
	"github.com/wzyjerry/auth/internal/ent/schema/authenticatorNested"
	"github.com/wzyjerry/auth/internal/util"
	"golang.org/x/crypto/bcrypt"
)

type ThirdParty[UniqueType any] interface {
	Login(ctx context.Context, code string) (BasicInfo[UniqueType], error)
}

type BasicInfo[UniqueType any] interface {
	GetUnique() UniqueType
	GetNickname() string
	GetAvatar() func(ctx context.Context) (string, error)
}

type GithubProvider = ThirdParty[int64]
type MicrosoftProvider = ThirdParty[string]

type LoginUsecase struct {
	repo              UserRepo
	conf              *conf.Security
	tokenHelper       *util.TokenHelper
	aliyunHelper      *util.AliyunHelper
	registerUsecase   *RegisterUsecase
	githubProvider    GithubProvider
	microsoftProvider MicrosoftProvider
}

func NewLoginUsecase(
	repo UserRepo,
	conf *conf.Security,
	tokenHelper *util.TokenHelper,
	aliyunHelper *util.AliyunHelper,
	registerUsecase *RegisterUsecase,
	githubProvider GithubProvider,
	microsoftProvider MicrosoftProvider,
) *LoginUsecase {
	return &LoginUsecase{
		repo:              repo,
		conf:              conf,
		tokenHelper:       tokenHelper,
		aliyunHelper:      aliyunHelper,
		registerUsecase:   registerUsecase,
		githubProvider:    githubProvider,
		microsoftProvider: microsoftProvider,
	}
}

func (uc *LoginUsecase) verifyPassword(hash string, password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) == nil
}

func (uc *LoginUsecase) password(ctx context.Context, kind authenticatorNested.Kind, unique *string, password string) (string, error) {
	if unique == nil {
		return "", ErrUniqueRequired
	}
	auth, err := uc.repo.GetAuthenticator(ctx, kind, *unique)
	if err != nil {
		if ent.IsNotFound(err) {
			return "", ErrPasswordLogin
		}
		return "", err
	}
	userPassword, ancestorId, err := uc.repo.GetUserPasswordAndAncestorId(ctx, *auth.UserID)
	if err != nil {
		return "", err
	}
	if userPassword == nil || !uc.verifyPassword(*userPassword, password) {
		return "", ErrPasswordLogin
	}
	return ancestorId, nil
}

// accountPassword 账户密码登录
func (uc *LoginUsecase) accountPassword(ctx context.Context, username *string, password string) (string, error) {
	return uc.password(ctx, authenticatorNested.Kind_KIND_ACCOUNT, username, password)
}

// AccountPassword 邮箱密码登录
func (uc *LoginUsecase) emailPassword(ctx context.Context, email *string, password string) (string, error) {
	return uc.password(ctx, authenticatorNested.Kind_KIND_EMAIL, email, password)
}

// AccountPassword 手机密码登录
func (uc *LoginUsecase) phonePassword(ctx context.Context, phone *string, password string) (string, error) {
	return uc.password(ctx, authenticatorNested.Kind_KIND_PHONE, phone, password)
}

// PreEmail 限定登录使用验证码，email为唯一值
func (uc *LoginUsecase) PreEmail(ctx context.Context, email string) error {
	code := util.Rnd6()
	if err := uc.aliyunHelper.SendEmail(email, uc.aliyunHelper.NewEmailHtmlCode(code)); err != nil {
		return err
	}
	if err := uc.repo.CacheLoginEmail(ctx, email, code, uc.conf.Expiration.Code.AsDuration()); err != nil {
		return err
	}
	return nil
}

// PrePhone 限定登录使用验证码，phone为唯一值
func (uc *LoginUsecase) PrePhone(ctx context.Context, phone string) error {
	code := util.Rnd6()
	// 区分国内国际模板
	switch {
	case strings.HasPrefix(phone, "+86"):
		if err := uc.aliyunHelper.SendSms(phone, uc.aliyunHelper.NewSms228845627(code)); err != nil {
			return err
		}
	default:
		if err := uc.aliyunHelper.SendSms(phone, uc.aliyunHelper.NewSms228852216(code)); err != nil {
			return err
		}
	}
	if err := uc.repo.CacheLoginPhone(ctx, phone, code, uc.conf.Expiration.Code.AsDuration()); err != nil {
		return err
	}
	return nil
}

func (uc *LoginUsecase) emailCode(ctx context.Context, email *string, code string) (string, error) {
	if email == nil {
		return "", ErrUniqueRequired
	}
	verified, err := uc.repo.VerifyLoginEmailCode(ctx, *email, code)
	if err != nil {
		return "", err
	}
	if !verified {
		return "", ErrCodeMismatch
	}
	auth, err := uc.repo.GetAuthenticator(ctx, authenticatorNested.Kind_KIND_EMAIL, email)
	var userId string
	if err != nil {
		id, err := uc.registerUsecase.Email(ctx, *email, nil, *email)
		if err != nil {
			return "", err
		}
		userId = id
	} else {
		userId = *auth.UserID
	}
	return uc.repo.GetAncestorId(ctx, userId)
}

func (uc *LoginUsecase) phoneCode(ctx context.Context, phone *string, code string) (string, error) {
	if phone == nil {
		return "", ErrUniqueRequired
	}
	verified, err := uc.repo.VerifyLoginPhoneCode(ctx, *phone, code)
	if err != nil {
		return "", err
	}
	if !verified {
		return "", ErrCodeMismatch
	}
	auth, err := uc.repo.GetAuthenticator(ctx, authenticatorNested.Kind_KIND_PHONE, phone)
	var userId string
	if err != nil {
		id, err := uc.registerUsecase.Phone(ctx, *phone, nil, *phone)
		if err != nil {
			return "", err
		}
		userId = id
	} else {
		userId = *auth.UserID
	}
	return uc.repo.GetAncestorId(ctx, userId)
}

func (uc *LoginUsecase) githubCode(ctx context.Context, code string) (string, error) {
	info, err := uc.githubProvider.Login(ctx, code)
	if err != nil {
		return "", err
	}
	auth, err := uc.repo.GetAuthenticator(ctx, authenticatorNested.Kind_KIND_GITHUB, info.GetUnique())
	var userId string
	if err != nil {
		if ent.IsNotFound(err) {
			id, err := uc.registerUsecase.Github(ctx, info.GetUnique(), info.GetNickname(), info.GetAvatar())
			if err != nil {
				return "", err
			}
			userId = id
		} else {
			return "", err
		}
	} else {
		userId = *auth.UserID
	}
	return uc.repo.GetAncestorId(ctx, userId)
}

func (uc *LoginUsecase) microsoftCode(ctx context.Context, code string) (string, error) {
	info, err := uc.microsoftProvider.Login(ctx, code)
	if err != nil {
		return "", err
	}
	auth, err := uc.repo.GetAuthenticator(ctx, authenticatorNested.Kind_KIND_MICROSOFT, info.GetUnique())
	var userId string
	if err != nil {
		if ent.IsNotFound(err) {
			id, err := uc.registerUsecase.Microsoft(ctx, info.GetUnique(), info.GetNickname(), info.GetAvatar())
			if err != nil {
				return "", err
			}
			userId = id
		} else {
			return "", err
		}
	} else {
		userId = *auth.UserID
	}
	return uc.repo.GetAncestorId(ctx, userId)
}

func (uc *LoginUsecase) Login(ctx context.Context, _type v1.Type, method v1.Method, unique *string, secret string) (string, error) {
	var ancestorId string
	switch method {
	case v1.Method_METHOD_PASSWORD:
		switch _type {
		case v1.Type_TYPE_ACCOUNT:
			result, err := uc.accountPassword(ctx, unique, secret)
			if err != nil {
				return ancestorId, err
			}
			ancestorId = result
		case v1.Type_TYPE_EMAIL:
			result, err := uc.emailPassword(ctx, unique, secret)
			if err != nil {
				return ancestorId, err
			}
			ancestorId = result
		case v1.Type_TYPE_PHONE:
			result, err := uc.phonePassword(ctx, unique, secret)
			if err != nil {
				return ancestorId, err
			}
			ancestorId = result
		default:
			return ancestorId, ErrUnknownKind
		}
	case v1.Method_METHOD_CODE:
		switch _type {
		case v1.Type_TYPE_EMAIL:
			result, err := uc.emailCode(ctx, unique, secret)
			if err != nil {
				return ancestorId, err
			}
			ancestorId = result
		case v1.Type_TYPE_PHONE:
			result, err := uc.phoneCode(ctx, unique, secret)
			if err != nil {
				return ancestorId, err
			}
			ancestorId = result
		case v1.Type_TYPE_GITHUB:
			result, err := uc.githubCode(ctx, secret)
			if err != nil {
				return ancestorId, err
			}
			ancestorId = result
		case v1.Type_TYPE_MICROSOFT:
			result, err := uc.microsoftCode(ctx, secret)
			if err != nil {
				return ancestorId, err
			}
			ancestorId = result
		default:
			return ancestorId, ErrUnknownKind
		}
	default:
		return ancestorId, ErrUnknownKind
	}
	return ancestorId, nil
}

func (uc *LoginUsecase) GenerateAccessToken(clientId string, ancestorId string) (string, error) {
	token := uc.tokenHelper.GenerateBasicToken(clientId, ancestorId)
	signed, err := uc.tokenHelper.SignJWT(token)
	if err != nil {
		return "", err
	}
	return signed, nil
}
