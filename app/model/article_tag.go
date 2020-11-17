package model

type ArticleTag struct {
	ArticleID uint32 `json:"article_id"`
	TagID uint32 `json:"tag_id"`
}

func (a ArticleTag) TableName() string {
	return Prefix + "article_tag"
}