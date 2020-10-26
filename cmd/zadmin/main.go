package main

import (
	"github.com/gin-gonic/gin"
	"github.com/iwhot/zadmin/router"
	"github.com/iwhot/zadmin/system/global"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"time"
)

func init() {
	global.Init()
}

func main() {
	//设置运行模式
	gin.SetMode(viper.GetString("server.mode"))
	//获取gin
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(gin.Logger())
	//设置静态文件位置
	r.Static("/static", "storage/static")
	r.Static("/uploads", "storage/uploads")
	//載入模板
	r.LoadHTMLGlob("template/**/**/*")
	//表单限制上传大小
	r.MaxMultipartMemory = 8 << 20
	//载入路由
	router.NewRouter(r)
	s := &http.Server{
		Addr:           viper.GetString("server.host") + ":" + viper.GetString("server.port"),
		Handler:        r,
		ReadTimeout:    time.Duration(viper.GetInt("server.readTimeout")) * time.Second,
		WriteTimeout:   time.Duration(viper.GetInt("server.writeTimeout")) * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	//打印错误
	log.Fatalf("router err : %v", s.ListenAndServe())
}
