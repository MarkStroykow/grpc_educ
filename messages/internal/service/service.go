package service

import (
	"context"
	"messages/internal/proto"
)

type GrpcService struct {
	proto.UnimplementedMessageServiceServer
}

func NewService() GrpcService {
	return GrpcService{}
}
func (s *GrpcService) SendMessage(_ context.Context, in *proto.CreateMessage) (*proto.MessageResponce, error) {
	return &proto.MessageResponce{
		Ok: true}, nil
}
