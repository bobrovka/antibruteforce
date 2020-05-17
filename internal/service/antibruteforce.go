package service

import (
	"context"
	"net"

	"github.com/bobrovka/antibruteforce/pkg/antibruteforce/api"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Service struct {
	bws BlackWhiteStorage
	lb  LeakyBucket
}

func NewService(bws BlackWhiteStorage, lb LeakyBucket) *Service {
	return &Service{
		bws: bws,
		lb:  lb,
	}
}

func (s *Service) AddBlackNet(_ context.Context, request *api.NetAddr) (*empty.Empty, error) {
	_, ipNet, err := net.ParseCIDR(request.GetNetAddr())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	err = s.bws.AddBlackNet(ipNet.String())
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *Service) RemoveBlackNet(_ context.Context, request *api.NetAddr) (*empty.Empty, error) {
	_, ipNet, err := net.ParseCIDR(request.GetNetAddr())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	err = s.bws.RemoveBlackNet(ipNet.String())
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *Service) AddWhiteNet(_ context.Context, request *api.NetAddr) (*empty.Empty, error) {
	_, ipNet, err := net.ParseCIDR(request.GetNetAddr())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	err = s.bws.AddWhiteNet(ipNet.String())
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *Service) RemoveWhiteNet(_ context.Context, request *api.NetAddr) (*empty.Empty, error) {
	_, ipNet, err := net.ParseCIDR(request.GetNetAddr())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	err = s.bws.RemoveWhiteNet(ipNet.String())
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *Service) ResetBlackWhiteLists(_ context.Context, _ *empty.Empty) (*empty.Empty, error) {
	err := s.bws.Clear()
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &emptypb.Empty{}, nil
}

func (s *Service) Check(_ context.Context, request *api.CheckRequest) (*api.Status, error) {
	ip := net.ParseIP(request.GetIp())
	if len(ip) == 0 {
		return nil, status.Error(codes.InvalidArgument, "field ip is required and should be appropriate format")
	}

	statusIP, err := s.bws.Check(ip)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	switch statusIP {
	case Rejected:
		return &api.Status{
			Ok: false,
		}, nil
	case Permitted:
		return &api.Status{
			Ok: true,
		}, nil
	}

	err = s.lb.Try(request.GetLogin(), request.GetPassword(), request.GetIp())
	if err == errRejected {
		return &api.Status{
			Ok: false,
		}, nil
	} else if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &api.Status{
		Ok: true,
	}, nil
}
