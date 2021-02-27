package model

import (
	"errors"
	"github.com/jinzhu/gorm"
	page2 "github.com/iwhot/zadmin/system/page"
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

	//更改文件状态
	if a.Thumb != "" {
		var file = Files{
			Url:   a.Thumb,
			State: 1,
		}

		if err := file.UpdateFiles(tx); err != nil {
			tx.Rollback()
			return err
		}
	}

	tx.Commit()
	return nil
}

//修改
func (a Article) Update(db *gorm.DB, tags []*Tag) error {
	art, err := a.FindOne(db)
	if err != nil || art == nil {
		return errors.New("文章不存在")
	}

	tx := db.Begin()
	if err := tx.Model(&a).Update(&a).Error; err != nil {
		tx.Rollback()
		return err
	}

	//关联文章-标签
	if len(tags) > 0 {
		if err := tx.Where("article_id=?", a.ID).Delete(ArticleTag{}).Error; err != nil {
			tx.Rollback()
			return err
		}
		//删除标签
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

	//更改文件状态
	if a.Thumb != "" && a.Thumb != art.Thumb {
		//删除原文件
		var f1 = Files{
			Url:   art.Thumb,
			State: 0,
		}

		if id := f1.SearchFileByAddress(tx); id != 0 {
			f1.ID = id
			_ = f1.Delete(tx)
		}

		var file = Files{
			Url:   a.Thumb,
			State: 1,
		}
		if err := file.UpdateFiles(tx); err != nil {
			tx.Rollback()
			return err
		}
	}

	tx.Commit()
	return nil
}

//删除
func (a Article) Delete(db *gorm.DB) error {
	art, err := a.FindOne(db)
	if err != nil || art == nil {
		return errors.New("文章不存在")
	}

	tx := db.Begin()
	//删除文章
	if err := tx.Where("id", a.ID).Delete(a).Error; err != nil {
		tx.Rollback()
		return err
	}
	//删除标签
	if err := tx.Where("article_id=?", a.ID).Delete(ArticleTag{}).Error; err != nil {
		tx.Rollback()
		return err
	}
	//删除文件
	if art.Thumb != "" {
		var f1 = Files{
			Url:   art.Thumb,
			State: 0,
		}

		if id := f1.SearchFileByAddress(tx); id != 0 {
			f1.ID = id
			_ = f1.Delete(tx)
		}
	}

	tx.Commit()
	return nil
}

//获取一条记录
func (a Article) FindOne(db *gorm.DB) (*Article, error) {
	if a.ID == 0 {
		return &a, errors.New("文章不存在")
	}

	var art = []*Article{}
	if err := db.Model(&a).Where("id=?", a.ID).Offset(0).Limit(1).Find(&art).Error; err != nil {
		return nil, err
	}

	return art[0], nil
}

//获取列表
func (a Article) GetArticleList(db *gorm.DB, page, pageSize int) ([]*Article,error) {
	var art = []*Article{}
	var offset = page2.GetOffset(page, pageSize)
	var mod *gorm.DB
	mod = db.Model(&a)
	if err := mod.Offset(offset).Limit(pageSize).Order("ctime desc,id desc").Find(&art).Error; err != nil {
		return nil, err
	}
	return art, nil
}
