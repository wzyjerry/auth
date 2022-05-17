package biz

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

type LoginUsecase struct {
	repo            UserRepo
	conf            *conf.Security
	tokenHelper     *util.TokenHelper
	aliyunHelper    *util.AliyunHelper
	registerUsecase *RegisterUsecase
}

func NewLoginUsecase(
	repo UserRepo,
	conf *conf.Security,
	tokenHelper *util.TokenHelper,
	aliyunHelper *util.AliyunHelper,
	registerUsecase *RegisterUsecase,
) *LoginUsecase {
	return &LoginUsecase{
		repo:            repo,
		conf:            conf,
		tokenHelper:     tokenHelper,
		aliyunHelper:    aliyunHelper,
		registerUsecase: registerUsecase,
	}
}

func (uc *LoginUsecase) verifyPassword(hash string, password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) == nil
}

func (uc *LoginUsecase) password(ctx context.Context, kind authenticatorNested.Kind, unique string, password string) (*ent.User, error) {
	auth, err := uc.repo.GetAuthenticator(ctx, kind, unique)
	if err != nil {
		return nil, ErrPasswordLogin
	}
	user, err := uc.repo.GetUser(ctx, *auth.UserID)
	if err != nil {
		return nil, ErrInternalInternalServerError
	}
	if user.Password == nil || !uc.verifyPassword(*user.Password, password) {
		return nil, ErrPasswordLogin
	}
	return user, nil
}

// accountPassword 账户密码登录
func (uc *LoginUsecase) accountPassword(ctx context.Context, username string, password string) (*ent.User, error) {
	return uc.password(ctx, authenticatorNested.Kind_KIND_ACCOUNT, username, password)
}

// AccountPassword 邮箱密码登录
func (uc *LoginUsecase) emailPassword(ctx context.Context, email string, password string) (*ent.User, error) {
	return uc.password(ctx, authenticatorNested.Kind_KIND_EMAIL, email, password)
}

// AccountPassword 手机密码登录
func (uc *LoginUsecase) phonePassword(ctx context.Context, phone string, password string) (*ent.User, error) {
	return uc.password(ctx, authenticatorNested.Kind_KIND_PHONE, phone, password)
}

// PreEmail 限定登录使用验证码，email为唯一值
func (uc *LoginUsecase) PreEmail(ctx context.Context, email string) error {
	code := util.Rnd6()
	if err := uc.aliyunHelper.SendEmail(email, uc.aliyunHelper.NewEmailHtmlCode(code)); err != nil {
		return err
	}
	if err := uc.repo.CacheLoginEmail(ctx, email, code); err != nil {
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
	if err := uc.repo.CacheLoginPhone(ctx, phone, code); err != nil {
		return err
	}
	return nil
}

func (uc *LoginUsecase) emailCode(ctx context.Context, email string, code string) (*ent.User, error) {
	verified, err := uc.repo.VerifyLoginEmailCode(ctx, email, code)
	if err != nil || !verified {
		return nil, ErrCodeMismatch
	}
	auth, err := uc.repo.GetAuthenticator(ctx, authenticatorNested.Kind_KIND_EMAIL, email)
	var userId string
	if err != nil {
		id, err := uc.registerUsecase.Email(ctx, email, nil, email)
		if err != nil {
			return nil, err
		}
		userId = id
	} else {
		userId = *auth.UserID
	}
	user, err := uc.repo.GetUser(ctx, userId)
	if err != nil {
		return nil, ErrInternalInternalServerError
	}
	return user, nil
}

func (uc *LoginUsecase) phoneCode(ctx context.Context, phone string, code string) (*ent.User, error) {
	verified, err := uc.repo.VerifyLoginPhoneCode(ctx, phone, code)
	if err != nil || !verified {
		return nil, ErrCodeMismatch
	}
	auth, err := uc.repo.GetAuthenticator(ctx, authenticatorNested.Kind_KIND_PHONE, phone)
	var userId string
	if err != nil {
		id, err := uc.registerUsecase.Phone(ctx, phone, nil, phone)
		if err != nil {
			return nil, err
		}
		userId = id
	} else {
		userId = *auth.UserID
	}
	user, err := uc.repo.GetUser(ctx, userId)
	if err != nil {
		return nil, ErrInternalInternalServerError
	}
	return user, nil
}

func (uc *LoginUsecase) Login(ctx context.Context, _type v1.Type, method v1.Method, unique string, secret string) (*ent.User, error) {
	var user *ent.User
	switch method {
	case v1.Method_METHOD_PASSWORD:
		switch _type {
		case v1.Type_TYPE_ACCOUNT:
			result, err := uc.accountPassword(ctx, unique, secret)
			if err != nil {
				return nil, err
			}
			user = result
		case v1.Type_TYPE_EMAIL:
			result, err := uc.emailPassword(ctx, unique, secret)
			if err != nil {
				return nil, err
			}
			user = result
		case v1.Type_TYPE_PHONE:
			result, err := uc.phonePassword(ctx, unique, secret)
			if err != nil {
				return nil, err
			}
			user = result
		default:
			return nil, ErrInternalInternalServerError
		}
	case v1.Method_METHOD_CODE:
		switch _type {
		case v1.Type_TYPE_EMAIL:
			result, err := uc.emailCode(ctx, unique, secret)
			if err != nil {
				return nil, err
			}
			user = result
		case v1.Type_TYPE_PHONE:
			result, err := uc.phoneCode(ctx, unique, secret)
			if err != nil {
				return nil, err
			}
			user = result
		default:
			return nil, ErrInternalInternalServerError
		}
	default:
		return nil, ErrInternalInternalServerError
	}
	return user, nil
}

func (uc *LoginUsecase) GenerateAccessToken(clientId string, user *ent.User) (string, error) {
	token := uc.tokenHelper.GenerateAccessToken(clientId, *user.AncestorID)
	signed, err := uc.tokenHelper.SignJWT(token)
	if err != nil {
		return "", ErrInternalInternalServerError
	}
	return signed, nil
}
