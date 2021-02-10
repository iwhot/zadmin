package backend

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/iwhot/zadmin/app/controller"
	"github.com/iwhot/zadmin/app/dao"
	"github.com/iwhot/zadmin/app/logic"
	"github.com/iwhot/zadmin/app/validate"
)

type Category struct {
	controller.Controller
}

func (this *Category) NewRouter(g *gin.RouterGroup) {
	g.GET("/category/list"+controller.BACKENDFIX, this.Index)
	g.GET("/category/add"+controller.BACKENDFIX, this.Add)
	g.POST("/category/add-post"+controller.BACKENDFIX, this.AddPost)
	g.GET("/category/edit/:id"+controller.BACKENDFIX, this.Edit)
	g.POST("/category/edit-post"+controller.BACKENDFIX, this.EditPost)
	g.GET("/category/delete"+controller.BACKENDFIX, this.Delete)
}

//栏目列表
func (this *Category) Index(ctx *gin.Context) {
	//获取分类树形结构
	var list, _ = dao.DefaultCategoryDao.GetCateList(ctx)
	var tree = logic.GetCategoryTree(list, 0)
	this.Render(ctx, "backend/category/index.html", gin.H{
		"list": tree,
	})
}

//添加栏目
func (this *Category) Add(ctx *gin.Context) {
	//取出一级分类
	var list, _ = dao.DefaultCategoryDao.GetCateListByPid(ctx)
	this.Render(ctx, "backend/category/add.html", gin.H{
		"list": list,
	})
}

//添加栏目提交
func (this *Category) AddPost(ctx *gin.Context) {
	//验证
	err := this.Validate(ctx, &validate.Category{})
	if err != nil {
		this.JSON(ctx, gin.H{"code": 0, "msg": fmt.Sprintf("%v", err)})
		return
	}

	err = dao.DefaultCategoryDao.Create(ctx)
	if err != nil {
		this.JSON(ctx, gin.H{"code": 0, "msg": fmt.Sprintf("%v", err)})
		return
	}

	this.JSON(ctx, gin.H{"code": 1, "msg": "添加分类成功"})
}

//编辑栏目
func (this *Category) Edit(ctx *gin.Context) {
	//取出一级分类
	var list, _ = dao.DefaultCategoryDao.GetCateListByPid(ctx)
	//获取分类详情
	var one, _ = dao.DefaultCategoryDao.GetOneCategory(ctx)

	this.Render(ctx, "backend/category/edit.html", gin.H{
		"list": list,
		"info": one,
	})
}

//编辑栏目提交
func (this *Category) EditPost(ctx *gin.Context) {
	//验证
	err := this.Validate(ctx, &validate.Category{})
	if err != nil {
		this.JSON(ctx, gin.H{"code": 0, "msg": fmt.Sprintf("%v", err)})
		return
	}

	err = dao.DefaultCategoryDao.Update(ctx)
	if err != nil {
		this.JSON(ctx, gin.H{"code": 0, "msg": fmt.Sprintf("%v", err)})
		return
	}

	this.JSON(ctx, gin.H{"code": 1, "msg": "编辑分类成功"})
}

//删除栏目
func (this *Category) Delete(ctx *gin.Context) {
	err := dao.DefaultCategoryDao.Delete(ctx)
	if err != nil {
		this.JSON(ctx, gin.H{"code": 0, "msg": fmt.Sprintf("%v", err)})
		return
	}

	this.JSON(ctx, gin.H{"code": 1, "msg": "删除成功"})
}
