package backend

import (
	"github.com/gin-gonic/gin"
	"github.com/iwhot/zadmin/app/controller"
)

type Role struct {
	controller.Controller
}

func (this *Role) NewRouter(g *gin.RouterGroup) {
	g.GET("/backend/auth-role/role-list", this.Index)
	g.GET("/backend/auth-role/role-add", this.Add)
	g.POST("/backend/auth-role/role-add-post", this.AddPost)
	g.GET("/backend/auth-role/role-edit", this.Edit)
	g.POST("/backend/auth-role/role-edit-post", this.EditPost)
	g.GET("/backend/auth-role/role-delete", this.Delete)
}

//角色列表
func (this *Role) Index(ctx *gin.Context) {

}

//添加角色
func (this *Role) Add(ctx *gin.Context) {

}

//添加角色提交
func (this *Role) AddPost(ctx *gin.Context) {

}

//编辑角色
func (this *Role) Edit(ctx *gin.Context) {

}

//编辑角色提交
func (this *Role) EditPost(ctx *gin.Context) {

}

//删除角色
func (this *Role) Delete(ctx *gin.Context) {

}
