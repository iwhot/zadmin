package model

import (
	page2 "github.com/iwhot/zadmin/system/page"
	"github.com/jinzhu/gorm"
)

type Ad struct {
	ID    uint32 `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id"`
	Title string `gorm:"column:title;" json:"title"`
	Url   string `gorm:"column:url;" json:"url"`
	PID   uint32 `gorm:"column:pid;" json:"pid"`
	Ctime uint32 `gorm:"column:ctime;" json:"ctime"`
	Utime uint32 `gorm:"column:utime;" json:"utime"`
}

func (a Ad) TableName() string {
	return Prefix + "ad"
}

//添加
func (a Ad) Create(DB *gorm.DB) error {
	return DB.Create(&a).Error
}

//修改
func (a Ad) Update(DB *gorm.DB) error {
	return DB.Model(&a).Updates(a).Error
}

//删除
func (a Ad) Delete(DB *gorm.DB) error {
	return DB.Where("id=?", a.ID).Delete(a).Error
}

//获取列表
func (a Ad) GetAdList(DB *gorm.DB, page, pageSize int) ([]*Ad, error) {
	var as = []*Ad{}
	var offset = page2.GetOffset(page, pageSize)

	if err := DB.Model(&a).Offset(offset).Limit(pageSize).Order("ctime desc,id desc").Find(&as).Error; err != nil {
		return nil, err
	}
	return as, nil
}

//获取一条记录
func (a Ad) GetOneAd(DB *gorm.DB) (*Ad,error) {
	var as = []*Ad{}
	if err := DB.Where("id=?",a.ID).Offset(0).Limit(1).First(&as).Error; err != nil {
		return nil, err
	}
	return as[0], nil
}
