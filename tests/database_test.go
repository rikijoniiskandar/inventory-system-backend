package tests

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"
	"strings"
	"testing"

	_ "github.com/go-sql-driver/mysql"

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
		envFile = "../.env-dev"
	case "production":
		envFile = "../.env"
	case "test":
		envFile = "../.env-test"
	default:
		t.Fatalf("Unknown environment: %s", env)
	}

	// Memuat variabel lingkungan dari file .env
	err := loadEnvVariables(envFile)
	assert.NoError(t, err, "Failed to load environment variables gess")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
	os.Getenv("DB_USERNAME"),
	os.Getenv("DB_PASSWORD"),
	os.Getenv("DB_HOST"),
	os.Getenv("DB_PORT"),
	os.Getenv("DB_NAME"))

	// Membuka koneksi database
	db, err := sql.Open("mysql", dsn)
	assert.NoError(t, err, "Failed to open database connection gess")

	// Menguji koneksi database
	err = db.Ping()
	assert.NoError(t, err, "Failed to ping the database")


	// Menutup koneksi database
	err = db.Close()
	assert.NoError(t, err, "Failed to close database connection")
}


