package backend

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/iwhot/zadmin/app/controller"
	"github.com/iwhot/zadmin/app/dao"
	"github.com/iwhot/zadmin/app/validate"
)

type Member struct {
	controller.Controller
}

func (this *Member) NewRouter(g *gin.RouterGroup) {
	g.GET("/member/user-list", this.Index)
	g.GET("/member/user-add", this.Add)
	g.POST("/member/user-add-post", this.AddPost)
	g.GET("/member/user-edit", this.Edit)
	g.POST("/member/user-edit-post", this.EditPost)
	g.GET("/member/user-delete", this.Delete)
}

//用户列表
func (this *Member) Index(ctx *gin.Context) {

	this.Render(ctx, "backend/member/index.html", nil)
}

//添加用户
func (this *Member) Add(ctx *gin.Context) {
	this.Render(ctx, "backend/member/add.html", nil)
}

//添加用户提交
func (this *Member) AddPost(ctx *gin.Context) {
	//验证
	err := this.Validate(ctx, &validate.User{})
	if err != nil {
		this.JSON(ctx, gin.H{"code": 0, "msg": fmt.Sprintf("%v", err)})
		return
	}

	//添加用户
	err = dao.DefaultUserDao.AddUser(ctx)
	if err != nil {
		this.JSON(ctx, gin.H{"code": 0, "msg": fmt.Sprintf("%v", err)})
		return
	}

	//成功添加用户
	this.JSON(ctx, gin.H{"code": 1, "msg": "添加用户成功"})
}

//编辑用户
func (this *Member) Edit(ctx *gin.Context) {
	//获取一个用户

	this.Render(ctx, "backend/member/edit.html", nil)
}

//编辑用户提交
func (this *Member) EditPost(ctx *gin.Context) {

}

//删除用户
func (this *Member) Delete(ctx *gin.Context) {

}
