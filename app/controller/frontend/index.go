package frontend

import (
	"github.com/gin-gonic/gin"
	"github.com/iwhot/zadmin/app/controller"
)

type Index struct {
	controller.Controller
}

func (this *Index) NewRouter(g *gin.RouterGroup) {
	g.GET("/", this.Index)
	g.GET("/index"+controller.FRONTENDFIX, this.Index)
	g.GET("/list"+controller.FRONTENDFIX, this.List)
	g.GET("/show"+controller.FRONTENDFIX, this.Show)
}

//首页
func (this *Index) Index(ctx *gin.Context) {
	this.Render(ctx, "frontend/index/index.html", nil)
}

//列表页
func (this *Index) List(ctx *gin.Context) {
	this.Render(ctx, "frontend/index/list.html", nil)
}

//详情页
func (this *Index) Show(ctx *gin.Context) {
	this.Render(ctx, "frontend/index/show.html", nil)
}
