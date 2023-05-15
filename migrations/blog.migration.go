package main

import (
	"github.com/ceejay1000/blog-api/database"
	"github.com/ceejay1000/blog-api/models"
)

// Migrate the schema
func main() {
	database.InitializeDB().AutoMigrate(&models.Blog{})
}
