package validate

type User struct {
	Username string `form:"username" binding:"required,max=30,min=3"`
	Email    string `form:"email" binding:"required,email,max=255,min=3"`
	Password string `form:"password" binding:"omitempty,gt=5"` //密码要么不传要传就大于6
}


type UserEdit struct {
	Email    string `form:"email" binding:"required,email,max=255,min=3"`
}