package initializers

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := os.Getenv("DB")
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	fmt.Printf("successfully conntected ")
	if err != nil {
		panic("Failed to connect to DB")
	}
}
