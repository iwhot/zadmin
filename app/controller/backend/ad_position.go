package backend

import (
	"github.com/gin-gonic/gin"
	"github.com/iwhot/zadmin/app/controller"
)

type AdPosition struct {
	controller.Controller
}

func (this *AdPosition) NewRouter(g *gin.RouterGroup) {
	g.GET("/backend/ad-position/user-list"+controller.BACKENDFIX, this.Index)
	g.GET("/backend/ad-position/user-add"+controller.BACKENDFIX, this.Add)
	g.POST("/backend/ad-position/user-add-post"+controller.BACKENDFIX, this.AddPost)
	g.GET("/backend/ad-position/user-edit"+controller.BACKENDFIX, this.Edit)
	g.POST("/backend/ad-position/user-edit-post"+controller.BACKENDFIX, this.EditPost)
	g.GET("/backend/ad-position/user-delete"+controller.BACKENDFIX, this.Delete)
}

//用户列表
func (this *AdPosition) Index(ctx *gin.Context) {

}

//添加用户
func (this *AdPosition) Add(ctx *gin.Context) {

}

//添加用户提交
func (this *AdPosition) AddPost(ctx *gin.Context) {

}

//编辑用户
func (this *AdPosition) Edit(ctx *gin.Context) {

}

//编辑用户提交
func (this *AdPosition) EditPost(ctx *gin.Context) {

}

//删除用户
func (this *AdPosition) Delete(ctx *gin.Context) {

}
