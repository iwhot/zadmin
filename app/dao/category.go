package dao

import (
	"github.com/gin-gonic/gin"
	"github.com/iwhot/zadmin/app/model"
	"github.com/iwhot/zadmin/system/page"
	"strconv"
	"github.com/satori/go.uuid"
	"time"
)

type category struct {
}

var DefaultCategoryDao = category{}

//获取列表
func (c category) GetCateList(ctx *gin.Context, pz int) ([]*model.Category, error) {
	var pid, _ = strconv.Atoi(ctx.Query("pid"))
	var cate = model.Category{
		PID:  uint32(pid),
		Name: ctx.Query("name"),
	}

	p := page.GetNowPage(ctx)
	return cate.GetCateList(masterDB, p, pz)
}

//添加
func (c category) Create(ctx *gin.Context, spath string) error {
	var pid, _ = strconv.Atoi(ctx.PostForm("pid"))
	ul := uuid.NewV4()
	var n = time.Now().Unix()
	var cate = model.Category{
		PID:      uint32(pid),
		Name:     ctx.PostForm("name"),
		Ename:    ctx.PostForm("ename"),
		Icon:     ctx.PostForm("icon"),
		Remark:   ctx.PostForm("remark"),
		Ctime:    uint32(n),
		Utime:    uint32(n),
		SeoTitle: ctx.PostForm("seo_title"),
		SeoKwds:  ctx.PostForm("seo_keywords"),
		SeoDesc:  ctx.PostForm("seo_desc"),
		Thumb:    ctx.PostForm("seo_thumb"),
		Uuid:     ul.String(),
	}

	return cate.Create(masterDB,spath)
}

//修改
func (c category) Update() {

}

//删除
func (c category) Delete() {

}
