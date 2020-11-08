package validate

type Menu struct {
	Mmenu string `form:"mmenu" binding:"required,max=30,min=3"`
	Url    string `form:"url" binding:"required,email,max=255,min=3"`
}
