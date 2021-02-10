package model

import (
	"errors"
	"github.com/iwhot/zadmin/app/controller"
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
	State uint8   `gorm:"column:state;type:tinyint(2);not null;default:0" json:"state"`
	Title string  `gorm:"column:title;" json:"title"`
	Desc  string  `gorm:"column:desc;" json:"desc"`
	Utime uint32  `gorm:"column:utime;type:int(10);not null;default:0" json:"utime"`
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
	//删除文件
	_ = os.Remove(controller.STORAGEPATH + f.Url)
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

//获取文件列表
func (f Files) GetFileList(DB *gorm.DB, page, pageSize int) ([]*Files, error) {
	var file = []*Files{}
	var offset = page2.GetOffset(page, pageSize)

	if err := DB.Model(&f).Offset(offset).Limit(pageSize).Order("role_utime desc,id desc").Find(&file).Error; err != nil {
		return nil, err
	}

	return file, nil
}

//获取统计
func (f Files)  Count(DB *gorm.DB) int {
	var count int

	if err := DB.Model(&f).Count(&count).Error; err != nil {
		return 0
	}

	return count
}

//查找某个文件
func (f Files) SearchFileByAddress(DB *gorm.DB) uint32 {
	if f.Url == ""{
		return 0
	}

	var file = []*Files{}
	if err := DB.Model(&f).Where("url=?",f.Url).Offset(0).Limit(1).Find(&file).Error;err != nil{
		return 0
	}

	return file[0].ID
}