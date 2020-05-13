package service

import (
	"context"

	"github.com/bobrovka/antibruteforce/pkg/antibruteforce/api"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) AddBlackNet(_ context.Context, request *api.NetAddr) (*empty.Empty, error) {
	return &emptypb.Empty{}, nil
}

func (s *Service) RemoveBlackNet(_ context.Context, _ *api.NetAddr) (*empty.Empty, error) {
	return &emptypb.Empty{}, nil
}

func (s *Service) AddWhiteNet(_ context.Context, _ *api.NetAddr) (*empty.Empty, error) {
	return &emptypb.Empty{}, nil
}

func (s *Service) RemoveWhiteNet(_ context.Context, _ *api.NetAddr) (*empty.Empty, error) {
	return &emptypb.Empty{}, nil
}
