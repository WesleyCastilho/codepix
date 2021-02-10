package grpc

import (
	"github.com/WesleyCastilho/codepix/application/grpc/pb"
	"github.com/WesleyCastilho/codepix/application/usecase"
)

type PixGrpcService struct {
	PixUseCase usecase.PixUseCase
	pb.UnimplementedPixServiceServer
}
