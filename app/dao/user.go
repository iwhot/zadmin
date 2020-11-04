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
func (u user) GetUserList(ctx *gin.Context) ([]*model.User, error) {
	var usr = model.User{
		Username: ctx.Query("username"),
		Email:    ctx.Query("email"),
	}
	p := page.GetNowPage(ctx)
	pz := page.GetPageSize(ctx)
	return usr.GetUserList(masterDB, p, pz)
}

//添加用户
func (u user) AddUser(ctx *gin.Context) error {
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
	}

	return usr.AddUser(masterDB)
}
