package server

import (
	"context"

	pb "github.com/dietzy1/Bar-Exchange/protos/event/v1"
	"github.com/dietzy1/Bar-Exchange/service"
)

type event interface {
	StartEvent(ctx context.Context, req service.Event) (service.Event, error)
	StopEvent(ctx context.Context, req service.Event) (service.Event, error)
	GetEvent(ctx context.Context, req service.Event) (service.Event, error)
}

func (s *server) StartEvent(ctx context.Context, req *pb.StartEventRequest) (*pb.StartEventResponse, error) {

	input := service.Event{
		FutureTimeStamp: req.FutureTimestamp,
	}

	res, err := s.event.StartEvent(ctx, input)
	if err != nil {
		return nil, err
	}

	return &pb.StartEventResponse{
		Id:       res.Id,
		Duration: 30,
	}, nil

}

func (s *server) StopEvent(ctx context.Context, req *pb.StopEventRequest) (*pb.StopEventResponse, error) {

	input := service.Event{
		Id: req.Id,
	}

	_, err := s.event.StopEvent(ctx, input)
	if err != nil {
		return nil, err
	}

	return &pb.StopEventResponse{}, nil

}

func (s *server) GetEvent(ctx context.Context, req *pb.GetEventRequest) (*pb.GetEventResponse, error) {
	//return s.event.GetEvent(ctx, req)
	input := service.Event{
		Id: req.Id,
	}

	res, err := s.event.GetEvent(ctx, input)
	if err != nil {
		return nil, err
	}

	return &pb.GetEventResponse{
		Id:              res.Id,
		FutureTimestamp: res.FutureTimeStamp,
	}, nil

}
