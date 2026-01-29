package database

import (
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// SetupDB membuka koneksi database dan mengatur schema default ke `hr`.
func SetupDB() (*gorm.DB, error) {

	// 1. Data Source Name
	dsn := "host=localhost user=postgres password=admin123 dbname=hr_db_test port=5432 sslmode=disable TimeZone=Asia/Jakarta"

	// 2. Open connection ke PostgreSQL
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), // tampilkan SQL log (dev mode)
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		log.Fatal("failed connect to database : ", err)
	}

	// 3. Set default schema ke `hr`
	// Semua query otomatis akan mengarah ke hr.*
	if err := db.Exec("SET search_path TO hr").Error; err != nil {
		log.Fatal("failed to set search_path to hr : ", err)
	}

	// 4. Ambil native *sql.DB untuk connection pooling
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	// 5. Konfigurasi connection pooling
	sqlDB.SetConnMaxIdleTime(10 * time.Second)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	log.Println("Connected to database successfully (schema: hr)")
	return db, nil
}

func InitAutoMigrate(db *gorm.DB) {
	// db.Exec("CREATE SCHEMA IF NOT EXISTS hr")

	/*
		err := db.AutoMigrate(
			&models.Region{},
		)
		if err != nil {
			log.Fatal("failed to migrate database : ", err)
		}
	*/

	/*
		// 1. Migrate Region dulu
		if err := db.AutoMigrate(&models.Region{}); err != nil {
			log.Fatal("Error migrating region", err)
		}

		// 2. Baru Migrate Country
		if err := db.AutoMigrate(&models.Country{}); err != nil {
			log.Fatal("Error migrating country", err)
		}
	*/
	log.Println("Skip AutoMigrate: using existing HR schema (no schema changes)")
}
