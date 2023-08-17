package server

import (
	"io"
	"net"
	"os"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"

	pb "github.com/dietzy1/Bar-Exchange/protos/event/v1"
)

type server struct {
	pb.UnimplementedEventServiceServer
	grpc *grpc.Server

	logger *zap.Logger
	config *Config

	//Domain interface
	event event
}

type Config struct {
	Addr   string
	Logger *zap.Logger
}

func New(c *Config, event event) *server {

	if c.Addr == "" {
		c.Addr = ":9000"
		c.Logger.Info("No port specified, defaulting to 9000")
	}

	//Unsure if this is even supposed to be here honestly
	log := grpclog.NewLoggerV2(os.Stdout, io.Discard, io.Discard)
	grpclog.SetLoggerV2(log)

	//Create new GRPC server object
	grpc := grpc.NewServer(
		grpc.UnaryInterceptor(loggingMiddleware(c.Logger)),
	)

	s := &server{
		grpc:   grpc,
		logger: c.Logger,
		config: c,
		tester: tester,
	}

	//Register the server object methods with the GRPC server
	pb.RegisterTesterServiceServer(grpc, s)

	return s
}

func (s *server) ListenAndServe() error {

	lis, err := net.Listen("tcp", s.config.Addr)
	if err != nil {
		s.logger.Error("Failed to listen:", zap.Error(err))
		return err
	}

	s.logger.Info("Serving gRPC on http://", zap.String("addr", s.config.Addr))

	if err := s.grpc.Serve(lis); err != nil {
		s.logger.Error("Failed to serve:", zap.Error(err))
		return err
	}

	//Here I want to call the gateway server

	return nil
}

func (s *server) Stop() {
	s.grpc.GracefulStop()
	s.logger.Info("gRPC server stopped gracefully")
}