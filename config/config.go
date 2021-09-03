package config

import (
	"os"
	"time"

	"path/filepath"

	"github.com/davecgh/go-spew/spew"
	"github.com/golang/glog"
	"github.com/spf13/viper"
)

type config struct {
	Database struct {
		User                 string
		Password             string
		Net                  string
		Addr                 string
		DBName               string
		AllowNativePasswords bool
		Params               struct {
			ParseTime string
		}
		Timeout time.Duration
	}
	Server struct {
		Address string
	}
}

var C config

func ReadConfig() {
	Config := &C

	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath(filepath.Join("config"))
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		glog.Info("Error to read config file: %+v", err)
	}

	if err := viper.Unmarshal(&Config); err != nil {
		glog.Info("Error to decode config file: %+v", err)
		os.Exit(1)
	}

	spew.Dump(C)
}
