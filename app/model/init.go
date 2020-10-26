package model

import "github.com/spf13/viper"

var pre = viper.GetString("database.prefix")

func init()  {
	//todo
}
