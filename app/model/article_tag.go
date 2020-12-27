package model

import "github.com/jinzhu/gorm"

type ArticleTag struct {
	ArticleID uint32 `json:"article_id"`
	TagID uint32 `json:"tag_id"`
}

func (a ArticleTag) TableName() string {
	return Prefix + "article_tag"
}

//创建
func (a ArticleTag) Create(db *gorm.DB) error {
	return db.Model(&a).Create(&a).Error
}

//删除
func (a ArticleTag) DeleteArticleTagByArtID(db *gorm.DB) error {
	return db.Where("article_id=?",a.ArticleID).Delete(ArticleTag{}).Error
}