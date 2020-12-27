package model

import (
	"errors"
	page2 "github.com/iwhot/zadmin/system/page"
	"github.com/jinzhu/gorm"
)

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
	tx := DB.Begin()
	if err := DB.Create(&c).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

//修改
func (c Category) Update(DB *gorm.DB) error {
	tx := DB.Begin()
	if err := DB.Model(&c).Updates(c).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

//获取一个分类
func (c Category) GetOneCategory(DB *gorm.DB) (*Category, error) {
	if c.ID == 0 {
		return nil, errors.New("分类不存在")
	}
	var cate []*Category
	if err := DB.Model(&c).Where("id=?", c.ID).Offset(0).Limit(1).First(&cate).Error; err != nil {
		return nil, err
	}
	return cate[0], nil
}

//删除
func (c Category) Delete(DB *gorm.DB) error {
	if c.ID == 0 {
		return errors.New("分类不存在")
	}

	tx := DB.Begin()
	//有子分类不可删除
	var cate *Category
	err := tx.Model(&c).Where("pid=?", c.ID).Offset(0).Limit(1).First(&cate).Error
	if err == nil && cate != nil {
		tx.Rollback()
		return errors.New("分类下有子分类不可删除")
	}

	//分类下有文章不可删除
	var art *Article
	var a = Article{}
	err = tx.Model(&a).Where("category_id=?", c.ID).Offset(0).Limit(1).First(&art).Error
	if err == nil && art != nil {
		tx.Rollback()
		return errors.New("分类下有文章不可删除")
	}

	if err2 := tx.Where("id=?",c.ID).Delete(Category{}).Error;err2 != nil{
		tx.Rollback()
		return err2
	}

	tx.Commit()
	return nil
}

//获取分类列表==>后台
func (c Category) GetCateList(DB *gorm.DB, page, pageSize int) ([]*Category, error) {
	var cate = []*Category{}
	var offset = page2.GetOffset(page, pageSize)
	var mod *gorm.DB

	mod = DB.Model(&c)
	//where条件
	if c.Name != "" {
		mod = mod.Where("username like ?", "%"+c.Name+"%")
	}

	if c.PID != 0 {
		mod = mod.Where("pid=?", c.PID)
	}

	if err := mod.Where("is_del = ?", 0).Offset(offset).Limit(pageSize).Order("utime desc,id desc").Find(&cate).Error; err != nil {
		return cate, err
	}

	return cate, nil
}

//================= todo 前台部分待定...
