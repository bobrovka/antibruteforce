// +build integration

package blackwhitestorage

import (
	"net"
	"os"
	"testing"

	"github.com/bobrovka/antibruteforce/internal/config"
	"github.com/bobrovka/antibruteforce/internal/service"
	"github.com/go-redis/redis"
	"github.com/stretchr/testify/assert"
)

func TestBucket(t *testing.T) {
	cfg := config.GetConfig(os.Getenv("CONFIG_PATH"))

	redisClient := redis.NewClient(&redis.Options{
		Addr:     cfg.Redis,
		Password: "",
		DB:       0,
	})
	bws := NewBlackwhitestorage(redisClient)

	t.Run("check_connection", func(t *testing.T) {
		pong, err := redisClient.Ping().Result()
		assert.NoError(t, err)
		assert.Equal(t, "PONG", pong)
	})

	t.Run("check_whitelist", func(t *testing.T) {
		bws.Clear()
		defer bws.Clear()

		ip := net.ParseIP("192.168.0.1")
		status, err := bws.Check(ip)
		assert.NoError(t, err)
		assert.Equal(t, service.Undefined, status)

		err = bws.AddWhiteNet("192.168.0.0/24")
		assert.NoError(t, err)

		ip = net.ParseIP("192.168.0.1")
		status, err = bws.Check(ip)
		assert.NoError(t, err)
		assert.Equal(t, service.Permitted, status)

		err = bws.RemoveWhiteNet("192.168.0.0/24")
		assert.NoError(t, err)

		ip = net.ParseIP("192.168.0.1")
		status, err = bws.Check(ip)
		assert.NoError(t, err)
		assert.Equal(t, service.Undefined, status)
	})

	t.Run("check_blacklist", func(t *testing.T) {
		bws.Clear()
		defer bws.Clear()

		ip := net.ParseIP("192.168.0.1")
		status, err := bws.Check(ip)
		assert.NoError(t, err)
		assert.Equal(t, service.Undefined, status)

		err = bws.AddBlackNet("192.168.0.0/24")
		assert.NoError(t, err)

		ip = net.ParseIP("192.168.0.1")
		status, err = bws.Check(ip)
		assert.NoError(t, err)
		assert.Equal(t, service.Rejected, status)

		err = bws.RemoveBlackNet("192.168.0.0/24")
		assert.NoError(t, err)

		ip = net.ParseIP("192.168.0.1")
		status, err = bws.Check(ip)
		assert.NoError(t, err)
		assert.Equal(t, service.Undefined, status)
	})
}
