package model

type RoleMenu struct {
	RoleID uint32 `gorm:"column:role_id;" json:"role_id"`
	MenuID uint32 `gorm:"column:menu_id;" json:"menu_id"`
}

func (rm RoleMenu) TableName() string {
	return Prefix + "role_menu"
}

//创建
func (rm RoleMenu) Create() {

}

//更新
func (rm RoleMenu) Update() {

}

//删除
func (rm RoleMenu) Delete() {

}

//获取某角色菜单
func (rm RoleMenu) GetOneRoleMenuList() {

}
