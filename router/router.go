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
	back.Use(middleware.CheckLogin(0), middleware.BackendAuth())
	{
		new(backend.Index).NewRouter(back)      //后台首页
		new(backend.Member).NewRouter(back)     //后台用户管理
		new(backend.Ad).NewRouter(back)         //后台广告位
		new(backend.AdPosition).NewRouter(back) //后台广告位管理
		new(backend.Article).NewRouter(back)    //文章管理
		new(backend.Category).NewRouter(back)   //分类管理
		new(backend.Links).NewRouter(back)      //链接管理
		new(backend.Menu).NewRouter(back)       //菜单管理
		new(backend.Permission).NewRouter(back) //权限管理
		new(backend.Role).NewRouter(back)       //角色管理
		new(backend.Setting).NewRouter(back)    //设置
	}

	//后台登录
	login := r.Group("/login")
	login.Use(middleware.CheckLogin(1))
	{
		new(backend.Login).NewRouter(login)
	}

}
