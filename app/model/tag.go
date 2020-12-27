package model

import "github.com/jinzhu/gorm"

type Tag struct {
	ID      uint32 `json:"id"`
	TagName string `json:"tag_name"`
}

func (t Tag) TableName() string {
	return Prefix + "tag"
}

//创建
func (t Tag) Create(db *gorm.DB) uint32 {
	if err := db.Create(&t).Error;err != nil{
		return 0
	}

	return t.ID
}

//修改

//删除

//获取一条记录

//获取列表
