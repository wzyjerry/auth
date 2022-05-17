package biz

import (
	"context"
	"strings"

	"github.com/dlclark/regexp2"
	"github.com/wzyjerry/auth/internal/conf"
	"github.com/wzyjerry/auth/internal/ent"
	"github.com/wzyjerry/auth/internal/ent/schema/authenticatorNested"
	"github.com/wzyjerry/auth/internal/middleware"
	"github.com/wzyjerry/auth/internal/util"
	"golang.org/x/crypto/bcrypt"
)

var (
	passwordRegex *regexp2.Regexp
)

type RegisterUsecase struct {
	repo   UserRepo
	helper *util.AliyunHelper
	conf   *conf.Security
}

func NewRegisterUsecase(
	repo UserRepo,
	conf *conf.Security,
	helper *util.AliyunHelper,
) *RegisterUsecase {
	return &RegisterUsecase{
		repo:   repo,
		conf:   conf,
		helper: helper,
	}
}

func (uc *RegisterUsecase) validatePassword(password string) bool {
	// 确保password正则存在
	if passwordRegex == nil {
		passwordRegex = regexp2.MustCompile(uc.conf.PasswordRegex, regexp2.Singleline)
	}
	b, err := passwordRegex.MatchString(password)
	return err == nil && b
}

func (uc *RegisterUsecase) hashPassword(password string) string {
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash)
}

func (uc *RegisterUsecase) generateAccountAnchor(username string) *authenticatorNested.Anchor {
	return &authenticatorNested.Anchor{
		Account: &authenticatorNested.Account{
			Username: util.P(strings.ToLower(username)),
		},
	}
}

func (uc *RegisterUsecase) generateEmailAnchor(email string) *authenticatorNested.Anchor {
	return &authenticatorNested.Anchor{
		Email: &authenticatorNested.Email{
			Email: util.P(strings.ToLower(email)),
		},
	}
}

func (uc *RegisterUsecase) generatePhoneAnchor(phone string) *authenticatorNested.Anchor {
	return &authenticatorNested.Anchor{
		Phone: &authenticatorNested.Phone{
			Phone: util.P(phone),
		},
	}
}

// Account 用户账户注册，username为唯一值，password和nickname必填，avatar选填
func (uc *RegisterUsecase) Account(ctx context.Context, username string, password string, nickname string) (string, error) {
	ip, err := middleware.GetIp(ctx)
	if err != nil {
		return "", err
	}
	if !uc.validatePassword(password) {
		return "", ErrInvalidPassword
	}
	hash := util.P(uc.hashPassword(password))
	if _, err := uc.repo.GetAuthenticator(ctx, authenticatorNested.Kind_KIND_ACCOUNT, username); !ent.IsNotFound(err) {
		return "", ErrAuthenticatorConflict
	}
	return uc.repo.CreateUser(ctx,
		int32(authenticatorNested.Kind_KIND_ACCOUNT),
		uc.generateAccountAnchor(username),
		hash,
		nickname,
		ip,
		nil,
	)
}

// PreEmail 限定注册使用验证码，email为唯一值
func (uc *RegisterUsecase) PreEmail(ctx context.Context, email string) error {
	if _, err := uc.repo.GetAuthenticator(ctx, authenticatorNested.Kind_KIND_EMAIL, email); !ent.IsNotFound(err) {
		return ErrAuthenticatorConflict
	}
	code := util.Rnd6()
	if err := uc.helper.SendEmail(email, uc.helper.NewEmailHtmlCode(code)); err != nil {
		return err
	}
	if err := uc.repo.CacheRegisterEmail(ctx, email, code); err != nil {
		return err
	}
	return nil
}

// PrePhone 限定注册使用验证码，phone为唯一值
func (uc *RegisterUsecase) PrePhone(ctx context.Context, phone string) error {
	if _, err := uc.repo.GetAuthenticator(ctx, authenticatorNested.Kind_KIND_PHONE, phone); !ent.IsNotFound(err) {
		return ErrAuthenticatorConflict
	}
	code := util.Rnd6()
	// 区分国内国际模板
	switch {
	case strings.HasPrefix(phone, "+86"):
		if err := uc.helper.SendSms(phone, uc.helper.NewSms228845627(code)); err != nil {
			return err
		}
	default:
		if err := uc.helper.SendSms(phone, uc.helper.NewSms228852216(code)); err != nil {
			return err
		}
	}
	if err := uc.repo.CacheRegisterPhone(ctx, phone, code); err != nil {
		return err
	}
	return nil
}

// email 邮箱注册，email为唯一值，nickname必填，password和avatar选填
func (uc *RegisterUsecase) Email(ctx context.Context, email string, password *string, nickname string) (string, error) {
	ip, err := middleware.GetIp(ctx)
	if err != nil {
		return "", err
	}
	var hash *string
	if password != nil {
		if !uc.validatePassword(*password) {
			return "", ErrInvalidPassword
		}
		hash = util.P(uc.hashPassword(*password))
	}
	if _, err := uc.repo.GetAuthenticator(ctx, authenticatorNested.Kind_KIND_EMAIL, email); !ent.IsNotFound(err) {
		return "", ErrAuthenticatorConflict
	}
	return uc.repo.CreateUser(ctx,
		int32(authenticatorNested.Kind_KIND_EMAIL),
		uc.generateEmailAnchor(email),
		hash,
		nickname,
		ip,
		nil,
	)
}

// phone 手机注册，phone为唯一值，nickname必填，password和avatar选填
func (uc *RegisterUsecase) Phone(ctx context.Context, phone string, password *string, nickname string) (string, error) {
	ip, err := middleware.GetIp(ctx)
	if err != nil {
		return "", err
	}
	var hash *string
	if password != nil {
		if !uc.validatePassword(*password) {
			return "", ErrInvalidPassword
		}
		hash = util.P(uc.hashPassword(*password))
	}
	if _, err := uc.repo.GetAuthenticator(ctx, authenticatorNested.Kind_KIND_PHONE, phone); !ent.IsNotFound(err) {
		return "", ErrAuthenticatorConflict
	}
	return uc.repo.CreateUser(ctx,
		int32(authenticatorNested.Kind_KIND_PHONE),
		uc.generatePhoneAnchor(phone),
		hash,
		nickname,
		ip,
		nil,
	)
}

// Phone 手机密码注册，phone为唯一值，nickname、code和password必填
func (uc *RegisterUsecase) PhonePassword(ctx context.Context, phone string, code string, password string, nickname string) (string, error) {
	verified, err := uc.repo.VerifyRegisterPhoneCode(ctx, phone, code)
	if err != nil || !verified {
		return "", ErrCodeMismatch
	}
	return uc.Phone(ctx, phone, &password, nickname)
}

// Email 邮箱密码注册，email为唯一值，nickname、code和password必填
func (uc *RegisterUsecase) EmailPassword(ctx context.Context, email string, code string, password string, nickname string) (string, error) {
	verified, err := uc.repo.VerifyRegisterEmailCode(ctx, email, code)
	if err != nil || !verified {
		return "", ErrCodeMismatch
	}
	return uc.Email(ctx, email, &password, nickname)
}
