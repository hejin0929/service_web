package service

import (
	"base/internal/biz"
	"context"

	pb "base/api/district/v1"
)

type DistrictService struct {
	pb.UnimplementedDistrictServer
	repo *biz.DistrictUse
}

func NewDistrictService(repo *biz.DistrictUse) *DistrictService {
	return &DistrictService{
		repo: repo,
	}
}

func (s *DistrictService) CreateDistrict(ctx context.Context, req *pb.CreateDistrictRequest) (*pb.CreateDistrictReply, error) {
	return &pb.CreateDistrictReply{}, nil
}
func (s *DistrictService) UpdateDistrict(ctx context.Context, req *pb.UpdateDistrictRequest) (*pb.UpdateDistrictReply, error) {
	return s.repo.SetDistrict(ctx, req)
}
func (s *DistrictService) DeleteDistrict(ctx context.Context, req *pb.DeleteDistrictRequest) (*pb.DeleteDistrictReply, error) {
	return &pb.DeleteDistrictReply{}, nil
}
func (s *DistrictService) GetDistrict(ctx context.Context, req *pb.GetDistrictRequest) (*pb.GetDistrictReply, error) {
	return s.repo.GetDistrict(ctx, req)
}
