package service

import "errors"

var errRejected = errors.New("bad auth request")

type LeakyBucket interface {
	Try(login, password, ip string) error
}
