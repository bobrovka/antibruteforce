package leakybucket

import "time"

const ttlUnused = 5 * time.Second // TODO to 15

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

			if 0 < b.load {
				b.load--
			} else {
				if b.unusedSince.IsZero() {
					b.unusedSince = time.Now()
				} else {
					if ttlUnused < time.Now().Sub(b.unusedSince) {
						delCh <- b.key
					}
				}
			}
		}
	}()

	return b
}
