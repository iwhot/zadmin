package backend

import (
	"github.com/gin-gonic/gin"
	"github.com/iwhot/zadmin/app/controller"
)

type Login struct {
	controller.Controller
}

func (this *Login) NewRouter(g *gin.RouterGroup) {

}
