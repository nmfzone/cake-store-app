package config

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"net/url"
	"strings"
	"sync"
)

var config Config
var singleton sync.Once

type Config struct {
	App struct {
		Host string `mapstructure:"host"`
		Port string `mapstructure:"port"`
	} `mapstructure:"app"`

	Db struct {
		Username string `mapstructure:"username"`
		Password string `mapstructure:"password"`
		Host     string `mapstructure:"host"`
		Port     string `mapstructure:"port"`
		Name     string `mapstructure:"name"`
	} `mapstructure:"db"`
}

func InitConfig() {
	singleton.Do(func() {
		viper.SetConfigFile("config.yml")
		viper.AutomaticEnv()
		viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

		err := viper.ReadInConfig()
		if err != nil {
			panic(err)
		}

		err = viper.Unmarshal(&config)
		if err != nil {
			log.Fatalln("cannot unmarshaling config")
		}
	})
}

func Get() Config {
	return config
}

func Dbdsn() string {
	dbUrl := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s",
		config.Db.Username,
		config.Db.Password,
		config.Db.Host,
		config.Db.Port,
		config.Db.Name)

	val := url.Values{}
	val.Add("parseTime", "1")
	val.Add("loc", "Asia/Jakarta")

	return fmt.Sprintf("%s?%s", dbUrl, val.Encode())
}
