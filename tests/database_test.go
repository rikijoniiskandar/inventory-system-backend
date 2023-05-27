package tests

import (
	"bufio"
	"database/sql"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func loadEnvVariables(envFile string) error {
	file, err := os.Open(envFile)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.SplitN(line, "=", 2)
		if len(parts) == 2 {
			key := strings.TrimSpace(parts[0])
			value := strings.TrimSpace(parts[1])
			os.Setenv(key, value)
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}

func TestDatabaseConnection(t *testing.T) {
	env := os.Getenv("APP_ENV")

	// Menyesuaikan nama file .env berdasarkan environment
	var envFile string
	switch env {
	case "development":
		envFile = ".env-dev"
	case "production":
		envFile = ".env"
	case "test":
		envFile = ".env-test"
	default:
		t.Fatalf("Unknown environment: %s", env)
	}

	// Memuat variabel lingkungan dari file .env
	err := loadEnvVariables(envFile)
	assert.NoError(t, err, "Failed to load environment variables")

	dsn := "username:password@tcp(localhost:3306)/database_name"
	// Membuka koneksi database
	db, err := sql.Open("mysql", dsn)
	assert.NoError(t, err, "Failed to open database connection")

	// Menguji koneksi database
	err = db.Ping()
	assert.NoError(t, err, "Failed to ping the database")


	// Menutup koneksi database
	err = db.Close()
	assert.NoError(t, err, "Failed to close database connection")
}


