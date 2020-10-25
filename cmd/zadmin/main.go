package main

import (
	"github.com/gin-gonic/gin"
	"github.com/iwhot/zadmin/router"
	"log"
	"net/http"
	"time"
)

func main() {
	//设置运行模式
	gin.SetMode("debug")
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
		Addr:           ":2020",
		Handler:        r,
		ReadTimeout:    60 * time.Second,
		WriteTimeout:   60 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	//打印错误
	log.Fatalf("router err : %v", s.ListenAndServe())
}
