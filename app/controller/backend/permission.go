package backend

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/iwhot/zadmin/app/controller"
	"github.com/iwhot/zadmin/app/dao"
	"github.com/iwhot/zadmin/app/validate"
	"github.com/iwhot/zadmin/system/page"
)

type Permission struct {
	controller.Controller
}

func (this *Permission) NewRouter(g *gin.RouterGroup) {
	g.GET("/backend/auth-permission/permission-list", this.Index)
	g.GET("/backend/auth-permission/permission-add", this.Add)
	g.POST("/backend/auth-permission/permission-add-post", this.AddPost)
	g.GET("/backend/auth-permission/permission-edit", this.Edit)
	g.POST("/backend/auth-permission/permission-edit-post", this.EditPost)
	g.GET("/backend/auth-permission/permission-delete", this.Delete)
}

//权限列表
func (this *Permission) Index(ctx *gin.Context) {
	var mname = ctx.Query("mname")
	var pid = ctx.Query("pid")
	var param string
	param = "zs=1";
	if mname != "" {
		param = "&mname=" + mname
	}

	if pid != "" {
		param = "&pid=" + pid
	}
	p := page.GetNowPage(ctx)
	me, _ := dao.DefaultMenuDao.GetMenuList(ctx, controller.PAGESIZE)
	count := dao.DefaultMenuDao.Count(ctx)
	pageHtml := page.NewPage().Pagination(p, controller.PAGESIZE, count, param)

	this.Render(ctx, "backend/member/index.html", gin.H{
		"list":  me,
		"page":  pageHtml,
		"mname": mname,
		"pid":   pid,
	})
}

//添加权限
func (this *Permission) Add(ctx *gin.Context) {
	this.Render(ctx, "backend/permission/add.html", nil)
}

//添加权限提交
func (this *Permission) AddPost(ctx *gin.Context) {
	err := this.Validate(ctx, &validate.Menu{})
	if err != nil {
		this.JSON(ctx, gin.H{"code": 0, "msg": fmt.Sprintf("%v", err)})
		return
	}

	err = dao.DefaultMenuDao.Create(ctx)
	if err != nil {
		this.JSON(ctx, gin.H{"code": 0, "msg": fmt.Sprintf("%v", err)})
		return
	}

	this.JSON(ctx, gin.H{"code": 1, "msg": "添加成功"})
}

//编辑权限
func (this *Permission) Edit(ctx *gin.Context) {
	me, _ := dao.DefaultMenuDao.GetOneMenu(ctx)
	this.Render(ctx, "backend/member/edit.html", gin.H{
		"MenuInfo": me,
	})
}

//编辑权限提交
func (this *Permission) EditPost(ctx *gin.Context) {

}

//删除权限
func (this *Permission) Delete(ctx *gin.Context) {
	err := this.Validate(ctx, &validate.Menu{})
	if err != nil {
		this.JSON(ctx, gin.H{"code": 0, "msg": fmt.Sprintf("%v", err)})
		return
	}

	err = dao.DefaultMenuDao.Update(ctx)
	if err != nil {
		this.JSON(ctx, gin.H{"code": 0, "msg": fmt.Sprintf("%v", err)})
		return
	}
	this.JSON(ctx, gin.H{"code": 1, "msg": "编辑成功"})
}
