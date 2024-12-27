package database

import (
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Activity model
type Activitylog struct {
	gorm.Model
	Who    string    `gorm:"type:varchar(255); NOT NULL" json:"who"`
	When   time.Time `gorm:"type:timestamp; default:CURRENT_TIMESTAMP; NOT NULL" json:"when"`
	What   string    `gorm:"type:varchar(255); NOT NULL" json:"what"`
	IPAddr string    `gorm:"type:varchar(255)" json:"ipaddr"` // client IP Address
	Agent  string    `json:"agent"`                           // Client user agent
}

// AddActivity write activity log
func AddActivitylog(who, what string, c *gin.Context) (bool, error) {
	addr := c.ClientIP()
	agent := c.Request.Header.Get("User-Agent")
	if err := DBConn.Create(&Activitylog{
		Who:    who,
		What:   what,
		IPAddr: addr,
		Agent:  agent,
	}).Error; err != nil {
		return false, err
	}

	return true, nil
}
