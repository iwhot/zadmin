package backend

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/iwhot/zadmin/app/controller"
	"github.com/iwhot/zadmin/app/dao"
	"github.com/iwhot/zadmin/app/logic"
	"github.com/iwhot/zadmin/app/validate"
	"github.com/iwhot/zadmin/system/page"
	"strconv"
)

type Permission struct {
	controller.Controller
}

func (this *Permission) NewRouter(g *gin.RouterGroup) {
	g.GET("/auth-permission/permission-list", this.Index)
	g.GET("/auth-permission/permission-add", this.Add)
	g.POST("/auth-permission/permission-add-post", this.AddPost)
	g.GET("/auth-permission/permission-edit/:id", this.Edit)
	g.POST("/auth-permission/permission-edit-post", this.EditPost)
	g.GET("/auth-permission/permission-delete", this.Delete)
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
	} else {
		param = "&pid=0"
	}

	p := page.GetNowPage(ctx)
	me, _ := dao.DefaultMenuDao.GetMenuList(ctx, controller.PAGESIZE)
	count := dao.DefaultMenuDao.Count(ctx)
	pageHtml := page.NewPage().Pagination(p, controller.PAGESIZE, count, param)

	var pid2, _ = strconv.Atoi(pid)
	b := logic.GetMenuBreadcrumbsFan(uint32(pid2))

	this.Render(ctx, "backend/permission/index.html", gin.H{
		"list":        me,
		"page":        pageHtml,
		"mname":       mname,
		"pid":         pid,
		"Breadcrumbs": b,
	})
}

//添加权限
func (this *Permission) Add(ctx *gin.Context) {
	//获取树形结构
	menus, _ := dao.DefaultMenuDao.GetAllMenuList()
	treeList := logic.GetMenuTree(menus, 0)
	this.Render(ctx, "backend/permission/add.html", gin.H{
		"TreeList": treeList,
	})
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
	//获取树形结构
	menus, _ := dao.DefaultMenuDao.GetAllMenuList()
	treeList := logic.GetMenuTree(menus, 0)
	me, _ := dao.DefaultMenuDao.GetOneMenu(ctx)
	this.Render(ctx, "backend/permission/edit.html", gin.H{
		"MenuInfo": me,
		"TreeList": treeList,
	})
}

//编辑权限提交
func (this *Permission) EditPost(ctx *gin.Context) {
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

//删除权限
func (this *Permission) Delete(ctx *gin.Context) {

	err := dao.DefaultMenuDao.Delete(ctx)
	if err != nil {
		this.JSON(ctx, gin.H{"code": 0, "msg": fmt.Sprintf("%v", err)})
		return
	}
	this.JSON(ctx, gin.H{"code": 1, "msg": "删除成功"})
}
