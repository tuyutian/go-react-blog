package models

type User struct {
	Username string `gorm:"type:varchar(30);unique;not null" json:"username" form:"username"`
	Email    string `gorm:"type:varchar(64);default:null;" json:"email"`
	Name     string `gorm:"type:varchar(30);default:null;" json:"name" form:"name"`
	Password string `gorm:"type:varchar(65);" json:"-" form:"password"`
	Status   int    `gorm:"type:tinyint(1);default:0" json:"status"`
	Model
}
