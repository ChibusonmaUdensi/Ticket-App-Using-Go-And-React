package db
import (
	
	"github.com/ChibusonmaUdensi/Ticket-App-Using-Go-And-React/models"
	"gorm.io/gorm"
)

func DBMigrator(db *gorm.DB) error {
	return db.AutoMigrate(&models.Event{})
}