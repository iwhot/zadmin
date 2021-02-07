package middleware

import "github.com/gin-gonic/gin"

//todo isLogin：0后台不是在登录页  1后台是在登录页
func CheckLogin(isLogin int8) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if isLogin == 2 {
			ctx.Next()
			return
		}
		ctx.Next()
	}
}
