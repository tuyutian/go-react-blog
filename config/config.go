package config

import (
	"gopkg.in/yaml.v2"
	"os"
	"time"
)

type Config struct {
	GinMode          string        `yaml:"gin_mode"`
	Addr             string        `yaml:"addr"`
	DSN              string        `yaml:"dsn"`
	MaxIdleConn      int           `yaml:"max_idle_conn"`
	JwtSecret        string        `yaml:"jwt_secret"`
	JwtExpireDay     int           `yaml:"jwt_expire_day"`
	RedisHost        string        `yaml:"redis_host"`
	RedisPassword    string        `yaml:"redis_password"`
	RedisMaxIdle     int           `yaml:"redis_max_idle"`
	RedisActive      int           `yaml:"redis_active"`
	RedisIdleTimeout time.Duration `yaml:"redis_idle_timeout"`
	SQLDebug         bool          `yaml:"sql_debug"`
	SQLLite          string        `yaml:"sql_lite"`
	RabbitHost       string        `yaml:"rabbit_host"`
}

var config *Config

func Load(path string) error {
	result, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	return yaml.Unmarshal(result, &config)
}

func Get() *Config {
	return config
}
