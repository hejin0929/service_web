// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.6.1
// - protoc             v3.21.12
// source: auth/v1/auth.proto

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationAuthExitUsersLogin = "/api.auth.v1.Auth/ExitUsersLogin"
const OperationAuthLoginUsers = "/api.auth.v1.Auth/LoginUsers"
const OperationAuthPatchUsersLogin = "/api.auth.v1.Auth/PatchUsersLogin"

type AuthHTTPServer interface {
	ExitUsersLogin(context.Context, *ExitUsersLoginRequest) (*ExitUsersLoginReply, error)
	LoginUsers(context.Context, *LoginUsersRequest) (*LoginUsersReply, error)
	PatchUsersLogin(context.Context, *PatchUsersLoginRequest) (*PatchUsersLoginReply, error)
}

func RegisterAuthHTTPServer(s *http.Server, srv AuthHTTPServer) {
	r := s.Route("/")
	r.POST("/api/v1/login", _Auth_LoginUsers0_HTTP_Handler(srv))
	r.DELETE("/api/v1/login", _Auth_ExitUsersLogin0_HTTP_Handler(srv))
	r.PATCH("/api/v1/login", _Auth_PatchUsersLogin0_HTTP_Handler(srv))
}

func _Auth_LoginUsers0_HTTP_Handler(srv AuthHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in LoginUsersRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAuthLoginUsers)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.LoginUsers(ctx, req.(*LoginUsersRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*LoginUsersReply)
		return ctx.Result(200, reply)
	}
}

func _Auth_ExitUsersLogin0_HTTP_Handler(srv AuthHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ExitUsersLoginRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAuthExitUsersLogin)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ExitUsersLogin(ctx, req.(*ExitUsersLoginRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ExitUsersLoginReply)
		return ctx.Result(200, reply)
	}
}

func _Auth_PatchUsersLogin0_HTTP_Handler(srv AuthHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in PatchUsersLoginRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAuthPatchUsersLogin)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.PatchUsersLogin(ctx, req.(*PatchUsersLoginRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*PatchUsersLoginReply)
		return ctx.Result(200, reply)
	}
}

type AuthHTTPClient interface {
	ExitUsersLogin(ctx context.Context, req *ExitUsersLoginRequest, opts ...http.CallOption) (rsp *ExitUsersLoginReply, err error)
	LoginUsers(ctx context.Context, req *LoginUsersRequest, opts ...http.CallOption) (rsp *LoginUsersReply, err error)
	PatchUsersLogin(ctx context.Context, req *PatchUsersLoginRequest, opts ...http.CallOption) (rsp *PatchUsersLoginReply, err error)
}

type AuthHTTPClientImpl struct {
	cc *http.Client
}

func NewAuthHTTPClient(client *http.Client) AuthHTTPClient {
	return &AuthHTTPClientImpl{client}
}

func (c *AuthHTTPClientImpl) ExitUsersLogin(ctx context.Context, in *ExitUsersLoginRequest, opts ...http.CallOption) (*ExitUsersLoginReply, error) {
	var out ExitUsersLoginReply
	pattern := "/api/v1/login"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationAuthExitUsersLogin))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AuthHTTPClientImpl) LoginUsers(ctx context.Context, in *LoginUsersRequest, opts ...http.CallOption) (*LoginUsersReply, error) {
	var out LoginUsersReply
	pattern := "/api/v1/login"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationAuthLoginUsers))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AuthHTTPClientImpl) PatchUsersLogin(ctx context.Context, in *PatchUsersLoginRequest, opts ...http.CallOption) (*PatchUsersLoginReply, error) {
	var out PatchUsersLoginReply
	pattern := "/api/v1/login"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationAuthPatchUsersLogin))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PATCH", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}