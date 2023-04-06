package initializers

import (
	"fmt"
	"os"
	"time"

	"DadGpt/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDatabase() {
	var err error
	dsn := os.Getenv("DB_PATH")

	num_retries := 0

	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		for num_retries < 4 {
			fmt.Printf("Failed to connect to Databse. Retrying in 5 seconds. (Attempt %d/4)", num_retries+1)
			time.Sleep(5 * time.Second)
			database, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
			if err == nil {
				break
			}
			num_retries++
		}

	}

	DB = database

}

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
}
