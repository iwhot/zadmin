package dao

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/iwhot/zadmin/app/model"
	"github.com/iwhot/zadmin/system/page"
	"strconv"
)

type files struct {
}

var DefaultFilesDao = files{}

func (f files) Create(fls model.Files) error {
	return fls.Create(masterDB)
}

//获取列表
func (f files) Find(ctx *gin.Context, pz int) ([]*model.Files, error) {
	var file = model.Files{}

	p := page.GetNowPage(ctx)
	return file.GetFileList(masterDB, p, pz)
}

//统计文件数量
func (f files) Count(ctx *gin.Context) int {
	return (model.Files{}).Count(masterDB)
}

//删除文件
func (f files) Delete(ctx *gin.Context) error {
	var id, _ = strconv.Atoi(ctx.Query("id"))
	if id == 0 {
		return errors.New("文件不存在")
	}

	var file = model.Files{
		ID: uint32(id),
	}

	return file.Delete(masterDB)
}

