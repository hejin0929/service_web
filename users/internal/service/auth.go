package service

import (
	"context"
	"users/internal/biz"

	pb "users/api/auth/v1"
)

type AuthService struct {
	pb.UnimplementedAuthServer
	repo *biz.AuthUse
}

func NewAuthService(repo *biz.AuthUse) *AuthService {
	return &AuthService{
		repo: repo,
	}
}

func (s *AuthService) LoginUsers(ctx context.Context, req *pb.LoginUsersRequest) (*pb.LoginUsersReply, error) {
	return s.repo.LoginUsers(ctx, req)
}
func (s *AuthService) ExitUsersLogin(ctx context.Context, req *pb.ExitUsersLoginRequest) (*pb.ExitUsersLoginReply, error) {
	return &pb.ExitUsersLoginReply{}, nil
}
func (s *AuthService) PatchUsersLogin(ctx context.Context, req *pb.PatchUsersLoginRequest) (*pb.PatchUsersLoginReply, error) {
	return &pb.PatchUsersLoginReply{}, nil
}
func (s *AuthService) AuthLogin(ctx context.Context, req *pb.AuthLoginRequest) (*pb.AuthLoginReply, error) {
	return s.repo.AuthLogin(ctx, req)
}
