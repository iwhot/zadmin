package backend

import (
	"github.com/gin-gonic/gin"
	"github.com/iwhot/zadmin/app/controller"
)

type Menu struct {
	controller.Controller
}

func (this *Menu) NewRouter(g *gin.RouterGroup) {
	g.GET("/backend/menu/menu-list"+controller.BACKENDFIX, this.Index)
	g.GET("/backend/menu/menu-add"+controller.BACKENDFIX, this.Add)
	g.POST("/backend/menu/menu-add-post"+controller.BACKENDFIX, this.AddPost)
	g.GET("/backend/menu/menu-edit"+controller.BACKENDFIX, this.Edit)
	g.POST("/backend/menu/menu-edit-post"+controller.BACKENDFIX, this.EditPost)
	g.GET("/backend/menu/menu-delete"+controller.BACKENDFIX, this.Delete)
}

//菜单列表
func (this *Menu) Index(ctx *gin.Context) {

}

//添加菜单
func (this *Menu) Add(ctx *gin.Context) {

}

//添加菜单提交
func (this *Menu) AddPost(ctx *gin.Context) {

}

//编辑菜单
func (this *Menu) Edit(ctx *gin.Context) {

}

//编辑菜单提交
func (this *Menu) EditPost(ctx *gin.Context) {

}

//删除菜单
func (this *Menu) Delete(ctx *gin.Context) {

}
