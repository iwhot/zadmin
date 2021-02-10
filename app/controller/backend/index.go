package backend

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/iwhot/zadmin/app/controller"
	"github.com/iwhot/zadmin/app/dao"
	"github.com/iwhot/zadmin/app/model"
	"github.com/iwhot/zadmin/system/common"
	"os"
	"path/filepath"
	"time"
)

type Index struct {
	controller.Controller
}

func (this *Index) NewRouter(g *gin.RouterGroup) {
	g.GET("/index"+controller.BACKENDFIX, this.Index)
	g.GET("/menu"+controller.BACKENDFIX, this.Menu)
	g.GET("/welcome"+controller.BACKENDFIX, this.Welcome)
	g.POST("/upload"+controller.BACKENDFIX, this.Upload)
}

//后台首页
func (this *Index) Index(ctx *gin.Context) {
	this.Render(ctx, "backend/index/index.html", nil)
}

//后台菜单
func (this *Index) Menu(ctx *gin.Context) {
	f, _ := os.Open("config/navs.json")
	defer f.Close()
	var res []interface{}
	d := json.NewDecoder(f)
	_ = d.Decode(&res)
	this.JSON(ctx, res)
}

//欢迎页面
func (this *Index) Welcome(ctx *gin.Context) {
	this.Render(ctx, "backend/index/welcome.html", nil)
}

//文件上传
func (this *Index) Upload(ctx *gin.Context) {
	file, err := ctx.FormFile("file")
	if err != nil {
		this.JSON(ctx, gin.H{"code": 0, "msg": fmt.Sprintf("%v", err)})
		return
	}

	paths := "/uploads/" + time.Now().Format("20060102") + "/"
	err = common.CreateFile(controller.STORAGEPATH + paths)
	if err != nil {
		this.JSON(ctx, gin.H{"code": 0, "msg": fmt.Sprintf("%v", err)})
		return
	}

	name := file.Filename
	size := file.Size
	//获取后缀
	ext := filepath.Ext(name)
	filename := paths + time.Now().Format("20060102") + common.CreateCaptcha() + ext
	err = ctx.SaveUploadedFile(file, controller.STORAGEPATH+filename)
	if err != nil {
		this.JSON(ctx, gin.H{"code": 0, "msg": fmt.Sprintf("%v", err)})
		return
	}

	var f = model.Files{
		Name:  name,
		Url:   filename,
		Size:  float64(size),
		Type:  0,
		Ctime: uint32(time.Now().Unix()),
		State: common.GetFileType(ext),
		Utime: uint32(time.Now().Unix()),
	}

	err = dao.DefaultFilesDao.Create(f)
	if err != nil {
		this.JSON(ctx, gin.H{"code": 0, "msg": fmt.Sprintf("%v", err)})
		return
	}

	this.JSON(ctx, gin.H{"code": 1, "msg": "上传成功", "url": filename})
}
