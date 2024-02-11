package grpc

import (
	"item-service/internal/usecase"

	"google.golang.org/grpc"
)

func NewGrpcServer(s *grpc.Server, uc *usecase.UseCase) *grpc.Server {
	return s
}
