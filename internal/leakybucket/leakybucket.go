package leakybucket

import (
	"errors"
	"sync"

	"golang.org/x/sync/errgroup"
)

const (
	maxLoadLogin    = 10
	maxLoadPassword = 100
	maxLoadIP       = 1000
)

var ErrBlocked = errors.New("blocked")

type LeakyBucket struct {
	lgmu      sync.Mutex
	lgbuckets map[string]*bucket
	lgdelCh   chan string

	pwmu      sync.Mutex
	pwbuckets map[string]*bucket
	pwdelCh   chan string

	ipmu      sync.Mutex
	ipbuckets map[string]*bucket
	ipdelCh   chan string
}

func NewLeakyBucket() *LeakyBucket {
	lb := &LeakyBucket{
		lgbuckets: map[string]*bucket{},
		lgdelCh:   make(chan string),

		pwbuckets: map[string]*bucket{},
		pwdelCh:   make(chan string),

		ipbuckets: map[string]*bucket{},
		ipdelCh:   make(chan string),
	}

	go func() {
		for {
			select {
			case login := <-lb.lgdelCh:
				lb.lgmu.Lock()
				delete(lb.lgbuckets, login)
				lb.lgmu.Unlock()
			case password := <-lb.pwdelCh:
				lb.pwmu.Lock()
				delete(lb.pwbuckets, password)
				lb.pwmu.Unlock()
			case ip := <-lb.ipdelCh:
				lb.ipmu.Lock()
				delete(lb.ipbuckets, ip)
				lb.ipmu.Unlock()
			}
		}
	}()

	return lb
}

func (s *LeakyBucket) Try(login string, password string, ip string) error {
	var eg errgroup.Group

	eg.Go(func() error {
		s.lgmu.Lock()
		defer s.lgmu.Unlock()

		return checkBucket(s.lgbuckets, login, maxLoadLogin, s.lgdelCh)
	})

	eg.Go(func() error {
		s.pwmu.Lock()
		defer s.pwmu.Unlock()

		return checkBucket(s.pwbuckets, password, maxLoadPassword, s.pwdelCh)
	})

	eg.Go(func() error {
		s.ipmu.Lock()
		defer s.ipmu.Unlock()

		return checkBucket(s.ipbuckets, ip, maxLoadIP, s.ipdelCh)
	})

	if err := eg.Wait(); err != nil {
		return err
	}

	return nil
}

func checkBucket(buckets map[string]*bucket, bucketKey string, maxload int64, delCh chan<- string) error {
	b, ok := buckets[bucketKey]
	if !ok {
		b = newBucket(bucketKey, maxload, delCh)
		buckets[bucketKey] = b
	}

	if b.load < b.maxload {
		b.load++
		return nil
	}

	return ErrBlocked
}
