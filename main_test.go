package main

import (
	"fmt"
	"os"
	"testing"

	tests "github.com/rikijoniiskandar/inventory-system/tests"
)

func TestMain(m *testing.M) {
	// Menjalankan setup sebelum unit test dimulai
	setup()

	// Menjalankan unit test
	exitCode := m.Run()

	// Menjalankan teardown setelah unit test selesai
	teardown()

	// Mengembalikan exit code
	os.Exit(exitCode)
}

func setup() {
	// Lakukan setup yang diperlukan sebelum unit test dimulai
	fmt.Println("Running setup...")
	// Contoh setup: Membuat koneksi ke database test
	err := tests.SetupDatabaseConnection()
	if err != nil {
		fmt.Println("Failed to set up database connection:", err)
		os.Exit(1)
	}
}


func teardown() {
	// Lakukan teardown yang diperlukan setelah unit test selesai
	fmt.Println("Running teardown...")
	// Contoh teardown: Menutup koneksi database test
	err := tests.CloseDatabaseConnection()
	if err != nil {
		fmt.Println("Failed to close database connection:", err)
		os.Exit(1)
	}
}

func TestAll(t *testing.T) {
	// Menjalankan semua unit test
	tests.TestDatabaseConnection(t)
	// Tambahkan unit test lainnya di sini
}