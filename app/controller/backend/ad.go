package backend

import (
	"github.com/gin-gonic/gin"
	"github.com/iwhot/zadmin/app/controller"
)

type Ad struct {
	controller.Controller
}

func (this *Ad) NewRouter(g *gin.RouterGroup) {
	g.GET("/backend/ad/ad-list"+controller.BACKENDFIX, this.Index)
	g.GET("/backend/ad/ad-add"+controller.BACKENDFIX, this.Add)
	g.POST("/backend/ad/ad-add-post"+controller.BACKENDFIX, this.AddPost)
	g.GET("/backend/ad/ad-edit"+controller.BACKENDFIX, this.Edit)
	g.POST("/backend/ad/ad-edit-post"+controller.BACKENDFIX, this.EditPost)
	g.GET("/backend/ad/ad-delete"+controller.BACKENDFIX, this.Delete)
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
