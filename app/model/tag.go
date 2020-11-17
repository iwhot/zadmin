package model

type Tag struct {
	ID      uint32 `json:"id"`
	TagName string `json:"tag_name"`
}

func (t Tag) TableName() string {
	return Prefix + "tag"
}
