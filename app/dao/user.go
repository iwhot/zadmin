package dao

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/iwhot/zadmin/app/model"
	"github.com/iwhot/zadmin/system/common"
	"github.com/iwhot/zadmin/system/page"
	"strconv"
	"time"
)

type user struct {
}

var DefaultUserDao = user{}

//获取用户列表
func (u user) GetUserList(ctx *gin.Context, pz int) ([]*model.User, error) {
	var usr = model.User{
		Username: ctx.Query("username"),
		Email:    ctx.Query("email"),
		Rname:    ctx.Query("rname"),
	}
	p := page.GetNowPage(ctx)
	return usr.GetUserList(masterDB, p, pz)
}

//添加用户
func (u user) AddUser(ctx *gin.Context, spath string) error {
	var role_id, _ = strconv.Atoi(ctx.PostForm("role_id"))
	if role_id == 0 {
		return errors.New("角色必须选择")
	}
	//手工验证密码
	var passwd = ctx.PostForm("password")
	var repasswd = ctx.PostForm("repassword")

	if len(passwd) < 6 {
		return errors.New("密码长度必须大于6位")
	}

	if passwd != repasswd {
		return errors.New("密码和确认密码不一致")
	}
	//创建盐值
	var salt = common.CreateSalt(10)
	//生成密码
	var pwd = common.Md5V("zadmin_" + passwd + salt + common.Md5V(salt) + "_end")

	var state, _ = strconv.Atoi(ctx.PostForm("state"))
	var t = time.Now().Unix()

	var usr = model.User{
		Username: ctx.PostForm("username"),
		Rname:    ctx.PostForm("rname"),
		Password: pwd,
		State:    uint8(state),
		Avatar:   ctx.PostForm("avatar"),
		Desc:     ctx.PostForm("desc"),
		Email:    ctx.PostForm("email"),
		Ctime:    uint32(t),
		Utime:    uint32(t),
		Ltime:    uint32(t),
		Salt:     salt,
		RoleID:   uint32(role_id),
	}

	return usr.AddUser(masterDB, spath)
}

//获取一条用户记录
func (u user) GetOneUserInfo(ctx *gin.Context) (*model.User, error) {
	var id, _ = strconv.Atoi(ctx.Param("id"))
	var usr = model.User{
		ID: uint32(id),
	}
	return usr.GetOneUserInfo(masterDB)
}

//统计数量
func (u user) Count(ctx *gin.Context) int {
	var usr = model.User{
		Username: ctx.Query("username"),
		Email:    ctx.Query("email"),
		Rname:    ctx.Query("rname"),
	}

	return usr.Count(masterDB)
}

//编辑用户
func (u user) Update(ctx *gin.Context, spath string) error {
	var role_id, _ = strconv.Atoi(ctx.PostForm("role_id"))
	if role_id == 0 {
		return errors.New("角色必须选择")
	}
	//手工验证密码
	var passwd = ctx.PostForm("password")
	var repasswd = ctx.PostForm("repassword")
	var id, _ = strconv.Atoi(ctx.PostForm("id"))
	if id == 0 {
		return errors.New("用户不存在")
	}
	var t = time.Now().Unix()

	if passwd != "" {
		if len(passwd) < 6 {
			return errors.New("密码长度必须大于6位")
		}

		if passwd != repasswd {
			return errors.New("密码和确认密码不一致")
		}
		//创建盐值
		var salt = common.CreateSalt(10)
		//生成密码
		var pwd = common.Md5V("zadmin_" + passwd + salt + common.Md5V(salt) + "_end")

		var usr = model.User{
			ID:       uint32(id),
			Rname:    ctx.PostForm("rname"),
			Avatar:   ctx.PostForm("avatar"),
			Desc:     ctx.PostForm("desc"),
			Email:    ctx.PostForm("email"),
			Utime:    uint32(t),
			Salt:     salt,
			Password: pwd,
			RoleID:   uint32(role_id),
		}

		return usr.UpdateUser(masterDB, spath)

	} else {
		var usr = model.User{
			ID:     uint32(id),
			Rname:  ctx.PostForm("rname"),
			Avatar: ctx.PostForm("avatar"),
			Desc:   ctx.PostForm("desc"),
			Email:  ctx.PostForm("email"),
			Utime:  uint32(t),
			RoleID: uint32(role_id),
		}

		return usr.UpdateUser(masterDB, spath)
	}

}

//删除用户
func (u user) Delete(ctx *gin.Context, spath string) error {
	var id, _ = strconv.Atoi(ctx.Query("id"))
	if id == 0 {
		return errors.New("用户不存在")
	}

	var usr = model.User{
		ID: uint32(id),
	}

	return usr.DeleteUser(masterDB, spath)
}
