package biz

import (
	"fmt"
	"net/http"
	"os"

	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	dm20151123 "github.com/alibabacloud-go/dm-20151123/client"
	dysmsapi20170525 "github.com/alibabacloud-go/dysmsapi-20170525/v2/client"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/wzyjerry/auth/internal/conf"
	"github.com/wzyjerry/auth/internal/util"
)

type EmailTemplate struct {
	subject string
	html    *string
	text    *string
}

// Html验证码
func NewEmailHtmlCode(code string) *EmailTemplate {
	return &EmailTemplate{
		subject: "AMiner邮件验证码 (Your AMiner verification code)",
		html: util.P(fmt.Sprintf(`<!DOCTYPE html>
		<html lang="en">
		<head>
		   <meta charset="UTF-8">
		   <meta http-equiv="X-UA-Compatible" content="IE=edge">
		   <meta name="viewport" content="width=device-width, initial-scale=1.0">
		   <title>Document</title>
		</head>
		<body>
		   <div>您的邮箱验证码见下文，验证码有效期是10分钟，请勿泄露。</div>
		   <div>Here is the AMiner verification code. It expires in 10 minutes. Don't share this code with anyone else.</div>
		   <div style="color: #539ae6;font-size: 30px">%s</div>
		</body>
		</html>`, code)),
	}
}

// Text验证码
func NewEmailTextCode(code string) *EmailTemplate {
	return &EmailTemplate{
		subject: "AMiner邮件验证码 (Your AMiner verification code)",
		text: util.P(fmt.Sprintf(`您的邮箱验证码见下文，验证码有效期是10分钟，请勿泄露。
		Here is the AMiner verification code. It expires in 10 minutes. Don't share this code with anyone else.
		%s`, code)),
	}
}

type SmsTemplate struct {
	code  string
	param string
}

// 国内验证码
func NewSms228845627(code string) *SmsTemplate {
	return &SmsTemplate{
		code:  "SMS_228845627",
		param: fmt.Sprintf(`{"code":"%s"}`, code),
	}
}

// 国际验证码
func NewSms228852216(code string) *SmsTemplate {
	return &SmsTemplate{
		code:  "SMS_228852216",
		param: fmt.Sprintf(`{"code":"%s"}`, code),
	}
}

type AliyunHelper struct {
	emailClient *dm20151123.Client
	smsClient   *dysmsapi20170525.Client
	conf        *conf.Security
	log         *log.Helper
}

func NewAliyunHelper(
	logger log.Logger,
	conf *conf.Security,
) *AliyunHelper {
	log := log.NewHelper(logger)
	emailConfig := &openapi.Config{
		AccessKeyId:     util.P(os.Getenv(conf.Aliyun.Email.AccessKeyId)),
		AccessKeySecret: util.P(os.Getenv(conf.Aliyun.Email.AccessKeySecret)),
		Endpoint:        util.P("dm.aliyuncs.com"),
	}
	emailClient, err := dm20151123.NewClient(emailConfig)
	if err != nil {
		log.Fatalf("aliyun email fatal")
	}
	smsConfig := &openapi.Config{
		AccessKeyId:     util.P(os.Getenv(conf.Aliyun.Sms.AccessKeyId)),
		AccessKeySecret: util.P(os.Getenv(conf.Aliyun.Sms.AccessKeySecret)),
		Endpoint:        util.P("dysmsapi.aliyuncs.com"),
	}
	smsClient, err := dysmsapi20170525.NewClient(smsConfig)
	if err != nil {
		log.Fatalf("aliyun sms fatal")
	}
	return &AliyunHelper{
		log:         log,
		conf:        conf,
		emailClient: emailClient,
		smsClient:   smsClient,
	}
}

func (h *AliyunHelper) SendEmail(email string, template *EmailTemplate) error {
	request := &dm20151123.SingleSendMailRequest{
		AccountName:    &h.conf.Aliyun.Email.AccountName,
		AddressType:    util.P[int32](1),
		FromAlias:      &h.conf.Aliyun.Email.FromAlias,
		ReplyToAddress: util.P(true),
		Subject:        &template.subject,
		HtmlBody:       template.html,
		TextBody:       template.text,
		ToAddress:      &email,
	}
	if _, err := h.emailClient.SingleSendMail(request); err != nil {
		return errors.New(http.StatusBadRequest, "ERR_SEND_EMAIL", err.Error())
	}
	return nil
}

func (h *AliyunHelper) SendSms(phone string, template *SmsTemplate) error {
	request := &dysmsapi20170525.SendSmsRequest{
		PhoneNumbers:  &phone,
		SignName:      &h.conf.Aliyun.Sms.SignName,
		TemplateCode:  &template.code,
		TemplateParam: &template.param,
	}
	response, err := h.smsClient.SendSms(request)
	if err != nil {
		return err
	}
	if *response.Body.Code != "OK" {
		return errors.New(http.StatusBadRequest, "ERR_SEND_SMS", *response.Body.Message)
	}
	return nil
}
