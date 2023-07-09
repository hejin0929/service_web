package service

import (
	pb "base/api/roles/v1"
	"base/internal/biz"
	"context"
)

type RolesService struct {
	pb.UnimplementedRolesServer
	rope *biz.RolesUse
}

func NewRolesService(rope *biz.RolesUse) *RolesService {
	return &RolesService{
		rope: rope,
	}
}

func (s *RolesService) CreateRoles(ctx context.Context, req *pb.CreateRolesRequest) (*pb.CreateRolesReply, error) {
	return &pb.CreateRolesReply{}, nil
}
func (s *RolesService) UpdateRoles(ctx context.Context, req *pb.UpdateRolesRequest) (*pb.UpdateRolesReply, error) {
	return &pb.UpdateRolesReply{}, nil
}
func (s *RolesService) DeleteRoles(ctx context.Context, req *pb.DeleteRolesRequest) (*pb.DeleteRolesReply, error) {
	return &pb.DeleteRolesReply{}, nil
}
func (s *RolesService) GetRoles(ctx context.Context, req *pb.GetRolesRequest) (*pb.GetRolesReply, error) {
	return &pb.GetRolesReply{}, nil
}
func (s *RolesService) ListRoles(ctx context.Context, req *pb.ListRolesRequest) (*pb.ListRolesReply, error) {
	return &pb.ListRolesReply{}, nil
}
