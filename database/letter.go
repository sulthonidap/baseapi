package database

import (
	"gorm.io/gorm"
)

type Letter struct {
	gorm.Model
	PlateNo     string `json:"plateNo"`
	VehicleType string `json:"VehicleType"`
	Name        string `json:"name"`
	Address     string `json:"address"`
	TimePeriod  string `json:"timePeriod"`
}
