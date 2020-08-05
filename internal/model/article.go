package model

import (
	"github.com/jinzhu/gorm"
)

type Article struct {
	gorm.Model
	CategoryId uint     `gorm:"type:tinyint(2);not null;index:cate_idx"`
	Category   Category `json:"category"`
	Title      string   `gorm:"type:varchar(100);not null;index:title_idx"`
	Content    string   `gorm:"type:text;not null;"`
	Visits     uint     `gorm:"type:mediumint(6);default:1;"`
}

//func (article *Article) BeforeCreate(scope *gorm.Scope) error {
//	return scope.SetColumn("CreatedOn", time.Now().Unix())
//}
