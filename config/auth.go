package config

import "time"

type AuthConfig struct {
	JWTSecret string        `mapstructure:"jwt_secret"`
	JWTExpire time.Duration `mapstructure:"jwt_expire"`
}
