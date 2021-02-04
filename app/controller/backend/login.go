package backend

import (
	"github.com/gin-gonic/gin"
	"github.com/iwhot/zadmin/app/controller"
)

type Login struct {
	controller.Controller
}

func (this *Login) NewRouter(g *gin.RouterGroup) {
	g.GET("/backend/login/index"+controller.BACKENDFIX, this.Index)
	g.GET("/backend/login/login"+controller.BACKENDFIX, this.Index)
}

//登录首页
func (this *Login) Index(ctx *gin.Context) {

}

//登录
func (this *Login) Login(ctx *gin.Context) {

}
