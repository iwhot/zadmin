package model

import (
	"errors"
	page2 "github.com/iwhot/zadmin/system/page"
	"github.com/jinzhu/gorm"
)

type Menu struct {
	ID        uint32 `gorm:"column:id;type:smallint(6);primary_key;AUTO_INCREMENT" json:"id"`
	Url       string `gorm:"column:url;type:varchar(255);unique_index;not null;default:''" json:"url"`
	Condition string `gorm:"column:condition;type:varchar(255);not null;default:''" json:"condition"`
	PID       uint32 `gorm:"column:pid;type:smallint(6);not null;default:0" json:"pid"`
	Ordernum  uint32 `gorm:"column:ordernum;type:int(10);not null;default:0" json:"ordernum"`
	Ctime     uint32 `gorm:"column:ctime;type:int(10);not null;default:0" json:"ctime"`
	Utime     uint32 `gorm:"column:utime;type:int(10);not null;default:0" json:"utime"`
	Icon      string `gorm:"column:icon;type:varchar(255);unique_index;not null;default:''" json:"icon"`
	Remark    string `gorm:"column:remark;type:varchar(255);unique_index;not null;default:''" json:"remark"`
	Mname     string `gorm:"column:mname;type:varchar(10);unique_index;not null;default:''" json:"mname"`
	Type      uint8  `gorm:"column:type;type:tinyint(2);not null;default:0" json:"type"`
}

func (u Menu) TableName() string {
	return Prefix + "menu"
}

//获取列表
func (m Menu) GetMenuList(DB *gorm.DB, page, pageSize int) ([]*Menu, error) {
	var menus []*Menu

	var offset = page2.GetOffset(page, pageSize)
	var mod *gorm.DB

	mod = DB.Model(&m)
	if m.ID != 0 {
		mod = mod.Where("id = ?", m.ID)
	} else {
		if m.Mname != "" {
			mod = mod.Where("mname like ?", "%"+m.Mname+"%")
		}

		if m.PID != 0 {
			mod = mod.Where("pid = ?", m.PID)
		}
	}

	if err := mod.Offset(offset).Limit(pageSize).Order("ordernum desc,id desc").Find(&menus).Error; err != nil {
		return nil, err
	}

	return menus, nil

}

//添加
func (m Menu) Create(DB *gorm.DB) error {
	return DB.Create(&m).Error
}

//修改
func (m Menu) Update(DB *gorm.DB) error {
	return DB.Model(&m).Updates(m).Error
}

//删除
func (m Menu) Delete(DB *gorm.DB) error {
	if m.ID == 0 {
		return errors.New("菜单不存在")
	}

	return DB.Delete(&m).Error
}

//获取一个菜单
func (m Menu) GetOneMenu(DB *gorm.DB) (*Menu, error) {
	if m.ID == 0 {
		return nil, errors.New("菜单不存在")
	}
	mu, err := m.GetMenuList(DB, 0, 1)
	if err != nil {
		return nil, err
	}
	return mu[0], nil
}

//统计
func (m Menu) Count(DB *gorm.DB) int {
	var count int
	var mod *gorm.DB
	mod = DB.Model(&m)
	if m.Mname != "" {
		mod = mod.Where("mname like ?", "%"+m.Mname+"%")
	}

	if m.PID != 0 {
		mod = mod.Where("pid = ?", m.PID)
	}

	if err := mod.Count(&count).Error; err != nil {
		return 0
	}

	return count
}

//获取所有菜单
func (m Menu) GetMenuByPid(DB *gorm.DB) ([]*Menu, error) {
	var menus []*Menu

	if err := DB.Model(&m).Where("pid = ?", m.PID).Order("pid asc").Find(&menus).Error; err != nil {
		return nil, err
	}

	return menus, nil
}
