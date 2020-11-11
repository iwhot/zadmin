package dao

import (
	"github.com/gin-gonic/gin"
	"github.com/iwhot/zadmin/app/model"
	"strconv"
)

type roleMenu struct {
}

var DefaultRoleMenuDao = roleMenu{}

func (rm roleMenu) GetOneRoleMenuList(ctx *gin.Context) ([]*model.Menu, error) {
	var roleID, _ = strconv.Atoi(ctx.Query("id"))

	var rms = model.RoleMenu{
		RoleID: uint32(roleID),
	}

	return rms.GetOneRoleMenuList(masterDB)
}
