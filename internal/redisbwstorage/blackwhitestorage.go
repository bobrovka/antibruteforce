package blackwhitestorage

import (
	"net"

	"github.com/bobrovka/antibruteforce/internal/service"
	"github.com/go-redis/redis"
)

type Blackwhitestorage struct {
	redisCli *redis.Client
}

func NewBlackwhitestorage(redisClient *redis.Client) *Blackwhitestorage {
	return &Blackwhitestorage{
		redisCli: redisClient,
	}
}

func (s *Blackwhitestorage) AddBlackNet(netAddr string) error {
	err := s.redisCli.SAdd("black", netAddr).Err()
	if err != nil {
		return err
	}

	return nil
}

func (s *Blackwhitestorage) RemoveBlackNet(netAddr string) error {
	err := s.redisCli.SRem("black", netAddr).Err()
	if err != nil {
		return err
	}

	return nil
}

func (s *Blackwhitestorage) AddWhiteNet(netAddr string) error {
	err := s.redisCli.SAdd("white", netAddr).Err()
	if err != nil {
		return err
	}

	return nil
}

func (s *Blackwhitestorage) RemoveWhiteNet(netAddr string) error {
	err := s.redisCli.SRem("white", netAddr).Err()
	if err != nil {
		return err
	}

	return nil
}

func (s *Blackwhitestorage) Clear() error {
	err := s.redisCli.FlushAll().Err()
	return err
}

func (s *Blackwhitestorage) Check(ip net.IP) (service.IpStatus, error) {
	blackList, err := s.redisCli.SMembers("black").Result()
	if err != nil {
		return service.Undefined, err
	}
	for _, blackNet := range blackList {
		_, ipNet, err := net.ParseCIDR(blackNet)
		if err != nil {
			return service.Undefined, err
		}

		if ipNet.Contains(ip) {
			return service.Rejected, nil
		}
	}

	whiteList, err := s.redisCli.SMembers("white").Result()
	if err != nil {
		return service.Undefined, err
	}
	for _, whiteNet := range whiteList {
		_, ipNet, err := net.ParseCIDR(whiteNet)
		if err != nil {
			return service.Undefined, err
		}

		if ipNet.Contains(ip) {
			return service.Permitted, nil
		}
	}

	return service.Undefined, nil
}
