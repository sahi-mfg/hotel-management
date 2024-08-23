package database

import (
	"hotel-management/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func InitDatabase() {
	dsn := "host=localhost user=sahi_hotel password=sahi dbname=hotel_db port=5432 sslmode=disable"
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Clean up inconsistent data before migration
	err = cleanupInconsistentData(DB)
	if err != nil {
		log.Fatal("Failed to clean up data:", err)
	}

	// Disable foreign key checks
	DB.Exec("SET CONSTRAINTS ALL DEFERRED")

	// Migration automatique des tables
	err = DB.AutoMigrate(&models.Client{}, &models.Chambre{}, &models.Reservation{})
	if err != nil {
		log.Fatal("Migration failed:", err)
	}

	// Re-enable foreign key checks
	DB.Exec("SET CONSTRAINTS ALL IMMEDIATE")
}

func cleanupInconsistentData(db *gorm.DB) error {
	// Delete reservations with non-existent client_ids
	result := db.Exec("DELETE FROM reservations WHERE client_id NOT IN (SELECT id FROM clients)")
	if result.Error != nil {
		return result.Error
	}
	return nil
}
