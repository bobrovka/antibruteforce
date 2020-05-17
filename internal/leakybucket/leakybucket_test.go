package leakybucket

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestBucket(t *testing.T) {
	t.Run("overflow_login", func(t *testing.T) {
		lb := NewLeakyBucket(10, 100, 100)
		for i := int64(0); i < 2*lb.maxLoadLogin; i++ {
			err := lb.Try("nagibator228", "123456", "127.0.0.1")
			if i < lb.maxLoadLogin {
				assert.NoError(t, err)
			} else {
				assert.Equal(t, ErrBlocked, err)
			}
		}
	})
	t.Run("overflow_password", func(t *testing.T) {
		lb := NewLeakyBucket(100, 10, 100)
		for i := int64(0); i < 2*lb.maxLoadPassword; i++ {
			err := lb.Try("nagibator228", "123456", "127.0.0.1")
			if i < lb.maxLoadPassword {
				assert.NoError(t, err)
			} else {
				assert.Equal(t, ErrBlocked, err)
			}
		}
	})
	t.Run("overflow_ip", func(t *testing.T) {
		lb := NewLeakyBucket(100, 100, 10)
		for i := int64(0); i < 2*lb.maxLoadIP; i++ {
			err := lb.Try("nagibator228", "123456", "127.0.0.1")
			if i < lb.maxLoadIP {
				assert.NoError(t, err)
			} else {
				assert.Equal(t, ErrBlocked, err)
			}
		}
	})
	t.Run("no_block", func(t *testing.T) {
		lb := NewLeakyBucket(100, 100, 100)
		for i := int64(0); i < 50; i++ {
			err := lb.Try("nagibator228", "123456", "127.0.0.1")
			if i < lb.maxLoadLogin && i < lb.maxLoadPassword && i < lb.maxLoadIP {
				assert.NoError(t, err)
			} else {
				assert.Equal(t, ErrBlocked, err)
			}
		}
	})
	t.Run("check_bandwidth_login", func(t *testing.T) {
		t.Parallel()
		// 60 rpm = 1 rps
		lb := NewLeakyBucket(60, 1000, 1000)

		// fill up the login bucket a little
		for i := int64(0); i < 10; i++ {
			err := lb.Try("nagibator228", "123456", "127.0.0.1")
			assert.NoError(t, err)
		}

		// wait till some requests flow away
		time.Sleep(5*time.Second + time.Second/100)

		// check bucket filling
		for i := int64(0); i < 100; i++ {
			err := lb.Try("nagibator228", "123456", "127.0.0.1")
			if i < 55 {
				assert.NoError(t, err)
			} else {
				assert.Equal(t, ErrBlocked, err)
			}
		}
	})
	t.Run("check_bandwidth_password", func(t *testing.T) {
		t.Parallel()
		// 60 rpm = 1 rps
		lb := NewLeakyBucket(1000, 60, 1000)

		// fill up the password bucket a little
		for i := int64(0); i < 10; i++ {
			err := lb.Try("nagibator228", "123456", "127.0.0.1")
			assert.NoError(t, err)
		}

		// wait till some requests flow away
		time.Sleep(5*time.Second + time.Second/100)

		// check bucket filling
		for i := int64(0); i < 100; i++ {
			err := lb.Try("nagibator228", "123456", "127.0.0.1")
			if i < 55 {
				assert.NoError(t, err)
			} else {
				assert.Equal(t, ErrBlocked, err)
			}
		}
	})
	t.Run("check_bandwidth_ip", func(t *testing.T) {
		t.Parallel()
		// 60 rpm = 1 rps
		lb := NewLeakyBucket(1000, 1000, 60)

		// fill up the ip bucket a little
		for i := int64(0); i < 10; i++ {
			err := lb.Try("nagibator228", "123456", "127.0.0.1")
			assert.NoError(t, err)
		}

		// wait till some requests flow away
		time.Sleep(5*time.Second + time.Second/100)

		// check bucket filling
		for i := int64(0); i < 100; i++ {
			err := lb.Try("nagibator228", "123456", "127.0.0.1")
			if i < 55 {
				assert.NoError(t, err)
			} else {
				assert.Equal(t, ErrBlocked, err)
			}
		}
	})
}
