package env

import (
	"log"

	"github.com/spf13/viper"
)

func InitConfig() *viper.Viper {
	viper.SetConfigType("json")
	viper.AddConfigPath("./env")
	viper.SetConfigName("config")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}

	return viper.GetViper()
}
