package main

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/bobrovka/antibruteforce/internal"
	"github.com/bobrovka/antibruteforce/internal/service"
	"github.com/bobrovka/antibruteforce/pkg/antibruteforce/api"
	"github.com/heetch/confita"
	"github.com/heetch/confita/backend/file"
	flag "github.com/spf13/pflag"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configPath string

func init() {
	flag.StringVarP(&configPath, "config", "c", "", "path to config file")
}

func main() {
	flag.Parse()

	cfg := getConfig()

	antibruteforceService := service.NewService()

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

func getConfig() *internal.Config {
	if configPath == "" {
		log.Fatal("no config file")
	}

	cfg := &internal.Config{
		HTTPListen: "127.0.0.1:50051",
	}

	loader := confita.NewLoader(
		file.NewBackend(configPath),
	)

	err := loader.Load(context.Background(), cfg)
	if err != nil {
		log.Fatal("cannot read config", err)
	}

	return cfg
}