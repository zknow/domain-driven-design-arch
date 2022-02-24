package config

import (
	"os"
	"path"

	log "github.com/sirupsen/logrus"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Config struct {
	Server struct {
		Port string `mapstructure:"port"`
	}

	Mail struct {
		Host    string `mapstructure:"host"`
		Port    int    `mapstructure:"port"`
		Account string `mapstructure:"user_account"`
		Passwd  string `mapstructure:"user_password"`
	}
}

var (
	conf       Config
	isWatching bool
)

func init() {
	cDir, err := os.Getwd()
	if err != nil {
		errHandle(err)
	}

	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath(path.Join(cDir, "/config"))
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		errHandle(err)
	}

	readConfigFromFile()

	if !isWatching {
		viper.WatchConfig()
		viper.OnConfigChange(func(e fsnotify.Event) {
			if e.Op == fsnotify.Write {
				log.Println("Config file changed! ...Reload")
				readConfigFromFile()
			}
		})
	}
}

func GetConfig() Config {
	return conf
}

func readConfigFromFile() {
	if err := viper.Unmarshal(&conf); err != nil {
		errHandle(err)
	}
}

func errHandle(err error) {
	log.Fatal("config error :", err)
}
