package api

import (
	"github.com/gin-gonic/gin"
	"github.com/iwhot/zadmin/app/controller"
)

type Index struct {
	controller.Controller
}

func (this *Index) NewRouter(g *gin.RouterGroup) {
	g.GET("/index", this.Index)
}

func (this *Index) Index(ctx *gin.Context) {
	this.JSON(ctx, gin.H{
		"code": 0,
		"msg":  "this is api index",
	})
}
