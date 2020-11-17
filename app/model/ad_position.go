package model

import "github.com/jinzhu/gorm"

type AdPosition struct {
	ID uint32 `gorm:"column:id;primary_key;AUTO_INCREMENT"json:"id"`
	Name string `gorm:"column:name;" json:"name"`
	Width uint32 `gorm:"column:width;"json:"width"`
	Height uint32 `gorm:"column:height;"json:"height"`
	Type uint8 `gorm:"column:type;"json:"type"`
	Ctime uint32 `gorm:"column:ctime;"json:"ctime"`
	Utime uint32 `gorm:"column:utime;"json:"utime"`
	Desc string `gorm:"column:desc;"json:"desc"`
	Style string `gorm:"column:style;"json:"style"`
}

func (a AdPosition) TableName() string {
	return Prefix + "ad_position"
}

//添加
func (a AdPosition) Create(DB *gorm.DB) {

}

//修改
func (a AdPosition) Update(DB *gorm.DB) {

}

//删除
func (a AdPosition) Delete(DB *gorm.DB) {

}

//获取一条记录
func (a AdPosition) GetOneAdPosition(DB *gorm.DB) {

}

//获取列表
func (a AdPosition) GetAdPositionList(DB *gorm.DB) {

}