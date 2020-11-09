package validate

type Menu struct {
	Mname string `form:"mname" binding:"required,max=30,min=2"`
	Url    string `form:"url" binding:"required,max=255,min=3"`
}
