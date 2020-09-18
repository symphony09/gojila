package global

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var Config = viper.New()

func init() {
	Config.SetConfigName("config")
	Config.SetConfigType("yml")
	Config.AddConfigPath(".")
	if err := Config.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			Config.SetDefault("address", "0.0.0.0:8080")
			Config.SetDefault("db.host", "127.0.0.1")
			Config.SetDefault("db.port", "3306")
			Config.SetDefault("db.user", "root")
			Config.SetDefault("db.password", "olives")
			Config.SetDefault("db.name", "gojila")
		} else {
			fmt.Println("Config file was found but another error was produced")
		}
	}
	Config.WatchConfig()
	Config.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
	})
}
