package config

import (
	"log"

	"github.com/spf13/viper"
)

var (
	MySQL  *mysql
	JWT    *jwt
	Server *server
)

func Init(path string) {
	viper.SetConfigType("yaml")
	viper.AddConfigPath(path)
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Println("could not find config files")
		} else {
			log.Panicln("read config error")
		}
		log.Fatal(err)
	}

	c := new(config)
	if err := viper.Unmarshal(&c); err != nil {
		log.Fatal(err)
	}

	MySQL = &c.MySQL

	JWT = &c.JWT
	JWT.Secret = []byte(viper.GetString("jwt.secret"))

	Server = &c.Server
}
