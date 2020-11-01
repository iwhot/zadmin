package router

import (
	"github.com/gin-gonic/gin"
	"github.com/iwhot/zadmin/app/controller"
	"github.com/iwhot/zadmin/app/controller/api"
	"github.com/iwhot/zadmin/app/controller/backend"
	"github.com/iwhot/zadmin/app/controller/frontend"
	"github.com/iwhot/zadmin/app/middleware"
)

func NewRouter(r *gin.Engine) {
	//设置默认路由当访问一个错误网站时返回
	r.NoRoute(new(controller.Controller).NotFound)
	//前台
	front := r.Group("")
	{
		new(frontend.Index).NewRouter(front)
	}

	//api
	apis := r.Group("/api")
	{
		new(api.Index).NewRouter(apis)
	}

	//后台
	back := r.Group("/backend")
	back.Use(middleware.CheckLogin(0))
	{
		new(backend.Index).NewRouter(back)//后台首页
		new(backend.Member).NewRouter(back)//后台用户管理
	}

	//登录
	login := r.Group("/login")
	login.Use(middleware.CheckLogin(1))
	{
		new(backend.Login).NewRouter(login)
	}

}
