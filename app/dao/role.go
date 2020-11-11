package dao

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/iwhot/zadmin/app/model"
	"github.com/iwhot/zadmin/system/page"
	"strconv"
	"time"
)

type role struct {
}

var DefaultRoleDao = role{}

//获取列表
func (r role) Find(ctx *gin.Context, pz int) ([]*model.Role, error) {
	var rol = model.Role{
		RoleName: ctx.Query("role_name"),
	}

	p := page.GetNowPage(ctx)
	return rol.GetRoleList(masterDB, p, pz)
}

//获取统计
func (r role) Count(ctx *gin.Context) int {
	var rol = model.Role{
		RoleName: ctx.Query("role_name"),
	}

	return rol.Count(masterDB)
}

//添加
func (r role) Create(ctx *gin.Context) error {
	var t = time.Now().Unix()
	var rol = model.Role{
		RoleName:  ctx.PostForm("role_name"),
		RoleCtime: uint32(t),
		RoleUtime: uint32(t),
		RoleDesc:  ctx.PostForm("role_desc"),
	}

	var menu = ctx.PostFormMap("menu")
	var me []uint32

	for _, v := range menu {
		id1, _ := strconv.Atoi(v)
		if id1 > 0 {
			me = append(me, uint32(id1))
		}
	}

	return rol.AddRole(masterDB,me)
}

//获取一条记录
func (r role) GetOneRoleInfo(ctx *gin.Context) (*model.Role, error) {
	var id, _ = strconv.Atoi(ctx.Param("id"))
	var rol = model.Role{
		ID: uint32(id),
	}
	return rol.FindOne(masterDB)
}

//编辑
func (r role) Update(ctx *gin.Context) error {
	var t = time.Now().Unix()

	var id, _ = strconv.Atoi(ctx.PostForm("id"))
	if id == 0 {
		return errors.New("角色不存在")
	}

	var rol = model.Role{
		ID:        uint32(id),
		RoleName:  ctx.PostForm("role_name"),
		RoleUtime: uint32(t),
		RoleDesc:  ctx.PostForm("role_desc"),
	}

	var menu = ctx.PostFormMap("menu")
	var me []uint32

	for _, v := range menu {
		id1, _ := strconv.Atoi(v)
		if id1 > 0 {
			me = append(me, uint32(id1))
		}
	}

	return rol.Update(masterDB,me)
}

//删除
func (r role) Delete(ctx *gin.Context) error {
	var id, _ = strconv.Atoi(ctx.Query("id"))
	if id == 0 {
		return errors.New("角色不存在")
	}

	var rol = model.Role{
		ID: uint32(id),
	}

	return rol.Delete(masterDB)
}

//获取所有记录
func (r role) GetAllRoleList() ([]*model.Role, error) {
	var rol = model.Role{}
	return rol.GetAllRoleList(masterDB)
}
