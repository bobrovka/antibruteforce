package service

import "net"

type IpStatus int8

const (
	Undefined IpStatus = iota
	Permitted
	Rejected
)

type BlackWhiteStorage interface {
	AddBlackNet(netAddr string) error
	RemoveBlackNet(netAddr string) error
	AddWhiteNet(netAddr string) error
	RemoveWhiteNet(netAddr string) error
	Check(ip net.IP) (IpStatus, error)
}
