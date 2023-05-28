package tests

import (
	"fmt"
	"os"
	"testing"

	"github.com/rikijoniiskandar/inventory-system/infrastructure/models"
	"github.com/rikijoniiskandar/inventory-system/migrations"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestRunAutoMigrations(t *testing.T) {
	// Konfigurasi koneksi database untuk testing
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
	os.Getenv("DB_USERNAME"),
	os.Getenv("DB_PASSWORD"),
	os.Getenv("DB_HOST"),
	os.Getenv("DB_PORT"),
	os.Getenv("DB_NAME"))
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	assert.NoError(t, err)

	// Menjalankan migrasi otomatis
	err = migrations.RunAutoMigrations(db)
	assert.NoError(t, err)

	// Cek apakah tabel-tabel migrasi telah dibuat
	assert.True(t, db.Migrator().HasTable(&models.Superadmin{}))
	assert.True(t, db.Migrator().HasTable(&models.Role{}))
	assert.True(t, db.Migrator().HasTable(&models.Permission{}))
}