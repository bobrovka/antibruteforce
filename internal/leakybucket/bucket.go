package leakybucket

import (
	"sync"
	"time"
)

const ttlUnused = 15 * time.Second

type bucket struct {
	mu          sync.Mutex
	key         string
	load        int64
	maxload     int64
	unusedSince time.Time
}

func newBucket(key string, maxLoad int64, delCh chan<- string) *bucket {
	b := &bucket{
		key:     key,
		maxload: maxLoad,
	}

	go func() {
		ticker := time.NewTicker(time.Duration(int64(time.Minute) / b.maxload))
		defer ticker.Stop()
		for {
			<-ticker.C

			b.mu.Lock()
			if 0 < b.load {
				b.load--
				b.mu.Unlock()
				continue
			}
			b.mu.Unlock()

			if b.unusedSince.IsZero() {
				b.unusedSince = time.Now()
			} else {
				if ttlUnused < time.Now().Sub(b.unusedSince) {
					delCh <- b.key
					break
				}
			}
		}
	}()

	return b
}

func (b *bucket) try() error {
	b.mu.Lock()
	defer b.mu.Unlock()

	if b.load < b.maxload {
		b.load++
		return nil
	}

	return ErrBlocked
}
