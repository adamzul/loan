package config

import (
	"log"

	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
)

type Config struct {
	Auth     AuthConfig          `mapstructure:"auth"`
	Postgres DBReplicationConfig `mapstructure:"postgres"`
	HTTP     HTTPConfig          `mapstructure:"http"`
}

type DBReplicationConfig struct {
	Primary DBConfig `mapstructure:"primary"`
	Standby DBConfig `mapstructure:"standby"`
}

func Load() (*Config, error) {
	err := readFromJSON()
	if err != nil {
		return nil, err
	}

	appConf := &Config{}
	err = viper.Unmarshal(appConf, func(dc *mapstructure.DecoderConfig) {
		// Prevent service to bootup if configuration is missing
		dc.ErrorUnset = true
		dc.ErrorUnused = false
	})

	if err != nil {
		return nil, err
	}

	return appConf, nil
}

func MustLoad() *Config {
	conf, err := Load()
	if err != nil {
		log.Fatalf("config: failed to load %s", err)
	}

	return conf
}

func readFromJSON() error {
	viper.SetConfigFile("./config.json")
	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	return nil
}
