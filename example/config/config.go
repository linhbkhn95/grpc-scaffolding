package config

import (
	"bytes"
	"strings"

	"github.com/spf13/viper"

	"github.com/linhbkhn95/golang-british/appmode"
	"github.com/linhbkhn95/golang-british/logger"

	"github.com/linhbkhn95/example/server/grpc"
)

// Config holds all settings
var defaultConfig = []byte(`
app_mode: development
log:
   enable_console: true
   console_log_format: false
   console_level: info
server:
  grpc:
   host: 0.0.0.0
   port: 10443
  http:
   host: 0.0.0.0
   port: 10080

`)

// AppConfig is centralized application configs
var (
	Mode   appmode.AppMode
	Log    logger.Configuration
	Server server.Config
)

type (
	Config struct {
		Log     logger.Configuration `mapstructure:"log"`
		AppMode string               `mapstructure:"app_mode"`
		Server  server.Config        `mapstructure:"server"`
	}
)

func init() {
	load()
}

func load() {
	var cfg = &Config{}

	viper.SetConfigType("yaml")
	err := viper.ReadConfig(bytes.NewBuffer(defaultConfig))
	if err != nil {
		logger.Fatalf("Failed to read viper config %v", err)
	}

	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	if err := viper.Unmarshal(&cfg); err != nil {
		logger.Fatalf("Failed to unmarshal config %v", err)
	}
	Mode = appmode.GetAppMode(cfg.AppMode)
	Log = cfg.Log
	Server = cfg.Server
}
