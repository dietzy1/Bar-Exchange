package server

import (
	"context"
	"fmt"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/dietzy1/Bar-Exchange/protos/event/v1"
)

// run the generated GRPC gateway server
func (s *server) RunGateway() error {

	//The reverse proxy connects to the GRPC server
	conn, err := grpc.DialContext(
		context.Background(),
		/* "dns:///0.0.0.0:8080", */
		"dns:///0.0.0.0"+s.config.Addr,
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return fmt.Errorf("failed to dial: %v", err)
	}

	//Main mux where options are added in -- no options are needed for now
	gwmux := runtime.NewServeMux()

	//Call function which handles registering services to the gateway
	if err := registerGateway(context.Background(), gwmux, conn); err != nil {
		return err
	}

	gwServer := &http.Server{
		Addr:    s.config.GatewayAddr,
		Handler: corsMiddleware(gwmux),
	}
	//Assign to server struct so we can gracefully shutdown
	s.gwServer = gwServer

	fmt.Println("Starting gateway server on port:", s.config.GatewayAddr)
	if err := gwServer.ListenAndServe(); err != nil {
		return fmt.Errorf("failed to listen and serve: %v", err)
	}

	return nil
}

func registerGateway(ctx context.Context, gwmux *runtime.ServeMux, conn *grpc.ClientConn) error {
	if err := pb.RegisterEventServiceHandler(context.Background(), gwmux, conn); err != nil {
		return fmt.Errorf("failed to register gateway: %v", err)
	}

	return nil
}
