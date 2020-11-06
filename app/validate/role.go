package validate

type Role struct {
	RoleName string `form:"role_name" binding:"required,max=20,min=3"`
}
