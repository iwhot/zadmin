package model

import (
	"errors"
	"github.com/jinzhu/gorm"
	"log"
)

type RoleMenu struct {
	RoleID uint32 `gorm:"column:role_id;" json:"role_id"`
	MenuID uint32 `gorm:"column:menu_id;" json:"menu_id"`
}

func (rm RoleMenu) TableName() string {
	return Prefix + "role_menu"
}

//创建
func (r Role) CreateRoleMenu() {

}


//删除
func (r Role) DeleteRoleMenu() {

}

//获取某角色菜单
func (r Role) GetOneRoleMenuList(DB *gorm.DB) (*Role, error) {
	if r.ID == 0 {
		return nil, errors.New("角色不存在")
	}

	var m []Menu
	if err := DB.Model(&r).Related(&m, "RoleMenu").Error; err != nil {
		return nil, err
	}
	log.Println(r)
	log.Println(m)
	return &r, nil
}
