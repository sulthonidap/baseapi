package database

import "gorm.io/gorm"

// Migrate - migrate databases
func Migrate(db *gorm.DB) {
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Activitylog{})
}
