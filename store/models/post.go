package models

import "gorm.io/gorm"

type Post struct {
	UserId    uint           `gorm:"type:int(11);unsigned;default:0" json:"user_id"`
	Content   string         `gorm:"type:text;" json:"content"`
	Title     string         `gorm:"type:varchar(150);default:''" json:"title"`
	Subtitle  string         `gorm:"type:varchar(255);default:''" json:"subtitle"`
	Slug      string         `gorm:"type:varchar(150);default:''"`
	View      string         `gorm:"type:int(11);unsigned;default:0" json:"view"`
	DeletedAt gorm.DeletedAt `gorm:"index"  json:"deleted_at" `
	Model
}
