package conf

import (
	"fmt"

	"github.com/spf13/viper"
)

var (
	HttpPort, WSPort                            int
	ReadTimeout, WriteTimeout                   int
	PayoutGoroutinue, RTPctrl                   int
	Dev                                         bool
	Base64Enable, LoginRepeatEnable             bool
	MaintainEnable, DevOpsMaintain, ServiceKill bool
	GameID, Env                                 string
)

func init() {
	viper.SetConfigName(".env") // name of config file (without extension)
	viper.SetConfigType("env")
	viper.AddConfigPath("./") // optionally look for config in the working directory

	if err := viper.ReadInConfig(); err != nil { // Find and read the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	HttpPort = viper.GetInt("HTTP_PORT")
	WSPort = viper.GetInt("WS_PORT")
	ReadTimeout = viper.GetInt("READ_TIMEOUT")
	WriteTimeout = viper.GetInt("WRITE_TIMEOUT")
	Dev = viper.GetBool("DEV")
	RTPctrl = viper.GetInt("RTP_CTRL")
	Base64Enable = viper.GetBool("BASE64_ENABLE")
	LoginRepeatEnable = viper.GetBool("LOGIN_REPEAT")
	MaintainEnable = viper.GetBool("MAINTAIN")
	DevOpsMaintain = viper.GetBool("DEVOPS_MAINTAIN")
	ServiceKill = viper.GetBool("SERVICE_KILL")
	PayoutGoroutinue = viper.GetInt("PAYOUT_GOROUTINUE")
	GameID = viper.GetString("GAMEID")
	Env = viper.GetString("ENV")
}
