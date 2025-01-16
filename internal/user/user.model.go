package user

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username *string `gorm:"type:varchar(100);uniqueIndex;default:null"`
	Name     *string `gorm:"type:varchar(100);default:null"`
	Email    string  `gorm:"type:varchar(100);uniqueIndex;not null"`
	Password string  `gorm:"type:varchar(255);not null"`
	Bio      *string `gorm:"type:text;default:null"`
}

func (u *User) TableName() string {
	return "user"
}
