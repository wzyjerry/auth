// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.15.6
// source: oauth2/v1/oauth.proto

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// OAuth2ServiceClient is the client API for OAuth2Service service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OAuth2ServiceClient interface {
	PreAuthorize(ctx context.Context, in *PreAuthorizeRequest, opts ...grpc.CallOption) (*PreAuthorizeReply, error)
	Authorize(ctx context.Context, in *AuthorizeRequest, opts ...grpc.CallOption) (*AuthorizeReply, error)
	Token(ctx context.Context, in *TokenRequest, opts ...grpc.CallOption) (*TokenReply, error)
}

type oAuth2ServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewOAuth2ServiceClient(cc grpc.ClientConnInterface) OAuth2ServiceClient {
	return &oAuth2ServiceClient{cc}
}

func (c *oAuth2ServiceClient) PreAuthorize(ctx context.Context, in *PreAuthorizeRequest, opts ...grpc.CallOption) (*PreAuthorizeReply, error) {
	out := new(PreAuthorizeReply)
	err := c.cc.Invoke(ctx, "/api.oauth2.v1.OAuth2Service/PreAuthorize", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *oAuth2ServiceClient) Authorize(ctx context.Context, in *AuthorizeRequest, opts ...grpc.CallOption) (*AuthorizeReply, error) {
	out := new(AuthorizeReply)
	err := c.cc.Invoke(ctx, "/api.oauth2.v1.OAuth2Service/Authorize", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *oAuth2ServiceClient) Token(ctx context.Context, in *TokenRequest, opts ...grpc.CallOption) (*TokenReply, error) {
	out := new(TokenReply)
	err := c.cc.Invoke(ctx, "/api.oauth2.v1.OAuth2Service/Token", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OAuth2ServiceServer is the server API for OAuth2Service service.
// All implementations must embed UnimplementedOAuth2ServiceServer
// for forward compatibility
type OAuth2ServiceServer interface {
	PreAuthorize(context.Context, *PreAuthorizeRequest) (*PreAuthorizeReply, error)
	Authorize(context.Context, *AuthorizeRequest) (*AuthorizeReply, error)
	Token(context.Context, *TokenRequest) (*TokenReply, error)
	mustEmbedUnimplementedOAuth2ServiceServer()
}

// UnimplementedOAuth2ServiceServer must be embedded to have forward compatible implementations.
type UnimplementedOAuth2ServiceServer struct {
}

func (UnimplementedOAuth2ServiceServer) PreAuthorize(context.Context, *PreAuthorizeRequest) (*PreAuthorizeReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PreAuthorize not implemented")
}
func (UnimplementedOAuth2ServiceServer) Authorize(context.Context, *AuthorizeRequest) (*AuthorizeReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Authorize not implemented")
}
func (UnimplementedOAuth2ServiceServer) Token(context.Context, *TokenRequest) (*TokenReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Token not implemented")
}
func (UnimplementedOAuth2ServiceServer) mustEmbedUnimplementedOAuth2ServiceServer() {}

// UnsafeOAuth2ServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OAuth2ServiceServer will
// result in compilation errors.
type UnsafeOAuth2ServiceServer interface {
	mustEmbedUnimplementedOAuth2ServiceServer()
}

func RegisterOAuth2ServiceServer(s grpc.ServiceRegistrar, srv OAuth2ServiceServer) {
	s.RegisterService(&OAuth2Service_ServiceDesc, srv)
}

func _OAuth2Service_PreAuthorize_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PreAuthorizeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OAuth2ServiceServer).PreAuthorize(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.oauth2.v1.OAuth2Service/PreAuthorize",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OAuth2ServiceServer).PreAuthorize(ctx, req.(*PreAuthorizeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OAuth2Service_Authorize_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthorizeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OAuth2ServiceServer).Authorize(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.oauth2.v1.OAuth2Service/Authorize",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OAuth2ServiceServer).Authorize(ctx, req.(*AuthorizeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OAuth2Service_Token_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OAuth2ServiceServer).Token(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.oauth2.v1.OAuth2Service/Token",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OAuth2ServiceServer).Token(ctx, req.(*TokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// OAuth2Service_ServiceDesc is the grpc.ServiceDesc for OAuth2Service service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OAuth2Service_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.oauth2.v1.OAuth2Service",
	HandlerType: (*OAuth2ServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PreAuthorize",
			Handler:    _OAuth2Service_PreAuthorize_Handler,
		},
		{
			MethodName: "Authorize",
			Handler:    _OAuth2Service_Authorize_Handler,
		},
		{
			MethodName: "Token",
			Handler:    _OAuth2Service_Token_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "oauth2/v1/oauth.proto",
}
