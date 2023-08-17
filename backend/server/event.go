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
	/* res, err := s.event.StartEvent(ctx, req)
	if err != nil {
		return nil, err
	} */
	return nil, nil

}

func (s *server) StopEvent(ctx context.Context, req *pb.StopEventRequest) (*pb.StopEventResponse, error) {
	//return s.event.StopEvent(ctx, req)
	return nil, nil
}

func (s *server) GetEvent(ctx context.Context, req *pb.GetEventRequest) (*pb.GetEventResponse, error) {
	//return s.event.GetEvent(ctx, req)
	return nil, nil
}
