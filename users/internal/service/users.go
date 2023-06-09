package service

import (
	"context"
	"users/internal/biz"

	pb "users/api/v1"
)

type UsersService struct {
	pb.UnimplementedUsersServer
	coll *biz.UsersUse
}

func NewUsersService(coll *biz.UsersUse) *UsersService {
	return &UsersService{
		coll: coll,
	}
}

func (s *UsersService) CreateUsers(ctx context.Context, req *pb.CreateUsersRequest) (*pb.CreateUsersReply, error) {
	return s.coll.CreateUsers(ctx, req)
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
