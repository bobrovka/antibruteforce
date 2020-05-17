package leakybucket

import (
	"sync/atomic"
	"time"
)

const ttlUnused = 15 * time.Second

type bucket struct {
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

			if 0 < atomic.LoadInt64(&b.load) {
				atomic.AddInt64(&b.load, -1)
			} else {
				if b.unusedSince.IsZero() {
					b.unusedSince = time.Now()
				} else {
					if ttlUnused < time.Now().Sub(b.unusedSince) {
						delCh <- b.key
						break
					}
				}
			}
		}
	}()

	return b
}

func (b *bucket) try() error {
	if atomic.LoadInt64(&b.load) < b.maxload {
		atomic.AddInt64(&b.load, 1)
		return nil
	}

	return ErrBlocked
}
