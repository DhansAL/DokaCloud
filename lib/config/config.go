package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

// Config stores all configuration of the application.
// The values are read by viper from a config file or environment variable.

type DatabaseConfig struct {
	PostgresDSN string `mapstructure:"dsn"`
}
type Config struct {
	Db DatabaseConfig `mapstructure:"database"`
}

// NewConfig reads configuration from file or environment variables.
func NewConfig() (config *Config, err error) {
	fmt.Print("Loading app configs :)")
	viper := viper.New()
	viper.SetConfigName("example")
	viper.AddConfigPath("./config")
	viper.SetConfigType("toml")
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("couldn't load config: %s", err)
		os.Exit(1)
	}
	var c Config
	if err := viper.Unmarshal(&c); err != nil {
		fmt.Printf("couldn't unmarshal config: %s", err)
	}
	fmt.Println("app config loaded")
	return &c, err
}
