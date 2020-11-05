package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
	"net/http"
	"strings"
)

const PAGESIZE = 20 //一页条数

type Controller struct {
}

//渲染页面
func (this *Controller) Render(ctx *gin.Context, tmpl string, data interface{}) {
	if data == nil {
		data = gin.H{}
	}
	ctx.HTML(http.StatusOK, tmpl, data)
}

//立即跳转
func (this *Controller) Redirect(ctx *gin.Context, url string) {
	ctx.Redirect(http.StatusMovedPermanently, url)
}

//返回json
func (this *Controller) JSON(ctx *gin.Context, data interface{}) {
	if data == nil {
		data = gin.H{}
	}
	ctx.JSON(http.StatusOK, data)
}

//验证器
func (this *Controller) Validate(ctx *gin.Context, valid interface{}) error {
	//todo ================= 此部分可放中间件 start
	uni := ut.New(zh.New(), en.New())
	trans, _ := uni.GetTranslator("zh")
	v, ok := binding.Validator.Engine().(*validator.Validate)
	if ok {
		//验证器翻译
		_ = zh_translations.RegisterDefaultTranslations(v, trans)
		//自定义验证方法
	}
	//todo ================= 此部分可放中间件 end
	return checkValidator(ctx, valid, trans)
}

//自定义验证方法
func checkValidator(ctx *gin.Context, valid interface{}, trans ut.Translator) error {
	err := ctx.ShouldBind(valid)

	if err != nil {
		//解析err
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			return nil
		}

		var errList []string
		for _, e := range errs {
			errList = append(errList, e.Translate(trans))
		}

		return errors.New(strings.Join(errList, "|"))
	}

	return nil
}

//404页面
func (this *Controller) NotFound(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"code": 404,
		"msg":  "404 not found !0.0",
	})
}
