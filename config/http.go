package config

type HTTPConfig struct {
	Port       string `mapstructure:"port"`
	Host       string `mapstructure:"host"`
	ExposePort string `mapstructure:"expose_port"`
}
