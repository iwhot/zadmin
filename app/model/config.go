package model

import "github.com/jinzhu/gorm"

type Config struct {
	ID      uint32 `json:"id"`
	Name    string `json:"name"`
	Content string `json:"content"`
}

func (c Config) TableName() string {
	return Prefix + "config"
}

//修改
func (c Config) Update(DB *gorm.DB) {

}

//取出某个配置
func (c Config) GetOneConfig(DB *gorm.DB) {

}
