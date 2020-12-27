package model

import (
	page2 "github.com/iwhot/zadmin/system/page"
	"github.com/jinzhu/gorm"
)

type AdPosition struct {
	ID     uint32 `gorm:"column:id;primary_key;AUTO_INCREMENT"json:"id"`
	Name   string `gorm:"column:name;" json:"name"`
	Width  uint32 `gorm:"column:width;"json:"width"`
	Height uint32 `gorm:"column:height;"json:"height"`
	Type   uint8  `gorm:"column:type;"json:"type"`
	Ctime  uint32 `gorm:"column:ctime;"json:"ctime"`
	Utime  uint32 `gorm:"column:utime;"json:"utime"`
	Desc   string `gorm:"column:desc;"json:"desc"`
	Style  string `gorm:"column:style;"json:"style"`
}

func (a AdPosition) TableName() string {
	return Prefix + "ad_position"
}

//添加
func (a AdPosition) Create(DB *gorm.DB) error {
	return DB.Create(&a).Error
}

//修改
func (a AdPosition) Update(DB *gorm.DB) error {
	return DB.Model(&a).Updates(a).Error
}

//删除
func (a AdPosition) Delete(DB *gorm.DB) error {
	return DB.Where("id=?", a.ID).Delete(a).Error
}

//获取一条记录
func (a AdPosition) GetOneAdPosition(DB *gorm.DB) (*AdPosition, error) {
	var as = []*AdPosition{}
	if err := DB.Model(&a).Where("id", a.ID).Offset(0).Limit(1).First(&as).Error; err != nil {
		return nil, err
	}

	return as[0], nil
}

//获取列表
func (a AdPosition) GetAdPositionList(DB *gorm.DB, page, pageSize int) ([]*AdPosition, error) {
	var as = []*AdPosition{}
	var offset = page2.GetOffset(page, pageSize)

	if err := DB.Model(&a).Offset(offset).Limit(pageSize).Order("ctime desc,id desc").Find(&as).Error; err != nil {
		return nil, err
	}
	return as, nil
}
