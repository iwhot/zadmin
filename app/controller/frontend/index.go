package frontend

import (
	"github.com/gin-gonic/gin"
	"github.com/iwhot/zadmin/app/controller"
	"net/http"
)

type Index struct {
	controller.Controller
}

func (this *Index) NewRouter(g *gin.RouterGroup) {
	g.GET("/", this.Index)
}

//首页
func (this *Index) Index(ctx *gin.Context) {
	ctx.String(http.StatusOK, "你好世界")
}
