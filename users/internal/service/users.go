package service

import (
	"context"
	"users/internal/biz"

	pb "users/api/v1"
)

type UsersService struct {
	pb.UnimplementedUsersServer
	biz *biz.UsersUse
}

func NewUsersService(repo *biz.UsersUse) *UsersService {
	return &UsersService{
		biz: repo,
	}
}

func (s *UsersService) CreateUsers(ctx context.Context, req *pb.CreateUsersRequest) (*pb.CreateUsersReply, error) {
	return s.biz.CreateUsers(ctx, req)
}
func (s *UsersService) UpdateUsers(ctx context.Context, req *pb.UpdateUsersRequest) (*pb.UpdateUsersReply, error) {
	return &pb.UpdateUsersReply{}, nil
}
func (s *UsersService) DeleteUsers(ctx context.Context, req *pb.DeleteUsersRequest) (*pb.DeleteUsersReply, error) {
	return &pb.DeleteUsersReply{}, nil
}
func (s *UsersService) GetUsers(ctx context.Context, req *pb.GetUsersRequest) (*pb.GetUsersReply, error) {
	return &pb.GetUsersReply{}, nil
}
func (s *UsersService) ListUsers(ctx context.Context, req *pb.ListUsersRequest) (*pb.ListUsersReply, error) {
	return &pb.ListUsersReply{}, nil
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
func (s *UsersService) PatchPassword(ctx context.Context, req *pb.PatchPasswordRequest) (*pb.PatchPasswordReply, error) {
	return &pb.PatchPasswordReply{}, nil
}
