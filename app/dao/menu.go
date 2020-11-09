package dao

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/iwhot/zadmin/app/model"
	"github.com/iwhot/zadmin/system/page"
	"strconv"
	"time"
)

type menu struct {
}

var DefaultMenuDao = menu{}

//列表
func (m menu) GetMenuList(ctx *gin.Context, pz int) ([]*model.Menu, error) {
	var pid, _ = strconv.Atoi(ctx.Query("pid"))
	var me = model.Menu{
		Mname: ctx.Query("mname"),
		PID:   uint32(pid),
	}
	p := page.GetNowPage(ctx)
	return me.GetMenuList(masterDB, p, pz)
}

//添加
func (m menu) Create(ctx *gin.Context) error {
	var pid, _ = strconv.Atoi(ctx.PostForm("pid"))
	var Ordernum, _ = strconv.Atoi(ctx.PostForm("Ordernum"))
	var ty, _ = strconv.Atoi(ctx.PostForm("type"))

	var t = time.Now().Unix()
	var me = model.Menu{
		Url:       ctx.PostForm("url"),
		Condition: ctx.PostForm("condition"),
		PID:       uint32(pid),
		Ordernum:  uint32(Ordernum),
		Ctime:     uint32(t),
		Utime:     uint32(t),
		Icon:      ctx.PostForm("icon"),
		Remark:    ctx.PostForm("remark"),
		Mname:     ctx.PostForm("mname"),
		Type:      uint8(ty),
	}

	return me.Create(masterDB)
}

//统计
func (m menu) Count(ctx *gin.Context) int {
	var pid, _ = strconv.Atoi(ctx.PostForm("pid"))

	var me = model.Menu{
		PID:   uint32(pid),
		Mname: ctx.Query("mname"),
	}

	return me.Count(masterDB)
}

//编辑
func (m menu) Update(ctx *gin.Context) error {
	var id, _ = strconv.Atoi(ctx.PostForm("id"))
	var pid, _ = strconv.Atoi(ctx.PostForm("pid"))
	var ty, _ = strconv.Atoi(ctx.PostForm("type"))
	var Ordernum, _ = strconv.Atoi(ctx.PostForm("Ordernum"))

	if id == 0 {
		return errors.New("角色不存在")
	}
	var t = time.Now().Unix()
	var me = model.Menu{
		ID:        uint32(id),
		Url:       ctx.PostForm("url"),
		Condition: ctx.PostForm("condition"),
		PID:       uint32(pid),
		Ordernum:  uint32(Ordernum),
		Utime:     uint32(t),
		Icon:      ctx.PostForm("icon"),
		Remark:    ctx.PostForm("remark"),
		Mname:     ctx.PostForm("mname"),
		Type:      uint8(ty),
	}

	return me.Update(masterDB)
}

//删除
func (m menu) Delete(ctx *gin.Context) error {
	var id, _ = strconv.Atoi(ctx.Query("id"))
	if id == 0 {
		return errors.New("菜单不存在")
	}

	var me = model.Menu{
		ID: uint32(id),
	}

	return me.Delete(masterDB)
}

//获取一个菜单
func (m menu) GetOneMenu(ctx *gin.Context) (*model.Menu, error) {
	var id, _ = strconv.Atoi(ctx.Param("id"))

	var me = model.Menu{
		ID: uint32(id),
	}

	return me.GetOneMenu(masterDB)
}

//获取所有列表
func (m menu) GetAllMenuList() ([]*model.Menu, error) {
	var me = model.Menu{}
	return me.GetAllMenuList(masterDB)
}

//根据id获取一条记录
func (m menu) FindOneToID(id uint32) (*model.Menu, error) {
	var me = model.Menu{
		ID: id,
	}

	return me.GetOneMenu(masterDB)
}
