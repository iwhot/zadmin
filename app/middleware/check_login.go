package middleware

import "github.com/gin-gonic/gin"

//0不是在登录页  1是在登录页  2直接全面通行
func CheckLogin(isLogin int8) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if isLogin == 2 {
			ctx.Next()
			return
		}
		ctx.Next()
	}
}
