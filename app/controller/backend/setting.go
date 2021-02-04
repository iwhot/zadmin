package backend

import (
	"github.com/gin-gonic/gin"
	"github.com/iwhot/zadmin/app/controller"
)

type Setting struct {
	controller.Controller
}

func (this *Setting) NewRouter(g *gin.RouterGroup) {
	g.GET("/backend/setting/index"+controller.BACKENDFIX, this.Index)
}

//配置管理
func (this *Setting) Index(ctx *gin.Context) {

}
