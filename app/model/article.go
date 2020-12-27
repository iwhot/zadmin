package model

import (
	"errors"
	"github.com/jinzhu/gorm"
)

type Article struct {
	ID         uint32 `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id"`
	Title      string `gorm:"column:title;" json:"title"`
	ShortTitle string `gorm:"column:short_title;" json:"short_title"`
	CategoryID uint32 `gorm:"column:category_id;" json:"category_id"`
	Content    string `gorm:"column:content;" json:"content"`
	Remark     string `gorm:"column:remark;" json:"remark"`
	Thumb      string `gorm:"column:thumb;" json:"thumb"`
	Ctime      uint32 `gorm:"column:ctime;" json:"ctime"`
	Utime      uint32 `gorm:"column:utime;" json:"utime"`
	UserID     uint32 `gorm:"column:user_id;" json:"user_id"`
	Click      uint32 `gorm:"column:click;" json:"click"`
	Like       uint32 `gorm:"column:like;" json:"like"`
	Type       uint8  `gorm:"column:type;" json:"type"`
	State      uint8  `gorm:"column:state;" json:"state"`
	SeoTitle   string `gorm:"column:seo_title;" json:"seo_title"`
	SeoKwds    string `gorm:"column:seo_kwds;" json:"seo_kwds"`
	SeoDesc    string `gorm:"column:seo_desc;" json:"seo_desc"`
	Uuid       string `gorm:"column:uuid;" json:"uuid"`
}

func (a Article) TableName() string {
	return Prefix + "article"
}

//创建
func (a Article) Create(db *gorm.DB, tags []*Tag) error {
	tx := db.Begin()

	//创建文章
	if err := tx.Create(&a).Error; err != nil {
		tx.Rollback()
		return err
	}

	//关联文章-标签
	if len(tags) > 0 {
		for _, v := range tags {
			tid := v.Create(tx)
			if tid > 0 {
				var at = ArticleTag{
					ArticleID: a.ID,
					TagID:     tid,
				}

				if err := at.Create(tx); err != nil {
					tx.Rollback()
					return err
				}
			} else {
				tx.Rollback()
				return errors.New("失败")
			}
		}
	}

	tx.Commit()
	return nil
}

//修改
func (a Article) Update(db *gorm.DB) {

}

//删除
func (a Article) Delete(db *gorm.DB) {

}

//获取一条记录
func (a Article) FindOne(db *gorm.DB) {

}

//获取列表
func (a Article) GetArticleList(db *gorm.DB) {

}
