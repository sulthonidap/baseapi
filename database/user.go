package database

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string     `gorm:"NOT NULL" json:"username"`
	Fullname string     `gorm:"NOT NULL" json:"fullname"`
	Phone    string     `gorm:"NOT NULL" json:"phone"`
	Active   bool       `gorm:"NOT NULL;default:TRUE" json:"active"`
	Password string     `gorm:"NOT NULL" json:"-" form:"password"`
	Role     string     `gorm:"type:varchar(255); NOT NULL" json:"role"` //admin, customer
	ForceCpw bool       `gorm:"NOT NULL;default:FALSE" json:"forceCpw"`
	LastCpw  *time.Time `gorm:"type:datetime;default:current_timestamp" json:"lastCpw" form:"lastCpw"`
}
