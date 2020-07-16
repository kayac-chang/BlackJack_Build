package conf

import (
	"fmt"

	"github.com/spf13/viper"
)

var ULG168APIHost, MaintainAPI, MaintainToken string

func init() {
	viper.SetConfigName(".env") // name of config file (without extension)
	viper.SetConfigType("env")
	viper.AddConfigPath("./") // optionally look for config in the working directory

	if err := viper.ReadInConfig(); err != nil { // Find and read the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	ULG168APIHost = viper.GetString("API_URL")
	MaintainAPI = viper.GetString("MAINTAIN_API")
	MaintainToken = viper.GetString("MAINTAIN_TOKEN")
}
