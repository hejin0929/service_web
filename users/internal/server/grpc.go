package server

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/validate"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	authV1 "users/api/auth/v1"
	v1 "users/api/users/v1"
	"users/internal/conf"
	"users/internal/service"
	"users/pkg"
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(c *conf.Server, greeter *service.UsersService, logger log.Logger, authService *service.AuthService) *grpc.Server {
	var opts = []grpc.ServerOption{
		grpc.Middleware(
			recovery.Recovery(),
			validate.Validator(),
		),
	}
	if c.Grpc.Network != "" {
		opts = append(opts, grpc.Network(c.Grpc.Network))
	}
	if c.Grpc.Addr != "" {
		opts = append(opts, grpc.Address(c.Grpc.Addr))
	}
	if c.Grpc.Timeout != nil {
		opts = append(opts, grpc.Timeout(c.Grpc.Timeout.AsDuration()))
	}

	opts = append(opts, grpc.Middleware(func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			if cc, ok := transport.FromServerContext(ctx); ok {

				token := cc.RequestHeader().Get("Authorization")
				if token == "" {
					return nil, errors.New("error：token deletion")
				}

				data, err := pkg.ParseToken(token, "password-hello-word")

				if err != nil {
					return nil, err
				}

				userID, ok := data["user_id"].(string)
				if !ok {
					return nil, fmt.Errorf("无法拿到token信息")
				}

				ctx = context.WithValue(ctx, "user_id", userID)

				if err != nil {
					return nil, err
				}

				return handler(ctx, req)
			}

			return handler(ctx, req)
		}
	}))

	srv := grpc.NewServer(opts...)
	v1.RegisterUsersServer(srv, greeter)
	authV1.RegisterAuthServer(srv, authService)
	return srv
}
