package conf

import (
	"fmt"

	"github.com/spf13/viper"
)

var (
	MysqlConf    *mysqlConf
	Debug        bool
	Singular     bool
	MaxIdleConns int
	MaxOpenConns int
)

func init() {
	viper.SetConfigName(".env") // name of config file (without extension)
	viper.SetConfigType("env")
	viper.AddConfigPath("./") // optionally look for config in the working directory

	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	MysqlConf = &mysqlConf{
		Host:     viper.GetString("DB_IP"),
		Port:     viper.GetString("DB_PORT"),
		Database: viper.GetString("DB_NAME"),
		// MysqlTimeZone : viper.GetString("timezone"),
		Charset:      viper.GetString("DB_CHARSET"),
		User:         viper.GetString("DB_USER"),
		Passwd:       viper.GetString("DB_PASSWORD"),
		ParseTime:    viper.GetString("DB_PARSE_TIME"),
		Debug:        viper.GetBool("DB_DEBUG"),
		Singular:     viper.GetBool("DB_SINGULAR"),
		MaxIdleConns: viper.GetInt("DB_MAX_IDLE_CONNECTION"),
		MaxOpenConns: viper.GetInt("DB_MAX_OPEN_CONNECTION"),
	}

	fmt.Println(MysqlConf.ConverToPath())
}

type mysqlConf struct {
	Host              string
	Port              string
	User              string
	Passwd            string
	Database          string
	Charset           string
	ColumnsWithAlias  bool
	InterpolateParams bool
	ParseTime         string
	Timeout           string
	ReadTimeout       string
	WriteTimeout      string

	Singular     bool
	Debug        bool
	MaxIdleConns int
	MaxOpenConns int
}

func (mc *mysqlConf) ConverToPath() string {
	path := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=%s",
		mc.User, mc.Passwd, mc.Host, mc.Port, mc.Database, mc.Charset, mc.ParseTime)

	return path
}
