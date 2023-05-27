package database

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

/*

 */
func LoadEvnVariables(){
	
	appEnv := os.Getenv("APP_ENV")

	// Default to production if APP_ENV is not set
	if appEnv == "" {
        appEnv = "production"
    }

	var envFilename string

	switch appEnv {
	case "development":
		envFilename = ".env-dev"
	case "test":
		envFilename = ".env-test"
	default:
		envFilename = ".env"
	}

	 // Load environment variables from the appropriate file
	 err := godotenv.Load(envFilename)
	 if err != nil {
		 log.Fatal("Error loading .env file:", err)
	 }
}

/*
 SetupDatabase digunakan sebagai entry point aplikasi
 dan berfungsi untuk inisialisasi router,
 memuat variabel lingkungan, serta memanggil fungsi SetupDatabase
 untuk mengatur koneksi database.
*/
func SetupDatabase() (*gorm.DB, error) {

	LoadEvnVariables()

	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	// Build database connection string
	dsn := dbUsername + ":" + dbPassword + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"

	// Open database connection using MySQL driver
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil

}