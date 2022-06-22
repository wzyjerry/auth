package userBiz

import (
	"context"
	"strings"
	"time"

	"github.com/dlclark/regexp2"
	"github.com/go-kratos/kratos/v2/log"
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
	logger *log.Helper
}

func NewRegisterUsecase(
	repo UserRepo,
	conf *conf.Security,
	helper *util.AliyunHelper,
	logger log.Logger,
) *RegisterUsecase {
	return &RegisterUsecase{
		repo:   repo,
		conf:   conf,
		helper: helper,
		logger: log.NewHelper(logger),
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

// Account 用户账户注册，username为唯一值，password和nickname必填
func (uc *RegisterUsecase) Account(ctx context.Context, username string, password string, nickname string) (string, error) {
	ip, err := middleware.GetIp(ctx)
	if err != nil {
		return "", err
	}
	if !uc.validatePassword(password) {
		return "", ErrInvalidPassword
	}
	hash := util.P(uc.hashPassword(password))
	if _, err := uc.repo.GetAuthenticator(ctx, authenticatorNested.Kind_KIND_ACCOUNT, username); err != nil {
		if ent.IsNotFound(err) {
			return uc.repo.CreateUser(ctx,
				int32(authenticatorNested.Kind_KIND_ACCOUNT),
				&authenticatorNested.Unique{
					Account: util.P(strings.ToLower(username)),
				},
				hash,
				nickname,
				ip,
			)
		}
		return "", err
	}
	return "", ErrAuthenticatorConflict
}

// PreEmail 限定注册使用验证码，email为唯一值
func (uc *RegisterUsecase) PreEmail(ctx context.Context, email string) error {
	if _, err := uc.repo.GetAuthenticator(ctx, authenticatorNested.Kind_KIND_EMAIL, email); err != nil {
		if ent.IsNotFound(err) {
			code := util.Rnd6()
			if err := uc.helper.SendEmail(email, uc.helper.NewEmailHtmlCode(code)); err != nil {
				return err
			}
			if err := uc.repo.CacheRegisterEmail(ctx, email, code, uc.conf.Expiration.Code.AsDuration()); err != nil {
				return err
			}
			return nil
		}
		return err
	}
	return ErrAuthenticatorConflict
}

// PrePhone 限定注册使用验证码，phone为唯一值
func (uc *RegisterUsecase) PrePhone(ctx context.Context, phone string) error {
	if _, err := uc.repo.GetAuthenticator(ctx, authenticatorNested.Kind_KIND_PHONE, phone); err != nil {
		if ent.IsNotFound(err) {
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
			if err := uc.repo.CacheRegisterPhone(ctx, phone, code, uc.conf.Expiration.Code.AsDuration()); err != nil {
				return err
			}
			return nil
		}
		return err
	}
	return ErrAuthenticatorConflict
}

// email 邮箱注册，email为唯一值，nickname必填，password选填
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
	if _, err := uc.repo.GetAuthenticator(ctx, authenticatorNested.Kind_KIND_EMAIL, email); err != nil {
		if ent.IsNotFound(err) {
			return uc.repo.CreateUser(ctx,
				int32(authenticatorNested.Kind_KIND_EMAIL),
				&authenticatorNested.Unique{
					Email: util.P(strings.ToLower(email)),
				},
				hash,
				nickname,
				ip,
			)
		}
		return "", err
	}
	return "", ErrAuthenticatorConflict
}

// phone 手机注册，phone为唯一值，nickname必填，password选填
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
	if _, err := uc.repo.GetAuthenticator(ctx, authenticatorNested.Kind_KIND_PHONE, phone); err != nil {
		if ent.IsNotFound(err) {
			return uc.repo.CreateUser(ctx,
				int32(authenticatorNested.Kind_KIND_PHONE),
				&authenticatorNested.Unique{
					Phone: &phone,
				},
				hash,
				nickname,
				ip,
			)
		}
		return "", err
	}
	return "", ErrAuthenticatorConflict
}

func (uc *RegisterUsecase) Github(ctx context.Context, id int64, nickname string, avatar func(ctx context.Context) (string, error)) (string, error) {
	ip, err := middleware.GetIp(ctx)
	if err != nil {
		return "", err
	}
	if _, err := uc.repo.GetAuthenticator(ctx, authenticatorNested.Kind_KIND_GITHUB, id); err != nil {
		if ent.IsNotFound(err) {
			id, err := uc.repo.CreateUser(ctx,
				int32(authenticatorNested.Kind_KIND_GITHUB),
				&authenticatorNested.Unique{
					Github: &id,
				},
				nil,
				nickname,
				ip,
			)
			if err != nil {
				return "", err
			}
			go func() {
				ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
				defer cancel()
				base64, err := avatar(ctx)
				if err == nil {
					uc.repo.CreateAvatar(ctx, id, base64)
				} else {
					uc.logger.Errorf("拉取Github头像失败: %v", err)
				}
			}()
			return id, nil
		}
		return "", err

	}
	return "", ErrAuthenticatorConflict
}

func (uc *RegisterUsecase) Microsoft(ctx context.Context, id string, nickname string, avatar func(ctx context.Context) (string, error)) (string, error) {
	ip, err := middleware.GetIp(ctx)
	if err != nil {
		return "", err
	}
	if _, err := uc.repo.GetAuthenticator(ctx, authenticatorNested.Kind_KIND_MICROSOFT, id); err != nil {
		if ent.IsNotFound(err) {
			id, err := uc.repo.CreateUser(ctx,
				int32(authenticatorNested.Kind_KIND_MICROSOFT),
				&authenticatorNested.Unique{
					Microsoft: &id,
				},
				nil,
				nickname,
				ip,
			)
			if err != nil {
				return "", err
			}
			go func() {
				ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
				defer cancel()
				base64, err := avatar(ctx)
				if err == nil {
					uc.repo.CreateAvatar(ctx, id, base64)
				} else {
					uc.logger.Errorf("拉取Microsoft头像失败: %v", err)
				}
			}()
			return id, nil
		}
		return "", err

	}
	return "", ErrAuthenticatorConflict
}

// Phone 手机密码注册，phone为唯一值，nickname、code和password必填
func (uc *RegisterUsecase) PhonePassword(ctx context.Context, phone string, code string, password string, nickname string) (string, error) {
	verified, err := uc.repo.VerifyRegisterPhoneCode(ctx, phone, code)
	if err != nil {
		return "", err
	}
	if !verified {
		return "", ErrCodeMismatch
	}
	return uc.Phone(ctx, phone, &password, nickname)
}

// Email 邮箱密码注册，email为唯一值，nickname、code和password必填
func (uc *RegisterUsecase) EmailPassword(ctx context.Context, email string, code string, password string, nickname string) (string, error) {
	verified, err := uc.repo.VerifyRegisterEmailCode(ctx, email, code)
	if err != nil {
		return "", err
	}
	if !verified {
		return "", ErrCodeMismatch
	}
	return uc.Email(ctx, email, &password, nickname)
}
