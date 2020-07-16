package model

import "github.com/jinzhu/gorm"

type Article struct {
	gorm.Model
	CategoryId uint   `gorm:"type:tinyint(2);not null;"`
	TagId      uint   `gorm:"type:tinyint(2);default:0;"`
	Title      string `gorm:"type:varchar(100);not null;index:title_idx"`
	Content    string `gorm:"type:text;not null;"`
	Visits     uint   `gorm:"type:mediumint(6);default:1;"`
}
