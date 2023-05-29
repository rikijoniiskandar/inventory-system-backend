package migrations

import (
	"github.com/rikijoniiskandar/inventory-system/infrastructure/models"
	"gorm.io/gorm"
)

func RunAutoMigrations(db *gorm.DB) error {
    // Menjalankan migrasi otomatis untuk semua model
    err := db.AutoMigrate(&models.Superadmin{}, &models.Role{}, &models.Permission{})
    if err != nil {
        return err
    }

    return nil
}