package initializers

import "github.com/kalyanKumarPokkula/Go-jwt/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
}