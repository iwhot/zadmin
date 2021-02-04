package backend

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/iwhot/zadmin/app/controller"
	"github.com/iwhot/zadmin/app/dao"
	"github.com/iwhot/zadmin/app/logic"
	"github.com/iwhot/zadmin/app/validate"
	"github.com/iwhot/zadmin/system/page"
)

type Role struct {
	controller.Controller
}

func (this *Role) NewRouter(g *gin.RouterGroup) {
	g.GET("/auth-role/role-list"+controller.BACKENDFIX, this.Index)
	g.GET("/auth-role/role-add"+controller.BACKENDFIX, this.Add)
	g.POST("/auth-role/role-add-post"+controller.BACKENDFIX, this.AddPost)
	g.GET("/auth-role/role-edit/:id"+controller.BACKENDFIX, this.Edit)
	g.POST("/auth-role/role-edit-post"+controller.BACKENDFIX, this.EditPost)
	g.GET("/auth-role/role-delete"+controller.BACKENDFIX, this.Delete)
	g.GET("/auth-role/role-auth"+controller.BACKENDFIX, this.Auth)
}

//角色列表
func (this *Role) Index(ctx *gin.Context) {
	var roleName = ctx.Query("role_name")

	var param string
	param = "zs=1";
	if roleName != "" {
		param = "&role_name=" + roleName
	}

	p := page.GetNowPage(ctx)
	rol, _ := dao.DefaultRoleDao.Find(ctx, controller.PAGESIZE)
	count := dao.DefaultRoleDao.Count(ctx)
	pageHtml := page.NewPage().Pagination(p, controller.PAGESIZE, count, param)
	this.Render(ctx, "backend/role/index.html", gin.H{
		"list":      rol,
		"page":      pageHtml,
		"role_name": roleName,
	})
}

//添加角色
func (this *Role) Add(ctx *gin.Context) {
	this.Render(ctx, "backend/role/add.html", nil)
}

//添加角色提交
func (this *Role) AddPost(ctx *gin.Context) {
	//验证
	err := this.Validate(ctx, &validate.Role{})
	if err != nil {
		this.JSON(ctx, gin.H{"code": 0, "msg": fmt.Sprintf("%v", err)})
		return
	}
	err = dao.DefaultRoleDao.Create(ctx)
	if err != nil {
		this.JSON(ctx, gin.H{"code": 0, "msg": fmt.Sprintf("%v", err)})
		return
	}
	//成功添加用户
	this.JSON(ctx, gin.H{"code": 1, "msg": "添加角色成功"})
}

//编辑角色
func (this *Role) Edit(ctx *gin.Context) {
	roles, _ := dao.DefaultRoleDao.GetOneRoleInfo(ctx)
	this.Render(ctx, "backend/role/edit.html", gin.H{
		"RoleInfo": roles,
	})
}

//编辑角色提交
func (this *Role) EditPost(ctx *gin.Context) {
	//验证
	err := this.Validate(ctx, &validate.Role{})
	if err != nil {
		this.JSON(ctx, gin.H{"code": 0, "msg": fmt.Sprintf("%v", err)})
		return
	}

	err = dao.DefaultRoleDao.Update(ctx)
	if err != nil {
		this.JSON(ctx, gin.H{"code": 0, "msg": fmt.Sprintf("%v", err)})
		return
	}

	//成功添加用户
	this.JSON(ctx, gin.H{"code": 1, "msg": "编辑角色成功"})
}

//删除角色
func (this *Role) Delete(ctx *gin.Context) {
	err := dao.DefaultRoleDao.Delete(ctx)
	if err != nil {
		this.JSON(ctx, gin.H{"code": 0, "msg": fmt.Sprintf("%v", err)})
		return
	}

	this.JSON(ctx, gin.H{"code": 1, "msg": "删除角色成功"})
}

//权限
func (this *Role) Auth(ctx *gin.Context) {
	ral, _ := dao.DefaultMenuDao.GetAllMenuList()

	r, _ := dao.DefaultRoleMenuDao.GetOneRoleMenuList(ctx)
	var chk []uint32
	for _, v := range r {
		chk = append(chk, v.ID)
	}

	RoleMenuTree := logic.GetRoleMenuNode(ral, 0, chk)
	this.JSON(ctx, RoleMenuTree)
}
