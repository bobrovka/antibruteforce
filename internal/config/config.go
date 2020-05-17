package config

import (
	"context"
	"log"

	"github.com/heetch/confita"
	"github.com/heetch/confita/backend/file"
)

// Config базовый конфиг приложения
type Config struct {
	HTTPListen string `config:"host"` // ip и port на котором должен слушать web-сервер
	Redis      string

	RPMLogin    int64
	RPMPassword int64
	RPMIP       int64
}

func GetConfig(configPath string) *Config {
	if configPath == "" {
		log.Fatal("no config file")
	}

	cfg := &Config{
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
