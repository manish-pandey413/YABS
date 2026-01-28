package config

import (
	"fmt"
	"log"
	"strings"

	"github.com/go-playground/validator/v10"
	_ "github.com/joho/godotenv/autoload"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/v2"
)

type Config struct {
	Server   ServerConfig   `koanf:"server" validate:"required"`
	Database DatabaseConfig `koanf:"database" validate:"required"`
	Auth     AuthConfig     `koanf:"auth" validate:"required"`
}

type Primary struct {
	Env string `koanf:"env" validate:"required"`
}

type ServerConfig struct {
	Port         string `koanf:"port" validate:"required"`
	ReadTimeout  int    `koanf:"read_timeout" validate:"required"`
	WriteTimeout int    `koanf:"write_timeout" validate:"required"`
	IdleTimeout  int    `koanf:"idle_timeout" validate:"required"`
}

type DatabaseConfig struct {
	Host            string `koanf:"host" validate:"required"`
	Port            int    `koanf:"port" validate:"required"`
	User            string `koanf:"user" validate:"required"`
	Password        string `koanf:"password"`
	Name            string `koanf:"name" validate:"required"`
	SSLMode         string `koanf:"ssl_mode" validate:"required"`
	MaxOpenConns    int    `koanf:"max_open_conns" validate:"required"`
	MaxIdleConns    int    `koanf:"max_idle_conns" validate:"required"`
	ConnMaxLifetime int    `koanf:"conn_max_lifetime" validate:"required"`
	ConnMaxIdleTime int    `koanf:"conn_max_idle_time" validate:"required"`
}

type AuthConfig struct {
	SecretKey string `koanf:"secret_key" validate:"required"`
}

func LoadConfig() (*Config, error) {
	k := koanf.New(".")

	err := k.Load(env.Provider("YABS_", ".", func(s string) string {
		return strings.ToLower(strings.TrimPrefix(s, "YABS_"))
	}), nil)
	if err != nil {
		log.Fatal("could not load initial env variables")
	}

	mainConfig := &Config{}

	err = k.Unmarshal("", mainConfig)
	if err != nil {
		log.Fatal("could not unmarshal main config")
	}

	validate := validator.New()

	err = validate.Struct(mainConfig)
	if err != nil {
		log.Fatal("config validation failed")
	} else {
		fmt.Println("config validation passed")
	}

	return mainConfig, nil
}
