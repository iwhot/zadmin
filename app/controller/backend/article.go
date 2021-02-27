package backend

import (
	"github.com/gin-gonic/gin"
	"github.com/iwhot/zadmin/app/controller"
)

type Article struct {
	controller.Controller
}

func (this *Article) NewRouter(g *gin.RouterGroup) {
	g.GET("/article/list"+controller.BACKENDFIX, this.Index)
	g.GET("/article/add"+controller.BACKENDFIX, this.Add)
	g.POST("/article/add-post"+controller.BACKENDFIX, this.AddPost)
	g.GET("/article/edit"+controller.BACKENDFIX, this.Edit)
	g.POST("/article/edit-post/:id"+controller.BACKENDFIX, this.EditPost)
	g.GET("/article/delete"+controller.BACKENDFIX, this.Delete)
}

//文章列表
func (this *Article) Index(ctx *gin.Context) {
	this.Render(ctx, "backend/article/index.html", nil)
}

//添加文章
func (this *Article) Add(ctx *gin.Context) {
	this.Render(ctx, "backend/article/add.html", nil)
}

//添加文章提交
func (this *Article) AddPost(ctx *gin.Context) {

}

//编辑文章
func (this *Article) Edit(ctx *gin.Context) {
	this.Render(ctx, "backend/article/edit.html", nil)
}

//编辑文章提交
func (this *Article) EditPost(ctx *gin.Context) {

}

//删除文章
func (this *Article) Delete(ctx *gin.Context) {

}
