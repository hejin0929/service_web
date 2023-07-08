package service

import (
	"context"

	pb "users/api/auth/v1"
)

type AuthService struct {
	pb.UnimplementedUsersServer
}

func NewAuthService() *UsersService {
	return &UsersService{}
}

func (s *UsersService) LoginUsers(ctx context.Context, req *pb.LoginUsersRequest) (*pb.LoginUsersReply, error) {
	return &pb.LoginUsersReply{}, nil
}
func (s *UsersService) ExitUsersLogin(ctx context.Context, req *pb.ExitUsersLoginRequest) (*pb.ExitUsersLoginReply, error) {
	return &pb.ExitUsersLoginReply{}, nil
}
func (s *UsersService) PatchUsersLogin(ctx context.Context, req *pb.PatchUsersLoginRequest) (*pb.PatchUsersLoginReply, error) {
	return &pb.PatchUsersLoginReply{}, nil
}
func (s *UsersService) AuthLogin(ctx context.Context, req *pb.AuthLoginRequest) (*pb.AuthLoginReply, error) {
	return &pb.AuthLoginReply{}, nil
}
