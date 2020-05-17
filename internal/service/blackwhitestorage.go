package service

import "net"

type IPStatus int8

const (
	Undefined IPStatus = iota
	Permitted
	Rejected
)

type BlackWhiteStorage interface {
	AddBlackNet(netAddr string) error
	RemoveBlackNet(netAddr string) error
	AddWhiteNet(netAddr string) error
	RemoveWhiteNet(netAddr string) error
	Check(ip net.IP) (IPStatus, error)
	Clear() error
}
