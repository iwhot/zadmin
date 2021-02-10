package dao

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/iwhot/zadmin/app/controller"
	"github.com/iwhot/zadmin/app/model"
	"github.com/satori/go.uuid"
	"log"
	"strconv"
	"strings"
	"time"
)

type category struct {
}

var DefaultCategoryDao = category{}

//获取列表
func (c category) GetCateList(ctx *gin.Context) ([]*model.Category, error) {
	var pid, _ = strconv.Atoi(ctx.Query("pid"))
	var cate = model.Category{
		PID:  uint32(pid),
		Name: ctx.Query("name"),
	}

	return cate.GetCateList(masterDB,true)
}

//获取列表
func (c category) GetCateListByPid(ctx *gin.Context) ([]*model.Category, error) {
	var cate = model.Category{
		PID:  0,
	}
	return cate.GetCateList(masterDB,false)
}

//添加
func (c category) Create(ctx *gin.Context) error {
	var pid, _ = strconv.Atoi(ctx.PostForm("pid"))
	var state, _ = strconv.Atoi(ctx.PostForm("state"))
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
		State:    uint8(state),
	}

	return cate.Create(masterDB)
}

//修改
func (c category) Update(ctx *gin.Context) error {
	var pid, _ = strconv.Atoi(ctx.PostForm("pid"))
	var id, _ = strconv.Atoi(ctx.PostForm("id"))
	var state, _ = strconv.Atoi(ctx.PostForm("state"))
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
		State:    uint8(state),
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

//获取一个分类
func (c category) GetOneCategory(ctx *gin.Context) (*model.Category,error) {
	var id, _ = strconv.Atoi(strings.TrimRight(ctx.Param("id" + controller.BACKENDFIX),controller.BACKENDFIX))
	var cate = model.Category{
		ID:       uint32(id),
	}
	log.Println(id)
	return cate.GetOneCategory(masterDB)
}
