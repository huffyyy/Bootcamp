package database

import (
	"log"
	"time"

	"github.com/codeid/hr-api/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func SetupDB() (*gorm.DB, error)  {
	// 1. set datasourcename db config
	dsn := "host=localhost user=postgres password=admin123 dbname=hr_db_test port=5432 sslmode=disable TimeZone=Asia/Jakarta"

	// 2. open connection
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	// 2.1 tampilkan error jika connection fail to db
	if err != nil {
		log.Fatal("failed connect to database : ", err)
	}

	// 3. test create schema oe di db
	// db.Exec("CREATE SCHEMA IF NOT EXISTS OE")

	// 3.1 create variable sqlDB agar bisa akses semua function di gorm.db
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	// 3. set connection pooling
	// default db saat development openConn=100
	sqlDB.SetConnMaxIdleTime(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	log.Println("Connected to database successfully")

	// 4. return value *gorm.DB
	return db, nil
}

func InitAutoMigrate(db *gorm.DB)  {
	err := db.AutoMigrate(
		&models.Region{}, // add model region
	)
	if err != nil {
		log.Fatal("failed to migrate database : ", err)
	}
}