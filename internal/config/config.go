package config

import (
	"flag"
	"github.com/glamostoffer/ValinorChat/internal/system/http"
	pg "github.com/glamostoffer/ValinorChat/pkg/pg_connector"
	authclient "github.com/glamostoffer/ValinorProtos/auth"
	"github.com/ilyakaznacheev/cleanenv"
	"os"
	"time"
)

type Config struct {
	Env          string            `yaml:"env"`
	StartTimeout time.Duration     `yaml:"start_timeout"`
	StopTimeout  time.Duration     `yaml:"stop_timeout"`
	GRPC         GRPCConfig        `yaml:"grpc"`
	HTTP         http.Config       `yaml:"http"`
	AuthCfg      authclient.Config `yaml:"auth_cfg"`
	Postgres     pg.Config         `yaml:"postgres"`
}

type GRPCConfig struct {
	Host    string        `yaml:"host"`
	Port    int           `yaml:"port"`
	Timeout time.Duration `yaml:"timeout"`
}

func LoadConfig() *Config {
	configPath := os.Getenv("CONFIG_PATH")

	if configPath == "" {
		flag.StringVar(&configPath, "config", "", "path to config file")
		flag.Parse()
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		panic("config file does not exist: " + configPath)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		panic("cannot read config: " + err.Error())
	}

	return &cfg
}
