package server

import (
	"context"
	"net/http"
	"os"
	"regexp"

	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func loggingMiddleware(
	logger *zap.Logger,
) grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		logger.Info("gRPC method", zap.String("method", info.FullMethod))

		resp, err := handler(ctx, req)

		// You can log the response and error here if needed.
		// For example:
		if err != nil {
			logger.Error("gRPC method encountered an error", zap.Error(err))
		} else {
			logger.Info("gRPC method completed successfully")
		}

		return resp, err
	}
}

// CORS middleware wrapper that allows origins -- configured in ENV
func corsMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if allowedOrigin(r.Header.Get("Origin")) {
			w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, DELETE")
			w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization, ResponseType, Origin")
			w.Header().Set("Access-Control-Allow-Credentials", "true")
		}
		if r.Method == "OPTIONS" {
			return
		}
		h.ServeHTTP(w, r)

	})
}

// Reads ENV file and determines if origin should be * or regex matching
func allowedOrigin(origin string) bool {
	if os.Getenv("CORS") == "*" {

		return true
	}
	if matched, _ := regexp.MatchString(os.Getenv(("CORS")), origin); matched {
		return true
	}
	return false
}
