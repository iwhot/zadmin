package model

import (
	"github.com/iwhot/zadmin/system/global"
	"github.com/spf13/viper"
)

var Prefix string

func init()  {
	global.Init()
	Prefix = viper.GetString("database.prefix")
}

