package config

import (
	"log"
	"os"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/v2"
)

type Config struct {
	Server   ServerConfig   `koanf:"server" validate:"required"`
	Database DatabaseConfig `koanf:"database" validate:"required"`
	Auth     AuthConfig     `koanf:"auth" validate:"required"`
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

func parseMapString(value string) (map[string]string, bool) {
	if !strings.HasPrefix(value, "map[") || !strings.HasSuffix(value, "]") {
		return nil, false
	}

	content := strings.TrimPrefix(value, "map[")
	content = strings.TrimSuffix(content, "]")

	if content == "" {
		return make(map[string]string), true
	}

	result := make(map[string]string)

	i := 0
	for i < len(content) {
		keyStart := i
		for i < len(content) && content[i] != ':' {
			i++
		}
		if i >= len(content) {
			break
		}

		key := strings.TrimSpace(content[keyStart:i])
		i++

		valueStart := i
		if i+4 <= len(content) && content[i:i+4] == "map[" {
			bracketCount := 0
			for i < len(content) {
				if i+4 <= len(content) && content[i:i+4] == "map[" {
					bracketCount++
					i += 4
				} else if content[i] == ']' {
					bracketCount--
					i++
					if bracketCount == 0 {
						break
					}
				} else {
					i++
				}
			}
		} else {
			for i < len(content) && content[i] != ' ' {
				i++
			}
		}

		value := strings.TrimSpace(content[valueStart:i])

		if nestedMap, isNested := parseMapString(value); isNested {
			for nestedKey, nestedValue := range nestedMap {
				result[key+"."+nestedKey] = nestedValue
			}
		} else {
			result[key] = value
		}

		for i < len(content) && content[i] == ' ' {
			i++
		}
	}

	return result, true
}

func LoadConfig() (*Config, error) {
	k := koanf.New(".")

	envVars := make(map[string]string)
	for _, env := range os.Environ() {
		parts := strings.SplitN(env, "=", 2)
		if len(parts) == 2 && strings.HasPrefix(parts[0], "YABS_") {
			key := parts[0]
			value := parts[1]

			configKey := strings.ToLower(strings.TrimPrefix(key, "YABS_"))

			if mapData, isMap := parseMapString(value); isMap {
				for mapKey, mapValue := range mapData {
					flatKey := configKey + "." + strings.ToLower(mapKey)
					envVars[flatKey] = mapValue
				}
			} else {
				envVars[configKey] = value
			}
		}
	}

	err := k.Load(env.ProviderWithValue("YABS_", ".", func(key, value string) (string, any) {
		return strings.ToLower(strings.TrimPrefix(key, "YABS_")), value
	}), nil)
	if err != nil {
		log.Fatal("could not load initial env variables")
	}

	for key, value := range envVars {
		k.Set(key, value)
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
		log.Println("config validation passed")
	}

	return mainConfig, nil
}
