package model

import (
	"errors"
	"github.com/jinzhu/gorm"
)

type Files struct {
	ID    uint32  `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id"`
	Name  string  `gorm:"column:name;type:varchar(64);unique_index;not null;default:''" json:"name"`
	Url   string  `gorm:"column:url;type:varchar(255);unique_index;not null;default:''" json:"url"`
	Size  float64 `gorm:"column:size;type:double(20,3);not null;default:0" json:"size"`
	Type  uint8   `gorm:"column:type;type:tinyint(2);not null;default:0" json:"type"`
	Ctime uint32  `gorm:"column:ctime;type:int(10);not null;default:0" json:"ctime"`
}

func (f Files) TableName() string {
	return Prefix + "files"
}

//添加数据
func (f Files) Create(DB *gorm.DB) error {
	return DB.Create(&f).Error
}

//删除数据
func (f Files) Delete(DB *gorm.DB) error {
	if f.ID == 0 {
		return errors.New("文件不存在")
	}

	return DB.Delete(&f).Error
}
