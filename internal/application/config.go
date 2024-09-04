package application

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	Server struct {
		Port string
	}
	Database struct {
		DSN string
	}
	Log struct {
		Level string
	}
}

func LoadConfig(env string) (*Config, error) {
	viper.SetConfigName(env)
	viper.AddConfigPath("./configs")
	viper.AutomaticEnv()

	var cfg Config
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
		return nil, err
	}

	err := viper.Unmarshal(&cfg)
	if err != nil {
		log.Fatalf("Unable to decode into struct, %v", err)
		return nil, err
	}

	return &cfg, nil
}
