package model

import (
	"errors"
	page2 "github.com/iwhot/zadmin/system/page"
	"github.com/jinzhu/gorm"
	"os"
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
func (f Files) Delete(DB *gorm.DB, spath string) error {
	if f.ID == 0 {
		return errors.New("文件不存在")
	}
	//删除文件
	_ = os.Remove(spath + f.Url)
	return DB.Delete(&f).Error
}

//获取列表
func (f Files) Find(DB *gorm.DB, page, pageSize int) ([]*Files, error) {
	var fileSlice []*Files
	var offset = page2.GetOffset(page, pageSize)
	var mod *gorm.DB

	mod = DB.Model(&f)
	if f.ID != 0 {
		mod = mod.Where("id = ?", f.ID)
	}

	if f.Url != "" {
		mod = mod.Where("url = ?", f.Url)
	}

	if err := mod.Offset(offset).Limit(pageSize).Order("ctime desc,id desc").Find(&fileSlice).Error; err != nil {
		return fileSlice, err
	}

	return fileSlice, nil
}

//获取一条记录
func (f Files) FindOne(DB *gorm.DB) (*Files, error) {
	if f.ID == 0 && f.Url == "" {
		return nil, errors.New("文件不存在")
	}

	fls, err := f.Find(DB, 0, 1)
	if err != nil {
		return nil, err
	}

	return fls[0], nil
}

//更新记录
func (f Files) Update(DB *gorm.DB) error {
	return DB.Model(&f).Updates(f).Error
}
