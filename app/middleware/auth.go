package middleware

import "github.com/gin-gonic/gin"

//后台权限
func BackendAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//todo ...
		ctx.Next()
	}
}
