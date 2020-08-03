package model

import "github.com/jinzhu/gorm"

// Category 分类
type Category struct {
	gorm.Model
	Name string `gorm:"type:varchar(15);not null;"`
}
