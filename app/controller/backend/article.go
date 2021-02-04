package backend

import (
	"github.com/gin-gonic/gin"
	"github.com/iwhot/zadmin/app/controller"
)

type Article struct {
	controller.Controller
}

func (this *Article) NewRouter(g *gin.RouterGroup) {
	g.GET("/backend/article/article-list"+controller.BACKENDFIX, this.Index)
	g.GET("/backend/article/article-add"+controller.BACKENDFIX, this.Add)
	g.POST("/backend/article/article-add-post"+controller.BACKENDFIX, this.AddPost)
	g.GET("/backend/article/article-edit"+controller.BACKENDFIX, this.Edit)
	g.POST("/backend/article/article-edit-post"+controller.BACKENDFIX, this.EditPost)
	g.GET("/backend/article/article-delete"+controller.BACKENDFIX, this.Delete)
}

//文章列表
func (this *Article) Index(ctx *gin.Context) {

}

//添加文章
func (this *Article) Add(ctx *gin.Context) {

}

//添加文章提交
func (this *Article) AddPost(ctx *gin.Context) {

}

//编辑文章
func (this *Article) Edit(ctx *gin.Context) {

}

//编辑文章提交
func (this *Article) EditPost(ctx *gin.Context) {

}

//删除文章
func (this *Article) Delete(ctx *gin.Context) {

}
