package main

import (
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/bobrovka/antibruteforce/internal/config"
	"github.com/bobrovka/antibruteforce/internal/leakybucket"
	blackwhitestorage "github.com/bobrovka/antibruteforce/internal/redisbwstorage"
	"github.com/bobrovka/antibruteforce/internal/service"
	"github.com/bobrovka/antibruteforce/pkg/antibruteforce/api"
	"github.com/go-redis/redis"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	cfg := config.GetConfig(os.Getenv("CONFIG_PATH"))

	redisClient := getRedisClient(cfg.Redis)

	bws := blackwhitestorage.NewBlackwhitestorage(redisClient)
	lb := leakybucket.NewLeakyBucket(cfg.RPMLogin, cfg.RPMPassword, cfg.RPMIP)

	antibruteforceService := service.NewService(bws, lb)

	// Create grpc server
	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)

	api.RegisterAntibruteforceServer(grpcServer, antibruteforceService)
	lis, err := net.Listen("tcp", cfg.HTTPListen)
	if err != nil {
		log.Fatal("cannot listen on ", cfg.HTTPListen)
	}

	shutdown := make(chan error, 1)
	go func() {
		err := grpcServer.Serve(lis)
		if err != nil {
			shutdown <- err
		}
	}()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case x := <-interrupt:
		log.Printf("Received `%v`.", x)
	case err := <-shutdown:
		log.Printf("Received shutdown message: %v", err)
	}

	grpcServer.GracefulStop()
}

func getRedisClient(addr string) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "",
		DB:       0,
	})

	_, err := rdb.Ping().Result()
	if err != nil {
		log.Fatal("cannot connect to redis, ", err)
	}

	return rdb
}
