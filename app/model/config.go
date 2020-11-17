package model

import (
	"errors"
	"github.com/jinzhu/gorm"
)

type Config struct {
	ID      uint32 `json:"id"`
	Name    string `json:"name"`
	Content string `json:"content"`
}

func (c Config) TableName() string {
	return Prefix + "config"
}

//修改
func (c Config) Update(DB *gorm.DB, conf []Config) error {
	tx := DB.Begin()

	for _, item := range conf {
		if err := tx.Model(&c).Where("name = ?", item.Name).Updates(item).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	tx.Commit()
	return nil
}

//取出某个配置
func (c Config) GetOneConfig(DB *gorm.DB) (*Config, error) {

	if c.Name == "" {
		return nil, errors.New("配置名不得为空")
	}

	if err := DB.Where("name=?", c.Name).First(&c).Error; err != nil {
		return nil, err
	}

	return &c, nil
}

//获取所有配置
func (c Config) GetAllConfig(DB *gorm.DB) (map[string]interface{}, error) {
	var data = make(map[string]interface{})
	var confs []*Config

	if err := DB.Model(&c).Find(&confs).Error; err != nil {
		return data, err
	}

	for _, item := range confs {
		data[item.Name] = item.Content
	}

	return data, nil
}
