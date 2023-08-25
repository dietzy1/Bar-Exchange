package server

import (
	"context"

	pb "github.com/dietzy1/Bar-Exchange/protos/beverage/v1"
	"github.com/dietzy1/Bar-Exchange/service"
)

type beverage interface {
	GetBeverages(ctx context.Context) ([]service.Beverage, error)
	CreateBeverage(ctx context.Context, req service.Beverage) error
	UpdateBeverage(ctx context.Context, req service.Beverage) error
	DeleteBeverage(ctx context.Context, req service.Beverage) error
}

func (s *server) GetBeverages(ctx context.Context, req *pb.GetBeveragesRequest) (*pb.GetBeveragesResponse, error) {

	res, err := s.beverage.GetBeverages(ctx)
	if err != nil {
		return nil, err
	}

	//convert from []service.Beverage to []*pb.Beverage
	beverages := make([]*pb.Beverage, len(res))
	for i, v := range res {
		beverages[i] = &pb.Beverage{
			Id:               v.Id,
			Name:             v.Name,
			Price:            v.Price,
			PercentageChange: int64(v.PercentageChange),
			Type:             pb.BeverageType(v.Type),
			Status:           pb.Status(v.Status),
		}
	}

	return &pb.GetBeveragesResponse{
		Beverages: beverages,
	}, nil

}

func (s *server) CreateBeverage(ctx context.Context, req *pb.CreateBeverageRequest) (*pb.CreateBeverageResponse, error) {

	input := service.Beverage{
		Name:  req.Beverage.Name,
		Price: req.Beverage.Price,
		Type:  service.BeverageType(req.Beverage.Type),
	}

	if err := s.beverage.CreateBeverage(ctx, input); err != nil {
		return nil, err
	}

	return &pb.CreateBeverageResponse{}, nil

}

func (s *server) UpdateBeverage(ctx context.Context, req *pb.UpdateBeverageRequest) (*pb.UpdateBeverageResponse, error) {

	input := service.Beverage{
		Id:    req.Beverage.Id,
		Name:  req.Beverage.Name,
		Price: req.Beverage.Price,
		Type:  service.BeverageType(req.Beverage.Type),
	}

	if err := s.beverage.UpdateBeverage(ctx, input); err != nil {
		return nil, err
	}

	return &pb.UpdateBeverageResponse{}, nil

}

func (s *server) DeleteBeverage(ctx context.Context, req *pb.DeleteBeverageRequest) (*pb.DeleteBeverageResponse, error) {

	input := service.Beverage{
		Id: req.Id,
	}

	if err := s.beverage.DeleteBeverage(ctx, input); err != nil {
		return nil, err
	}

	return &pb.DeleteBeverageResponse{}, nil

}
