package model

import "github.com/jinzhu/gorm"

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
func (l Links) GetOneLinks(DB *gorm.DB) {

}

//获取列表
func (l Links) GetLinksList(DB *gorm.DB) {

}