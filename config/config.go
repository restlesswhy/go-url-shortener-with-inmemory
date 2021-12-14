package config

import (
	"errors"

	"github.com/spf13/viper"
)

type Config struct {
	Logger Logger
}

type Logger struct {
	Encoding string
	Level string
}

func Loadconfig(configName string) (*viper.Viper, error) {
	v := viper.New()

	v.SetConfigName(configName)
	v.AddConfigPath(".")
	v.AutomaticEnv()
	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileAlreadyExistsError); ok {
			return nil, errors.New("config file not found")
		}
		return nil, err
	}

	return v, nil
}

func ParseConfig(v *viper.Viper) (*Config, error) {
	var c Config

	if err := v.Unmarshal(&c); err != nil {
		return nil, err
	}

	return &c, nil
}

func GetConfig(configPath string) (*Config, error) {
	cfgFile, err := Loadconfig(configPath)
	if err != nil {
		return nil, err
	}

	cfg, err := ParseConfig(cfgFile)
	if err != nil {
		return nil, err
	}

	return cfg, err
}

func GetConfigPath(configPath string) string {
	if configPath == "some" {
		return "some path"
	}
	
	return "./config/config-local"
}