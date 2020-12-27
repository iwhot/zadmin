package dao

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/iwhot/zadmin/app/model"
	"github.com/iwhot/zadmin/system/page"
	"github.com/satori/go.uuid"
	"strconv"
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
func (c category) Create(ctx *gin.Context) error {
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

	return cate.Create(masterDB)
}

//修改
func (c category) Update(ctx *gin.Context) error {
	var pid, _ = strconv.Atoi(ctx.PostForm("pid"))
	var id, _ = strconv.Atoi(ctx.PostForm("id"))
	var n = time.Now().Unix()
	var cate = model.Category{
		ID:       uint32(id),
		PID:      uint32(pid),
		Name:     ctx.PostForm("name"),
		Ename:    ctx.PostForm("ename"),
		Icon:     ctx.PostForm("icon"),
		Remark:   ctx.PostForm("remark"),
		Utime:    uint32(n),
		SeoTitle: ctx.PostForm("seo_title"),
		SeoKwds:  ctx.PostForm("seo_keywords"),
		SeoDesc:  ctx.PostForm("seo_desc"),
		Thumb:    ctx.PostForm("seo_thumb"),
	}

	return cate.Update(masterDB)
}

//删除
func (c category) Delete(ctx *gin.Context) error {
	var id, _ = strconv.Atoi(ctx.Query("id"))
	if id == 0 {
		return errors.New("用户不存在")
	}

	var cate = model.Category{
		ID: uint32(id),
	}
	return cate.Delete(masterDB)
}
