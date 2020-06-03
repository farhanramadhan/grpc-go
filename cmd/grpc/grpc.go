package grpc

import "google.golang.org/grpc"

type grpcServer struct {
	GrpcHandler grpcHandler
	Server      *grpc.Server
}

func NewGrpcServer(grpcHandler grpcHandler) *grpcServer {
	return &grpcServer{
		GrpcHandler: grpcHandler,
		Server:      grpc.NewServer(),
	}
}
