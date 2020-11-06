package model

import (
	"errors"
	page2 "github.com/iwhot/zadmin/system/page"
	"github.com/jinzhu/gorm"
)

type Role struct {
	ID        uint32 `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id"`
	RoleName  string `gorm:"column:role_name;type:varchar(20);unique_index;not null;default:''" json:"role_name"`
	RoleCtime uint32 `gorm:"column:role_ctime;type:int(10);not null;default:0" json:"role_ctime"`
	RoleUtime uint32 `gorm:"column:role_utime;type:int(10);not null;default:0" json:"role_utime"`
	RoleDesc  string `gorm:"column:desc;type:varchar(1000);unique_index;not null;default:''" json:"role_desc"`
}

//获取角色列表
func (r Role) GetRoleList(DB *gorm.DB, page, pageSize int) ([]*Role, error) {
	var roles []*Role

	var offset = page2.GetOffset(page, pageSize)
	var mod *gorm.DB

	mod = DB.Model(&r)

	if r.RoleName != "" {
		mod = mod.Where("role_name like ?", "%"+r.RoleName+"%")
	}

	if r.ID != 0 {
		mod = mod.Where("id = ?", r.ID)
	}

	if err := mod.Where("is_del = ?", 0).Offset(offset).Limit(pageSize).Order("utime desc,id desc").Find(&roles).Error; err != nil {
		return nil, err
	}

	return roles, nil
}

//添加角色
func (r Role) AddRole(DB *gorm.DB) error {
	return DB.Create(&r).Error
}

//编辑角色
func (r Role) Update(DB *gorm.DB) error {
	return DB.Model(&r).Updates(r).Error
}

//获取一个角色
func (r Role) FindOne(DB *gorm.DB) (*Role, error) {
	if r.ID == 0 || r.RoleName == "" {
		return nil, errors.New("角色不存在")
	}

	roles, err := r.GetRoleList(DB, 0, 1)
	if err != nil {
		return nil, err
	}

	return roles[0], nil
}

//删除一个角色
func (r Role) Delete(DB *gorm.DB) error {
	tx := DB.Begin()
	//查找对应角色
	rs, err := r.FindOne(DB)
	if err != nil {
		tx.Rollback()
		return err
	}

	//如果角色已经绑定用户则不允许删除

	//删除角色
	if err = tx.Delete(&rs).Error; err != nil {
		tx.Rollback()
		return err
	}

	//删除角色下的权限

	tx.Commit()
	return nil
}

//统计
func (r Role) Count(DB *gorm.DB) int {
	var count int
	var mod *gorm.DB

	if r.RoleName != "" {
		mod = mod.Where("role_name like ?", "%"+r.RoleName+"%")
	}

	if err := mod.Count(&count).Error; err != nil {
		return 0
	}

	return count
}
