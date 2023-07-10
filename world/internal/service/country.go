package service

import (
	"context"
	"world/internal/biz"

	pb "world/api/country/v1"
)

type CountryService struct {
	pb.UnimplementedCountryServer
	repo *biz.CountryUse
}

func NewCountryService(repo *biz.CountryUse) *CountryService {
	return &CountryService{
		repo: repo,
	}
}

func (s *CountryService) CreateCountry(ctx context.Context, req *pb.CreateCountryRequest) (*pb.CreateCountryReply, error) {
	return s.repo.CreateCountry(ctx, req)
}
func (s *CountryService) UpdateCountry(ctx context.Context, req *pb.UpdateCountryRequest) (*pb.UpdateCountryReply, error) {
	return &pb.UpdateCountryReply{}, nil
}
func (s *CountryService) DeleteCountry(ctx context.Context, req *pb.DeleteCountryRequest) (*pb.DeleteCountryReply, error) {
	return &pb.DeleteCountryReply{}, nil
}
func (s *CountryService) GetCountry(ctx context.Context, req *pb.GetCountryRequest) (*pb.GetCountryReply, error) {
	return &pb.GetCountryReply{}, nil
}
func (s *CountryService) ListCountry(ctx context.Context, req *pb.ListCountryRequest) (*pb.ListCountryReply, error) {
	return s.repo.GetCountryLists(ctx, req)
}
