package models

type User struct {
	Id       string `json:"id"` // 图书ISBN ID
	Username string `gorm:"type:varchar(30);unique" json:"username" form:"username"`
	Phone    string `gorm:"type:varchar(11);default:null;unique" json:"phone"`
	Email    string `gorm:"type:varchar(64);default:null;unique" json:"email"`
	Name     string `gorm:"type:varchar(30);" json:"name" form:"name"`
	Password string `gorm:"type:varchar(65);" json:"-" form:"password"`
	Sex      int    `gorm:"type:tinyint(1);default:0" json:"sex"`
	Status   int    `gorm:"type:tinyint(1);default:0" json:"status"`
	Model
}
