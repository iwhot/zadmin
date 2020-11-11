package model

import (
	"errors"
	"github.com/jinzhu/gorm"
)

type RoleMenu struct {
	RoleID uint32 `gorm:"column:role_id;" json:"role_id"`
	MenuID uint32 `gorm:"column:menu_id;" json:"menu_id"`
}

func (rm RoleMenu) TableName() string {
	return Prefix + "role_menu"
}

//创建
func (rm RoleMenu) CreateRoleMenu() {

}

//删除
func (rm RoleMenu) DeleteRoleMenu() {

}

//获取某角色菜单
func (rm RoleMenu) GetOneRoleMenuList(DB *gorm.DB) ([]*Menu, error) {
	if rm.RoleID == 0 {
		return nil, errors.New("角色不存在")
	}

	var m []*Menu

	if err := DB.Table(Prefix+"role_menu rm").Select("m.*").Joins("left join "+Prefix+"menu m on m.id=rm.menu_id").Where("rm.role_id=?", rm.RoleID).Find(&m).Error; err != nil {
		return nil, err
	}
	return m, nil
}
