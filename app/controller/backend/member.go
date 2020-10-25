package backend

import (
	"github.com/gin-gonic/gin"
	"github.com/iwhot/zadmin/app/controller"
)

type Member struct {
	controller.Controller
}

func (this *Member) NewRouter(g *gin.RouterGroup) {
	g.GET("/backend/member/user-list", this.Index)
	g.GET("/backend/member/user-add", this.Add)
	g.POST("/backend/member/user-add-post", this.AddPost)
	g.GET("/backend/member/user-edit", this.Edit)
	g.POST("/backend/member/user-edit-post", this.EditPost)
	g.GET("/backend/member/user-delete", this.Delete)
}

//用户列表
func (this *Member) Index(ctx *gin.Context) {

}

//添加用户
func (this *Member) Add(ctx *gin.Context) {

}

//添加用户提交
func (this *Member) AddPost(ctx *gin.Context) {

}

//编辑用户
func (this *Member) Edit(ctx *gin.Context) {

}

//编辑用户提交
func (this *Member) EditPost(ctx *gin.Context) {

}

//删除用户
func (this *Member) Delete(ctx *gin.Context) {

}
