package backend

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/iwhot/zadmin/app/controller"
	"github.com/iwhot/zadmin/app/dao"
	"github.com/iwhot/zadmin/system/page"
)

type Media struct {
	controller.Controller
}

func (this *Media) NewRouter(g *gin.RouterGroup) {
	g.GET("/media/list"+controller.BACKENDFIX, this.Index)
	g.GET("/media/edit"+controller.BACKENDFIX, this.Edit)
	g.GET("/media/edit-post"+controller.BACKENDFIX, this.EditPost)
	g.GET("/media/delete"+controller.BACKENDFIX, this.Delete)
}

//文件列表
func (this *Media) Index(ctx *gin.Context) {
	p := page.GetNowPage(ctx)
	files, _ := dao.DefaultFilesDao.Find(ctx, p)
	count := dao.DefaultFilesDao.Count(ctx)
	pageHtml := page.NewPage().Pagination(p, controller.PAGESIZE, count, "")
	this.Render(ctx, "backend/media/index.html", gin.H{
		"list": files,
		"page": pageHtml,
	})
}

//编辑文件
func (this *Media) Edit(ctx *gin.Context) {

}

//编辑文件提交
func (this *Media) EditPost(ctx *gin.Context) {

}

//删除文件
func (this *Media) Delete(ctx *gin.Context) {
	err := dao.DefaultFilesDao.Delete(ctx)
	if err != nil {
		this.JSON(ctx, gin.H{"code": 0, "msg": fmt.Sprintf("%v", err)})
		return
	}
	this.JSON(ctx, gin.H{"code": 1, "msg": "删除文件成功"})
}
