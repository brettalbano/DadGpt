package initializers

import (
	"fmt"
	"os"

	"github.com/brettalbano/DadGpt/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDatabase() {
	var err error
	dsn := os.Getenv("DB_PATH")

	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	DB = database

	if err != nil {
		fmt.Println("Failed to connect to Databse.")
	}
}

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
}
