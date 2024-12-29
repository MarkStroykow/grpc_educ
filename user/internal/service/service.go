package service

import (
	"context"
	"strconv"
	"users/internal/proto"
	"users/internal/storage"
)

type GrpcService struct {
	proto.UnimplementedUserServiceServer
	userStorage storage.Storage
}

func NewService() GrpcService {
	return GrpcService{
		userStorage: storage.NewStorage(),
	}
}
func (s *GrpcService) GetUser(_ context.Context, in *proto.GetUserRequest) (*proto.UserResponse, error) {
	id, _ := strconv.Atoi(in.UserId)
	user := s.userStorage.GetUser(uint(id))
	s.userStorage.GetUser(uint(id))
	return &proto.UserResponse{
		UserId: in.UserId,
		Name:   user.Name,
		Email:  user.Email}, nil
}

func (s *GrpcService) CreateUser(_ context.Context, in *proto.CreateUserRequest) (*proto.UserResponse, error) {
	id := s.userStorage.CreateUser(in.Name, in.Email)
	return &proto.UserResponse{
		UserId: id,
		Name:   in.Name,
		Email:  in.Email}, nil
}
