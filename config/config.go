package config

import (
	"errors"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	Postgres  PostgresConfig
	Server    ServerConfig
	Shortener ShortenerConfig
	CallAt CallAtConfig
}

type CallAtConfig struct {
	SleepingTime time.Duration

}

type ServerConfig struct {
	Port              string
	ReadTimeout       time.Duration
	WriteTimeout      time.Duration
	CtxDefaultTimeout time.Duration
	MaxConnectionIdle time.Duration
	Timeout           time.Duration
	MaxConnectionAge  time.Duration
	Time              time.Duration
}

type PostgresConfig struct {
	PostgresqlHost     string
	PostgresqlPort     string
	PostgresqlUser     string
	PostgresqlPassword string
	PostgresqlDbname   string
	PostgresqlSSLMode  bool
	PgDriver           string
}

type ShortenerConfig struct {
	StringLength int
	Runes        string
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

	return "./config/config"
}
