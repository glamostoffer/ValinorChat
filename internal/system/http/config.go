package http

import "time"

type Config struct {
	StopTimeout   time.Duration `yaml:"stop_timeout" validate:"required"`
	Host          string        `yaml:"host" validate:"required"`
	Port          string        `yaml:"port" validate:"required"`
	AllowOrigins  string        `yaml:"allow_origins" validate:"required"`
	AllowHeaders  string        `yaml:"allow_headers" validate:"required"`
	ExposeHeaders string        `yaml:"expose_headers" validate:"required"`
	PProfEnabled  bool
}
