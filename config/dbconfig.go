package config

import (
	"fmt"
	"goAuthTodo/database"
	"goAuthTodo/models"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDB() {
	if errenv := godotenv.Load(); errenv != nil {
		log.Fatal("Gagal memuat file .env")
	}
	var err error
	dsn := os.Getenv("DSN")
	database.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Gagal koneksi database!")
	} else {
		fmt.Println("Koneksi database berhasil!")
	}
}

func RunMigration() {
	if err := database.DB.AutoMigrate(&models.User{}, &models.Todo{}); err != nil {
		log.Fatal("Migrasi table gagal!")
	} else {
		fmt.Println("Migrasi table berhasil!")
	}
}
