// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// protoc-gen-go-http v2.2.2

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

type LoginServiceHTTPServer interface {
	Login(context.Context, *LoginRequest) (*LoginReply, error)
	PreEmail(context.Context, *LoginPreEmailRequest) (*emptypb.Empty, error)
	PrePhone(context.Context, *LoginPrePhoneRequest) (*emptypb.Empty, error)
}

func RegisterLoginServiceHTTPServer(s *http.Server, srv LoginServiceHTTPServer) {
	r := s.Route("/")
	r.POST("/user/v1/login/pre_phone", _LoginService_PrePhone0_HTTP_Handler(srv))
	r.POST("/user/v1/login/pre_email", _LoginService_PreEmail0_HTTP_Handler(srv))
	r.POST("/user/v1/login/login", _LoginService_Login0_HTTP_Handler(srv))
}

func _LoginService_PrePhone0_HTTP_Handler(srv LoginServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in LoginPrePhoneRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.user.v1.LoginService/PrePhone")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.PrePhone(ctx, req.(*LoginPrePhoneRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*emptypb.Empty)
		return ctx.Result(200, reply)
	}
}

func _LoginService_PreEmail0_HTTP_Handler(srv LoginServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in LoginPreEmailRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.user.v1.LoginService/PreEmail")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.PreEmail(ctx, req.(*LoginPreEmailRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*emptypb.Empty)
		return ctx.Result(200, reply)
	}
}

func _LoginService_Login0_HTTP_Handler(srv LoginServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in LoginRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.user.v1.LoginService/Login")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Login(ctx, req.(*LoginRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*LoginReply)
		return ctx.Result(200, reply)
	}
}

type LoginServiceHTTPClient interface {
	Login(ctx context.Context, req *LoginRequest, opts ...http.CallOption) (rsp *LoginReply, err error)
	PreEmail(ctx context.Context, req *LoginPreEmailRequest, opts ...http.CallOption) (rsp *emptypb.Empty, err error)
	PrePhone(ctx context.Context, req *LoginPrePhoneRequest, opts ...http.CallOption) (rsp *emptypb.Empty, err error)
}

type LoginServiceHTTPClientImpl struct {
	cc *http.Client
}

func NewLoginServiceHTTPClient(client *http.Client) LoginServiceHTTPClient {
	return &LoginServiceHTTPClientImpl{client}
}

func (c *LoginServiceHTTPClientImpl) Login(ctx context.Context, in *LoginRequest, opts ...http.CallOption) (*LoginReply, error) {
	var out LoginReply
	pattern := "/user/v1/login/login"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/api.user.v1.LoginService/Login"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *LoginServiceHTTPClientImpl) PreEmail(ctx context.Context, in *LoginPreEmailRequest, opts ...http.CallOption) (*emptypb.Empty, error) {
	var out emptypb.Empty
	pattern := "/user/v1/login/pre_email"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/api.user.v1.LoginService/PreEmail"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *LoginServiceHTTPClientImpl) PrePhone(ctx context.Context, in *LoginPrePhoneRequest, opts ...http.CallOption) (*emptypb.Empty, error) {
	var out emptypb.Empty
	pattern := "/user/v1/login/pre_phone"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/api.user.v1.LoginService/PrePhone"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
