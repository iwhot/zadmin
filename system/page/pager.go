package page

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

//获取当前页
func GetNowPage(ctx *gin.Context) int {
	var p, _ = strconv.Atoi(ctx.Query("page"))
	if p < 1 {
		return 1
	}

	return p
}

//获取一页大小
func GetPageSize(ctx *gin.Context) int {
	var pz, _ = strconv.Atoi(ctx.Query("page_size"))
	if pz < 1 {
		return 1
	}

	return pz
}

//获取偏移量
func GetOffset(page, pageSize int) int {
	if page < 1 {
		return 0
	}

	return (page - 1) * pageSize
}
