package config

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Config struct {
	App   AppConfig
	DB    DBConfig
	Redis RedisConfig
}

func LoadAllConfigs(dir string) (*Config, error) {
	_ = godotenv.Load(filepath.Join(dir, ".env"))

	cfg := &Config{}
	load := []struct {
		Name   string
		Target interface{}
		EnvFn  func(*viper.Viper)
	}{
		{"app", &cfg.App, nil},
		{"database", &cfg.DB, func(v *viper.Viper) { overrideEnv(v, "database") }},
		{"redis", &cfg.Redis, func(v *viper.Viper) { overrideEnv(v, "redis") }},
	}

	for _, item := range load {
		v := viper.New()
		v.SetConfigName(item.Name)
		v.SetConfigType("yaml")
		v.AddConfigPath(dir)
		v.AutomaticEnv()
		v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

		if err := v.ReadInConfig(); err != nil {
			return nil, fmt.Errorf("error loading %s.yaml: %w", item.Name, err)
		}
		if err := v.Unmarshal(item.Target); err != nil {
			return nil, fmt.Errorf("error decoding %s config: %w", item.Name, err)
		}
		if item.EnvFn != nil {
			item.EnvFn(v)
		}
	}

	return cfg, nil
}
