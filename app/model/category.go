package model

import "github.com/jinzhu/gorm"

type Category struct {
	ID       uint32 `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id"`
	PID      uint32 `gorm:"column:pid;" json:"pid"`
	Name     string `gorm:"column:name;" json:"name"`
	Ename    string `gorm:"column:ename;" json:"ename"`
	Icon     string `gorm:"column:icon;" json:"icon"`
	Remark   string `gorm:"column:remark;" json:"remark"`
	Ctime    uint32 `gorm:"column:ctime;" json:"ctime"`
	Utime    uint32 `gorm:"column:utime;" json:"utime"`
	SeoTitle string `gorm:"column:seo_title;" json:"seo_title"`
	SeoKwds  string `gorm:"column:seo_kwds;" json:"seo_kwds"`
	SeoDesc  string `gorm:"column:seo_desc;" json:"seo_desc"`
	Thumb    string `gorm:"column:thumb;" json:"thumb"`
	Uuid     string `gorm:"column:uuid;" json:"uuid"`
}

func (c Category) TableName() string {
	return Prefix + "category"
}

//添加
func (c Category) Create(DB *gorm.DB) error {
	return DB.Create(&c).Error
}

//修改
func (c Category) Update(DB *gorm.DB) {

}

//获取一个分类
func (c Category) GetOneCategory(DB *gorm.DB) {

}

//删除
func (c Category) Delete(DB *gorm.DB) {

}

//获取分类列表==>后台
func (c Category) GetCateList(DB *gorm.DB) {

}

//================= todo 前台部分待定...
