package dao

import (
	"github.com/gin-gonic/gin"
	"github.com/iwhot/zadmin/app/model"
	"github.com/iwhot/zadmin/system/page"
)

type user struct {
}

//var DefaultUserDao = user{}

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
