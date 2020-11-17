package model

import (
	"errors"
	page2 "github.com/iwhot/zadmin/system/page"
	"github.com/jinzhu/gorm"
)

type Links struct {
	ID     uint32 `json:"id"`
	Name   string `json:"name"`
	Url    string `json:"url"`
	Target string `json:"target"`
	State  uint8  `json:"state"`
	Icon   string `json:"icon"`
	Type   uint8  `json:"type"`
	Ctime  uint32 `json:"ctime"`
	Utime  uint32 `json:"utime"`
}

func (l Links) TableName() string {
	return Prefix + "links"
}

//添加
func (l Links) Create(DB *gorm.DB) error {
	return DB.Create(&l).Error
}

//更新
func (l Links) Update(DB *gorm.DB) error {
	return DB.Model(&l).Updates(l).Error
}

//删除
func (l Links) Delete(DB *gorm.DB) error {
	return DB.Delete(&l).Error
}

//获取一条
func (l Links) GetOneLinks(DB *gorm.DB) (*Links,error) {
	if l.ID  == 0{
		return nil,errors.New("链接不存在")
	}

	if err := DB.First(&l,l.ID).Error;err != nil{
		return nil,err
	}
	return &l,nil
}

//获取列表
func (l Links) GetLinksList(DB *gorm.DB, page, pageSize int) ([]*Links,error) {
	var linkSlice []*Links
	var offset = page2.GetOffset(page, pageSize)
	var mod *gorm.DB

	mod = DB.Model(&l)

	//where条件
	if l.Name != "" {
		mod = mod.Where("name like ?", "%"+l.Name+"%")
	}

	if err := mod.Where("is_del = ?", 0).Offset(offset).Limit(pageSize).Order("utime desc,id desc").Find(&linkSlice).Error; err != nil {
		return linkSlice, err
	}

	return linkSlice, nil
}