package model

import "github.com/jinzhu/gorm"

type Ad struct {
	ID uint32 `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id"`
	Title string `gorm:"column:title;" json:"title"`
	Url string `gorm:"column:url;" json:"url"`
	PID uint32 `gorm:"column:pid;" json:"pid"`
	Ctime uint32 `gorm:"column:ctime;" json:"ctime"`
	Utime uint32 `gorm:"column:utime;" json:"utime"`
}

func (a Ad) TableName() string {
	return Prefix + "ad"
}

//添加
func (a Ad) Create(DB *gorm.DB) {

}

//修改
func (a Ad) Update(DB *gorm.DB) {

}

//删除
func (a Ad) Delete(DB *gorm.DB) {

}

//获取列表
func (a Ad) GetAdList(DB *gorm.DB) {

}

//获取一条记录
func (a Ad) GetOneAdList(DB *gorm.DB) {

}