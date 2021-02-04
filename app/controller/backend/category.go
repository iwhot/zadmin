package backend

import (
	"github.com/gin-gonic/gin"
	"github.com/iwhot/zadmin/app/controller"
)

type Category struct {
	controller.Controller
}

func (this *Category) NewRouter(g *gin.RouterGroup) {
	g.GET("/backend/category/category-list"+controller.BACKENDFIX, this.Index)
	g.GET("/backend/category/category-add"+controller.BACKENDFIX, this.Add)
	g.POST("/backend/category/category-add-post"+controller.BACKENDFIX, this.AddPost)
	g.GET("/backend/category/category-edit"+controller.BACKENDFIX, this.Edit)
	g.POST("/backend/category/category-edit-post"+controller.BACKENDFIX, this.EditPost)
	g.GET("/backend/category/category-delete"+controller.BACKENDFIX, this.Delete)
}

//栏目列表
func (this *Category) Index(ctx *gin.Context) {

	this.Render(ctx, "backend/category/index.html", nil)
}

//添加栏目
func (this *Category) Add(ctx *gin.Context) {
	this.Render(ctx, "backend/category/add.html", nil)
}

//添加栏目提交
func (this *Category) AddPost(ctx *gin.Context) {

}

//编辑栏目
func (this *Category) Edit(ctx *gin.Context) {
	this.Render(ctx, "backend/category/edit.html", nil)
}

//编辑栏目提交
func (this *Category) EditPost(ctx *gin.Context) {

}

//删除栏目
func (this *Category) Delete(ctx *gin.Context) {

}
