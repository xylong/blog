package model

type Tag struct {
	ID   uint   `gorm:"primary_key"`
	Name string `gorm:"type:varchar(15);not null;"`
}
