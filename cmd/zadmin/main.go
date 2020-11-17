package main

import (
	"github.com/gin-gonic/gin"
	"github.com/iwhot/zadmin/router"
	"github.com/iwhot/zadmin/system/common"
	"github.com/iwhot/zadmin/system/global"
	"github.com/spf13/viper"
	"html/template"
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

	//设置模板函数要在载入模板前面
	r.SetFuncMap(template.FuncMap{
		"UnixToDateTime": common.UnixToDateTime,
	})

	//載入模板
	r.LoadHTMLGlob("template/**/**/*")
	//表单限制上传大小
	r.MaxMultipartMemory = 8 << 20

	rd, _ := time.ParseDuration(viper.GetString("server.readTimeout"))
	wd, _ := time.ParseDuration(viper.GetString("server.writeTimeout"))

	//载入路由
	router.NewRouter(r)
	s := &http.Server{
		Addr:           viper.GetString("server.host") + ":" + viper.GetString("server.port"),
		Handler:        r,
		ReadTimeout:    rd,
		WriteTimeout:   wd,
		MaxHeaderBytes: 1 << 20,
	}
	//打印错误
	log.Fatalf("router err : %v", s.ListenAndServe())
}
