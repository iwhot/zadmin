package backend

import (
	"github.com/gin-gonic/gin"
	"github.com/iwhot/zadmin/app/controller"
)

type Permission struct {
	controller.Controller
}

func (this *Permission) NewRouter(g *gin.RouterGroup) {
	g.GET("/backend/auth/permission-list", this.Index)
	g.GET("/backend/auth/permission-add", this.Add)
	g.POST("/backend/auth/permission-add-post", this.AddPost)
	g.GET("/backend/auth/permission-edit", this.Edit)
	g.POST("/backend/auth/permission-edit-post", this.EditPost)
	g.GET("/backend/auth/permission-delete", this.Delete)
}

//权限列表
func (this *Permission) Index(ctx *gin.Context) {

}

//添加权限
func (this *Permission) Add(ctx *gin.Context) {

}

//添加权限提交
func (this *Permission) AddPost(ctx *gin.Context) {

}

//编辑权限
func (this *Permission) Edit(ctx *gin.Context) {

}

//编辑权限提交
func (this *Permission) EditPost(ctx *gin.Context) {

}

//删除权限
func (this *Permission) Delete(ctx *gin.Context) {

}
