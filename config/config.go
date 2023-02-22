package config

import (
	"log"
	"os"
	"path"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Config struct {
	Port string `mapstructure:"port"`
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
