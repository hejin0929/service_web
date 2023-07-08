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
	"github.com/go-kratos/kratos/v2/transport/http"
	"os"
	v1 "users/api/users/v1"
	"users/internal/conf"
	"users/internal/service"
	"users/pkg"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Server, greeter *service.UsersService, logger log.Logger) *http.Server {
	var (
		opts = []http.ServerOption{
			http.Middleware(
				recovery.Recovery(),
				validate.Validator(),
			),
		}
	)
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}

	opts = append(opts, http.Middleware(func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			if cc, ok := transport.FromServerContext(ctx); ok {
				if cc.Operation() == "/api.v1.Users/LoginUsers" || cc.Operation() == "/api.v1.Users/CreateUsers" {
					return handler(ctx, req)
				}
				session := cc.RequestHeader().Get("session")

				fmt.Println("this is a value ?? ", session)

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

	srv := http.NewServer(opts...)

	srv.Route("/doc").GET("/doc.json", func(c http.Context) error {
		dir, _ := os.Getwd()

		file, err := os.ReadFile(dir + "/doc.json")

		if err != nil {
			c.JSON(404, err)
		}
		c.Response().Write(file)

		return err
	})

	v1.RegisterUsersHTTPServer(srv, greeter)
	return srv
}
