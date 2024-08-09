package config

import (
	"time"
)

type DBConfig struct {
	Username    string        `mapstructure:"username"`
	Password    string        `mapstructure:"password"`
	Name        string        `mapstructure:"name"`
	Host        string        `mapstructure:"host"`
	Port        string        `mapstructure:"port"`
	MaxOpen     int32         `mapstructure:"max-open"`
	MaxIdle     int32         `mapstructure:"max-idle"`
	MaxIdleTime time.Duration `mapstructure:"max-idle-time"`
}
