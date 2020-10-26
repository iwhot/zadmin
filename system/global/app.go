package global

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
	"github.com/spf13/viper"
)

var once sync.Once
func Init()  {
	once.Do(func() {
		// 随机数种子
		rand.Seed(time.Now().UnixNano())
		//配置名称
		viper.SetConfigName("config")
		viper.AddConfigPath("config/")
		viper.SetConfigType("yaml")
		//读取配置文件
		err := viper.ReadInConfig()
		if err != nil{
			panic(fmt.Errorf("Fatal error config file: %s \n", err))
		}
	})
}
