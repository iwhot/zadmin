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
	RoleDesc  string `gorm:"column:role_desc;type:varchar(1000);unique_index;not null;default:''" json:"role_desc"`
	//RoleMenu []Menu `gorm:"many2many:zs_role_menu" json:"role_menu"`//多对多(=@__@=)
}

func (u Role) TableName() string {
	return Prefix + "role"
}

//获取角色列表
func (r Role) GetRoleList(DB *gorm.DB, page, pageSize int) ([]*Role, error) {
	var roles = []*Role{}

	var offset = page2.GetOffset(page, pageSize)
	var mod *gorm.DB

	mod = DB.Model(&r)

	if r.ID != 0 {
		mod = mod.Where("id = ?", r.ID)
	} else {
		if r.RoleName != "" {
			mod = mod.Where("role_name like ?", "%"+r.RoleName+"%")
		}
	}

	if err := mod.Offset(offset).Limit(pageSize).Order("role_utime desc,id desc").Find(&roles).Error; err != nil {
		return nil, err
	}

	return roles, nil
}

//添加角色
func (r Role) AddRole(DB *gorm.DB, me []uint32) error {
	tx := DB.Begin()

	//角色添加
	if err := tx.Create(&r).Error; err != nil {
		tx.Rollback()
		return err
	}

	//删除角色相关权限
	if err := tx.Where("role_id=?", r.ID).Delete(RoleMenu{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	//角色权限添加
	if len(me) < 1 {
		tx.Commit()
		return nil
	}

	for _, v := range me {
		var rm = RoleMenu{
			RoleID: r.ID,
			MenuID: v,
		}

		if err := tx.Create(&rm).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	tx.Commit()
	return nil
}

//编辑角色
func (r Role) Update(DB *gorm.DB, me []uint32) error {
	tx := DB.Begin()

	if err := DB.Model(&r).Updates(r).Error; err != nil {
		tx.Rollback()
		return err
	}

	//删除角色相关权限
	if err := tx.Where("role_id=?", r.ID).Delete(RoleMenu{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	if len(me) < 1 {
		tx.Commit()
		return nil
	}

	for _, v := range me {
		var rm = RoleMenu{
			RoleID: r.ID,
			MenuID: v,
		}

		if err := tx.Create(&rm).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	tx.Commit()
	return nil
}

//获取一个角色
func (r Role) FindOne(DB *gorm.DB) (*Role, error) {
	if r.ID == 0 && r.RoleName == "" {
		return nil, errors.New("角色不存在")
	}

	roles, err := r.GetRoleList(DB, 0, 1)
	if err != nil {
		return &Role{}, err
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

	//删除角色
	if err = tx.Delete(&rs).Error; err != nil {
		tx.Rollback()
		return err
	}

	//删除角色下的权限
	if err := DB.Where("role_id=?", r.ID).Delete(RoleMenu{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

//统计
func (r Role) Count(DB *gorm.DB) int {
	var count int
	var mod *gorm.DB
	mod = DB.Model(&r)
	if r.RoleName != "" {
		mod = mod.Where("role_name like ?", "%"+r.RoleName+"%")
	}

	if err := mod.Count(&count).Error; err != nil {
		return 0
	}

	return count
}

//获取所有角色
func (r Role) GetAllRoleList(DB *gorm.DB) ([]*Role, error) {
	var roles []*Role
	if err := DB.Model(&r).Find(&roles).Error; err != nil {
		return roles, err
	}

	return roles, nil
}
