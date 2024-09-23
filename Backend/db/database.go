package db

import (
	"fmt"

	"github.com/gofiber/fiber/v2/log"
	"github.com/ChibusonmaUdensi/Ticket-App-Using-Go-And-React/config"
		"gorm.io/gorm"
		"gorm.io/gorm/logger"
		"gorm.io/driver/postgres"
)
func Init(config *config.EnvConfig, DBMigrator func(db *gorm.DB) error) *gorm.DB{
	uri := fmt.Sprintf(`
	    host=%s user=%s dbname=%s password=%s sslmode=%s port=5433`,
		config.DBHost, config.DBUser, config.DBName, config.DBPassword, config.DBSSLMode,
	)
	db, err := gorm.Open(postgres.Open(uri), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err!= nil {
        log.Fatal("Error connecting to the database", err)
    }
	log.Info("Database initialized!")

	if err := DBMigrator(db); err!= nil {
        log.Fatal("Error migrating the database", err)
    }
	return db
}