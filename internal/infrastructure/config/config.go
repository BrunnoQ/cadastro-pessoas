package config

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
)

// Config holds all application configuration
type Config struct {
	Server  ServerConfig  `mapstructure:"server"`
	MongoDB MongoDBConfig `mapstructure:"mongodb"`
	Logging LoggingConfig `mapstructure:"logging"`
	App     AppConfig     `mapstructure:"app"`
}

// ServerConfig holds HTTP server configuration
type ServerConfig struct {
	Port         int           `mapstructure:"port"`
	Mode         string        `mapstructure:"mode"` // debug, release
	ReadTimeout  time.Duration `mapstructure:"read_timeout"`
	WriteTimeout time.Duration `mapstructure:"write_timeout"`
	IdleTimeout  time.Duration `mapstructure:"idle_timeout"`
}

// MongoDBConfig holds MongoDB configuration
type MongoDBConfig struct {
	URI             string        `mapstructure:"uri"`
	Database        string        `mapstructure:"database"`
	Timeout         time.Duration `mapstructure:"timeout"`
	MinPoolSize     uint64        `mapstructure:"min_pool_size"`
	MaxPoolSize     uint64        `mapstructure:"max_pool_size"`
	MaxConnIdleTime time.Duration `mapstructure:"max_conn_idle_time"`
}

// LoggingConfig holds logging configuration
type LoggingConfig struct {
	Level  string `mapstructure:"level"`  // debug, info, warn, error
	Format string `mapstructure:"format"` // json, console
}

// AppConfig holds application-specific configuration
type AppConfig struct {
	Name        string `mapstructure:"name"`
	Version     string `mapstructure:"version"`
	Environment string `mapstructure:"environment"` // local, beta, prod
}

// Load reads configuration from YAML file based on environment
func Load(env string) (*Config, error) {
	v := viper.New()

	// Set config file name based on environment
	configFile := fmt.Sprintf("config.%s", env)

	v.SetConfigName(configFile)
	v.SetConfigType("yaml")
	v.AddConfigPath("./configs")
	v.AddConfigPath("../configs")
	v.AddConfigPath("../../configs")

	// Enable environment variable override
	v.AutomaticEnv()

	// Read configuration
	if err := v.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	// Unmarshal into struct
	var config Config
	if err := v.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	// Set environment in app config
	config.App.Environment = env

	// Validate configuration
	if err := config.Validate(); err != nil {
		return nil, fmt.Errorf("invalid configuration: %w", err)
	}

	return &config, nil
}

// Validate checks if configuration is valid
func (c *Config) Validate() error {
	// Validate server config
	if c.Server.Port <= 0 || c.Server.Port > 65535 {
		return fmt.Errorf("invalid server port: %d", c.Server.Port)
	}

	if c.Server.Mode != "debug" && c.Server.Mode != "release" {
		return fmt.Errorf("invalid server mode: %s (must be 'debug' or 'release')", c.Server.Mode)
	}

	// Validate MongoDB config
	if c.MongoDB.URI == "" {
		return fmt.Errorf("mongodb URI is required")
	}

	if c.MongoDB.Database == "" {
		return fmt.Errorf("mongodb database name is required")
	}

	if c.MongoDB.MinPoolSize > c.MongoDB.MaxPoolSize {
		return fmt.Errorf("mongodb min_pool_size (%d) cannot be greater than max_pool_size (%d)",
			c.MongoDB.MinPoolSize, c.MongoDB.MaxPoolSize)
	}

	// Validate logging config
	validLogLevels := map[string]bool{"debug": true, "info": true, "warn": true, "error": true}
	if !validLogLevels[c.Logging.Level] {
		return fmt.Errorf("invalid log level: %s (must be debug, info, warn, or error)", c.Logging.Level)
	}

	validLogFormats := map[string]bool{"json": true, "console": true}
	if !validLogFormats[c.Logging.Format] {
		return fmt.Errorf("invalid log format: %s (must be json or console)", c.Logging.Format)
	}

	return nil
}
