package backend

import (
	"github.com/gin-gonic/gin"
	"github.com/iwhot/zadmin/app/controller"
)

type Ad struct {
	controller.Controller
}

func (this *Ad) NewRouter(g *gin.RouterGroup) {
	g.GET("/backend/ad/ad-list", this.Index)
	g.GET("/backend/ad/ad-add", this.Add)
	g.POST("/backend/ad/ad-add-post", this.AddPost)
	g.GET("/backend/ad/ad-edit", this.Edit)
	g.POST("/backend/ad/ad-edit-post", this.EditPost)
	g.GET("/backend/ad/ad-delete", this.Delete)
}

//广告列表
func (this *Ad) Index(ctx *gin.Context) {

}

//添加广告
func (this *Ad) Add(ctx *gin.Context) {

}

//添加广告提交
func (this *Ad) AddPost(ctx *gin.Context) {

}

//编辑广告
func (this *Ad) Edit(ctx *gin.Context) {

}

//编辑广告提交
func (this *Ad) EditPost(ctx *gin.Context) {

}

//删除广告
func (this *Ad) Delete(ctx *gin.Context) {

}
