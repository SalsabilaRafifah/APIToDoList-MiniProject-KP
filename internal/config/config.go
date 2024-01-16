package config

import (
	"fmt"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// membuat kondeksi ke database PostgreSQL menggunakan GORM
func ConnectDB() (*gorm.DB, error) {
	//membaca nilai variabel lingkungan dari .env
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_PORT"))

	//membuka koneksi ke database PostgreSQL dengan beberapa konfigurasi dari dsn
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		return nil, err
	}

	// Mengaktifkan fitur autoUpdateTime dan autoCreateTime
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	// Set konfigurasi koneksi database
	sqlDB.SetMaxIdleConns(10)           // Jumlah maksimum koneksi yang diizinkan dalam pool yang tidak aktif
	sqlDB.SetMaxOpenConns(100)          // Jumlah maksimum koneksi yang diizinkan dalam pool (termasuk yang sedang digunakan)
	sqlDB.SetConnMaxLifetime(time.Hour) // Maksimum waktu on untuk koneksi dalam pool

	//Mengembalikan objek database GORM yang sudah terhubung dan dikonfigurasi dengan benar.
	return db, nil
}
