package server

import (
	v1 "base/api/district/v1"
	roles "base/api/roles/v1"
	"base/internal/conf"
	"base/internal/service"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Server, greeter *service.DistrictService, logger log.Logger, rolesService *service.RolesService) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
		),
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}

	// 暂时取消认真，开发阶段
	//opts = append(opts, http.Middleware(func(handler middleware.Handler) middleware.Handler {
	//	return func(ctx context.Context, req interface{}) (interface{}, error) {
	//		if cc, ok := transport.FromServerContext(ctx); ok {
	//			token := cc.RequestHeader().Get("Authorization")
	//
	//			conn, err := grpc.Dial("localhost:9001", grpc.WithInsecure())
	//
	//			ctx2 := metadata.AppendToOutgoingContext(context.Background(), "Authorization", token)
	//
	//			if err != nil {
	//				log.Fatalf("failed to connect: %v", err)
	//			}
	//
	//			defer conn.Close()
	//
	//			client := pb.NewAuthClient(conn)
	//
	//			req2 := pb.AuthLoginRequest{Token: token}
	//
	//			_, err = client.AuthLogin(ctx2, &req2)
	//			if err != nil {
	//				return nil, fmt.Errorf("failed to call YourMethod: %v", err)
	//			}
	//
	//			if token == "" {
	//				return nil, errors.New("error：token deletion")
	//			}
	//
	//			data, err := pkg.ParseToken(token, "password-hello-word")
	//
	//			if err != nil {
	//				return nil, err
	//			}
	//
	//			userID, ok := data["user_id"].(string)
	//			if !ok {
	//				return nil, fmt.Errorf("无法拿到token信息")
	//			}
	//
	//			ctx = context.WithValue(ctx, "user_id", userID)
	//
	//			if err != nil {
	//				return nil, err
	//			}
	//
	//			return handler(ctx, req)
	//		}
	//
	//		return handler(ctx, req)
	//	}
	//}))

	srv := http.NewServer(opts...)
	v1.RegisterDistrictHTTPServer(srv, greeter)
	roles.RegisterRolesHTTPServer(srv, rolesService)
	return srv
}
