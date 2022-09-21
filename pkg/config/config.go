package config

import (
	"context"
	"fmt"

	"github.com/sethvargo/go-envconfig"
)

type Config struct {
	Port int `env:"PORT,required"`
	DB   *databaseConfig
}

type databaseConfig struct {
	DSN              string `env:"MYSQL_DSN,required"`
	MaxOpenConns     int    `env:"MAX_OPEN_CONNS,default=100"`
	MaxIdleConns     int    `env:"MAX_IDLE_CONNS,default=100"`
	ConnsMaxLifetime int    `env:"CONNS_MAX_LIFETIME,default=100"`
}

func LoadConfig(ctx context.Context) (*Config, error) {
	var cfg Config
	if err := envconfig.Process(ctx, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

func (cfg *Config) Address() string {
	return fmt.Sprintf(":%d", cfg.Port)
}
