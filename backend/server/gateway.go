package server

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/dietzy1/Bar-Exchange/protos/event/v1"
)

// run the generated GRPC gateway server
func runGateway(logger *zap.Logger) error {

	//The reverse proxy connects to the GRPC server
	conn, err := grpc.DialContext(
		context.Background(),
		/* "dns:///0.0.0.0:8080", */
		"dns:///0.0.0.0"+os.Getenv("GRPC"),
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return fmt.Errorf("failed to dial: %v", err)
	}

	//Main mux where options are added in
	gwmux := runtime.NewServeMux()

	//middleware chaining

	middleware := corsMiddleware(gwmux)
	loggingMiddleware(middleware, logger)

	//middleware := logger(cors(gwmux))

	gwServer := &http.Server{
		Addr:    gatewayAddress,
		Handler: middleware,
	}

	log.Info("Serving gRPC-Gateway", gatewayAddress)
	log.Fatalln(gwServer.ListenAndServe())

	return nil
}

func registerGateway(ctx context.Context, gwmux *runtime.ServeMux, conn *grpc.ClientConn) error {
	if err := pb.RegisterEventServiceHandler(context.Background(), gwmux, conn); err != nil {
		return fmt.Errorf("failed to register gateway: %v", err)
	}

	return nil
}
