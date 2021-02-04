package backend

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/iwhot/zadmin/app/controller"
	"github.com/iwhot/zadmin/app/dao"
	"github.com/iwhot/zadmin/app/validate"
	"github.com/iwhot/zadmin/system/page"
)

type Member struct {
	controller.Controller
}

func (this *Member) NewRouter(g *gin.RouterGroup) {
	g.GET("/member/user-list"+controller.BACKENDFIX, this.Index)
	g.GET("/member/user-add"+controller.BACKENDFIX, this.Add)
	g.POST("/member/user-add-post"+controller.BACKENDFIX, this.AddPost)
	g.GET("/member/user-edit/:id"+controller.BACKENDFIX, this.Edit)
	g.POST("/member/user-edit-post"+controller.BACKENDFIX, this.EditPost)
	g.GET("/member/user-delete"+controller.BACKENDFIX, this.Delete)
}

//用户列表
func (this *Member) Index(ctx *gin.Context) {
	var username = ctx.Query("username")
	var email = ctx.Query("email")
	var rname = ctx.Query("rname")
	var param string
	param = "zs=1";
	if username != "" {
		param = "&username=" + username
	}

	if email != "" {
		param = "&email=" + email
	}

	if rname != "" {
		param = "&rname=" + rname
	}

	p := page.GetNowPage(ctx)
	user, _ := dao.DefaultUserDao.GetUserList(ctx, controller.PAGESIZE)
	count := dao.DefaultUserDao.Count(ctx)

	pageHtml := page.NewPage().Pagination(p, controller.PAGESIZE, count, param)
	this.Render(ctx, "backend/member/index.html", gin.H{
		"list":     user,
		"page":     pageHtml,
		"username": username,
		"email":    email,
		"rname":    rname,
	})

}

//添加用户
func (this *Member) Add(ctx *gin.Context) {
	//获取所有角色
	roles, _ := dao.DefaultRoleDao.GetAllRoleList()
	this.Render(ctx, "backend/member/add.html", gin.H{
		"Roles": roles,
	})
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
	//获取所有角色
	roles, _ := dao.DefaultRoleDao.GetAllRoleList()
	//获取一个用户
	user, _ := dao.DefaultUserDao.GetOneUserInfo(ctx)
	this.Render(ctx, "backend/member/edit.html", gin.H{
		"UserInfo": user,
		"Roles":    roles,
	})
}

//编辑用户提交
func (this *Member) EditPost(ctx *gin.Context) {
	//验证
	err := this.Validate(ctx, &validate.UserEdit{})
	if err != nil {
		this.JSON(ctx, gin.H{"code": 0, "msg": fmt.Sprintf("%v", err)})
		return
	}

	//编辑用户
	err = dao.DefaultUserDao.Update(ctx)
	if err != nil {
		this.JSON(ctx, gin.H{"code": 0, "msg": fmt.Sprintf("%v", err)})
		return
	}

	//成功添加用户
	this.JSON(ctx, gin.H{"code": 1, "msg": "编辑用户成功"})
}

//删除用户
func (this *Member) Delete(ctx *gin.Context) {
	err := dao.DefaultUserDao.Delete(ctx)
	if err != nil {
		this.JSON(ctx, gin.H{"code": 0, "msg": fmt.Sprintf("%v", err)})
		return
	}

	this.JSON(ctx, gin.H{"code": 1, "msg": "删除用户成功"})
}
