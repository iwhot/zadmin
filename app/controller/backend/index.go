package backend

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/iwhot/zadmin/app/controller"
	"os"
)

type Index struct {
	controller.Controller
}

func (this *Index) NewRouter(g *gin.RouterGroup) {
	g.GET("/index", this.Index)
	g.GET("/menu", this.Menu)
	g.GET("/welcome", this.Welcome)
}

//后台首页
func (this *Index) Index(ctx *gin.Context) {
	this.Render(ctx, "backend/index/index.html", nil)
}

//后台菜单
func (this *Index) Menu(ctx *gin.Context) {
	f, _ := os.Open("config/navs.json")
	defer f.Close()
	var res []interface{}
	d := json.NewDecoder(f)
	_ = d.Decode(&res)
	this.JSON(ctx, res)
}

//欢迎页面
func (this *Index) Welcome(ctx *gin.Context) {
	this.Render(ctx, "backend/index/welcome.html", nil)
}
