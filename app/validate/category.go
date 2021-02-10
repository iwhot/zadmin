package validate

type Category struct {
	Name string `form:"name" binding:"required,max=20,min=2"`
	Ename string `form:"ename" binding:"required,max=20,min=3"`
}