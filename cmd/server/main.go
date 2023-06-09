package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"log"

	database "github.com/rikijoniiskandar/inventory-system/infrastructure/databases"
	"github.com/rikijoniiskandar/inventory-system/migrations"
)

func main(){
	// Setup database connection
	db, err := database.SetupDatabase()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	r := gin.Default()

	 // Menjalankan migrasi otomatis
	 err = migrations.RunAutoMigrations(db)
	 if err != nil {
		 panic(err)
	 }

	// Set database connection to context
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	r.LoadHTMLGlob("templates/*")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})

	r.Run()
}