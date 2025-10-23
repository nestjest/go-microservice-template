package config

import (
	"flag"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

// Config struct make config file in struct
type Config struct {
	Env            string     `yaml:"env" env-default:"local"` // current env
	StoragePath    string     `yaml:"storage_path" env-required:"true"` // path to db file
	GRPC           GPRCConfig `yaml:"grpc"` 
	MigrationsPath string // path to migrations files
	TokenTTL       time.Duration `yaml:"token_ttl" env-default:"1h"` // ttl of jwt token
}

// GPRCConfig make decomposition grps config on GRPSConfig struct
type GPRCConfig struct {
	Port int `yaml:"port"` // grpc server port
	Timeout time.Duration `yaml:"timeout"` // grpc server timeout
}

// MustLoad load a config from path (only panic no err returns).
func MustLoad() *Config {
	configPath := fetchConfigPath()
	if configPath == "" {
		panic("config path is empty")
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		panic("config file does not exist: " + configPath)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		panic("config path is empty: " + err.Error())
	}

	return &cfg
}

// fetchConfigPath fetches config path from command line.
// Priority: flag > env > default.
// Default value is empty string.
func fetchConfigPath() string {
	var res string

	flag.StringVar(&res, "config", "", "path to config file")
	flag.Parse()

	if res == "" {
		res = os.Getenv("CONFIG_PATH")
	}

	return res
}