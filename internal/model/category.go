package model

// Category 分类
type Category struct {
	ID   uint   `gorm:"primary_key"`
	Name string `gorm:"type:varchar(15);not null;"`
}
