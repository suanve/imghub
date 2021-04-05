package config

import (
	"os"
	"time"

	"gorm.io/gorm"
)

var (
	RootDir             = "./web"
	SecretKey           []byte
	DB                  *gorm.DB
	HubHost             = "127.0.0.1"
	Port                = 8081
	TokenExpireDuration = time.Hour * 24
	Protocol            = "https"
)

func init() {

	if os.Getenv("hubHost") != "" {
		HubHost = os.Getenv("hubHost")
	}
	// DB, _ = gorm.Open(sqlite.Open("imghub.db"), &gorm.Config{})
	SecretKey = []byte("457431ecdc83778ac2f5a3441d776b4a")

}
