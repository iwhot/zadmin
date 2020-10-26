package dao

import (
	"fmt"
	"github.com/iwhot/zadmin/system/global"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
)

var masterDB *gorm.DB

func init() {
	global.Init()
	var err error
	//拼接dsn
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		viper.GetString("database.user"),
		viper.GetString("database.password"),
		viper.GetString("database.host"),
		viper.GetString("database.port"),
		viper.GetString("database.dbname"),
		viper.GetString("database.charset"))

	masterDB, err = gorm.Open(viper.GetString("database.prefix"), dsn)
	if err != nil {
		panic(fmt.Sprintf("gorm.Open err :%v,", err))
	}

	if viper.GetString("server.mode") == "debug" {
		masterDB.LogMode(true)
	}

	masterDB.SingularTable(true) //全局设置表名不可以为复数形式。不开表名后面会加s
	// 最大连接数
	masterDB.DB().SetMaxOpenConns(viper.GetInt("database.maxOpenConns"))
	// 闲置连接数
	masterDB.DB().SetMaxIdleConns(viper.GetInt("database.maxIdleConns"))
}
