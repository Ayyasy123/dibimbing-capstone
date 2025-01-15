package config

import (
	"fmt"
	"log"
	"os"

	"github.com/Ayyasy123/dibimbing-capstone.git/entity"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	// Ambil variabel lingkungan dari .env
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	// String koneksi ke MYSQL with docker
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPassword, dbHost, dbPort, dbName)

	// Konfigurasi koneksi database with xampp
	// dsn := "root:@tcp(127.0.0.1:3306)/capstone?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Auto-migrasi semua entitas
	err = db.AutoMigrate(
		&entity.User{},
		&entity.Service{},
		&entity.Booking{},
		&entity.Payment{},
		&entity.Review{},
	)
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	DB = db
	log.Println("Database connected and migrated successfully")
}
