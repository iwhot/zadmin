package backend

import (
	"github.com/gin-gonic/gin"
	"github.com/iwhot/zadmin/app/controller"
)

type Links struct {
	controller.Controller
}

func (this *Links) NewRouter(g *gin.RouterGroup) {
	g.GET("/backend/links/links-list", this.Index)
	g.GET("/backend/links/links-add", this.Add)
	g.POST("/backend/links/links-add-post", this.AddPost)
	g.GET("/backend/links/links-edit", this.Edit)
	g.POST("/backend/links/links-edit-post", this.EditPost)
	g.GET("/backend/links/links-delete", this.Delete)
}

//友链列表
func (this *Links) Index(ctx *gin.Context) {

}

//添加友链
func (this *Links) Add(ctx *gin.Context) {

}

//添加友链提交
func (this *Links) AddPost(ctx *gin.Context) {

}

//编辑友链
func (this *Links) Edit(ctx *gin.Context) {

}

//编辑友链提交
func (this *Links) EditPost(ctx *gin.Context) {

}

//删除友链
func (this *Links) Delete(ctx *gin.Context) {

}
