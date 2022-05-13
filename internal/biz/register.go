package biz

import (
	"context"
	"net/http"
	"strings"

	"github.com/dlclark/regexp2"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/wzyjerry/auth/internal/conf"
	"github.com/wzyjerry/auth/internal/ent"
	"github.com/wzyjerry/auth/internal/ent/schema/authenticatorNested"
	"github.com/wzyjerry/auth/internal/middleware"
	"github.com/wzyjerry/auth/internal/util"
	"golang.org/x/crypto/bcrypt"
)

var (
	passwordRegex            *regexp2.Regexp
	ErrUnknownKind           = errors.New(http.StatusInternalServerError, "UNKNOWN_KIND", "unknown kind")
	ErrInvalidPassword       = errors.New(http.StatusBadRequest, "INVALID_PASSWORD", "incalid password")
	ErrCodeMismatch          = errors.New(http.StatusBadRequest, "CODE_MISMATCH", "code mismatch")
	ErrAuthenticatorConflict = errors.New(http.StatusConflict, "AUTHENTICATOR_CONFLICT", "authenticator conflict")
)

type RegisterRepo interface {
	// redis部分
	CachePreEmail(ctx context.Context, email string, code string) error
	CachePrePhone(ctx context.Context, phone string, code string) error
	VerifyPreEmailCode(ctx context.Context, email string, code string) (bool, error)
	VerifyPrePhoneCode(ctx context.Context, phone string, code string) (bool, error)
	// db部分
	GetAuthenticator(ctx context.Context, kind authenticatorNested.Kind, unique string) (*ent.Authenticator, error)
	CreateUser(ctx context.Context, kind int32, anchor *authenticatorNested.Anchor, password *string, nickname string, ip string, avatar *string) (string, error)
}

type RegisterUsecase struct {
	repo   RegisterRepo
	helper *AliyunHelper
	conf   *conf.Security
	log    *log.Helper
}

func NewRegisterUsecase(
	repo RegisterRepo,
	conf *conf.Security,
	helper *AliyunHelper,
	logger log.Logger,
) *RegisterUsecase {
	return &RegisterUsecase{
		repo:   repo,
		conf:   conf,
		helper: helper,
		log:    log.NewHelper(logger),
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
	if err := uc.helper.SendEmail(email, NewEmailHtmlCode(code)); err != nil {
		return err
	}
	if err := uc.repo.CachePreEmail(ctx, email, code); err != nil {
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
		if err := uc.helper.SendSms(phone, NewSms228845627(code)); err != nil {
			return err
		}
	default:
		if err := uc.helper.SendSms(phone, NewSms228852216(code)); err != nil {
			return err
		}
	}
	if err := uc.repo.CachePrePhone(ctx, phone, code); err != nil {
		return err
	}
	return nil
}

// email 邮箱注册，email为唯一值，nickname必填，password和avatar选填
func (uc *RegisterUsecase) email(ctx context.Context, email string, password *string, nickname string) (string, error) {
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
func (uc *RegisterUsecase) phone(ctx context.Context, phone string, password *string, nickname string) (string, error) {
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
func (uc *RegisterUsecase) Phone(ctx context.Context, phone string, code string, password string, nickname string) (string, error) {
	verified, err := uc.repo.VerifyPrePhoneCode(ctx, phone, code)
	if err != nil || !verified {
		return "", ErrCodeMismatch
	}
	return uc.phone(ctx, phone, &password, nickname)
}

// Email 邮箱密码注册，email为唯一值，nickname、code和password必填
func (uc *RegisterUsecase) Email(ctx context.Context, email string, code string, password string, nickname string) (string, error) {
	verified, err := uc.repo.VerifyPreEmailCode(ctx, email, code)
	if err != nil || !verified {
		return "", ErrCodeMismatch
	}
	return uc.email(ctx, email, &password, nickname)
}
