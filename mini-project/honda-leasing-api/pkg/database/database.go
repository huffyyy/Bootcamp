package database

import (
	"fmt"
	"log"
	"time"

	"github.com/codeid/honda-leasing-api/internal/configs"
	"github.com/codeid/honda-leasing-api/internal/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Database struct {
	*gorm.DB
}

func InitDB(cfg *configs.Config) (*Database, error) {
	dsn := generateDSN(cfg.Database)

	log.Printf("Connecting to database: %s@%s:%s/%s",
		cfg.Database.User,
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.Name,
	)

	gormConfig := &gorm.Config{}

	if cfg.Environment == "development" {
		gormConfig.Logger = logger.Default.LogMode(logger.Info)
	} else {
		gormConfig.Logger = logger.Default.LogMode(logger.Silent)
	}

	db, err := gorm.Open(postgres.Open(dsn), gormConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to connect database: %w", err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get sql instance: %w", err)
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(20)
	sqlDB.SetConnMaxLifetime(time.Hour)

	if err := sqlDB.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}
	log.Println("Database connected successfully")
	return &Database{DB: db}, nil

}

func generateDSN(dbConfig configs.DatabaseConfig) string {
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		dbConfig.Host,
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Name,
		dbConfig.Port,
		dbConfig.SSLMode,
		dbConfig.TimeZone,
	)
}

func CloseDB(db *Database) error {
	if db == nil || db.DB == nil {
		return nil
	}

	sqlDB, err := db.DB.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}

func InitAutoMigrate(database *Database) error {
	schemas := []string{
		"account",
		"dealer",
		"finance",
		"leasing",
		"mst",
	}
	for _, schema := range schemas {
		if err := database.Exec(
			fmt.Sprintf("CREATE SCHEMA IF NOT EXISTS %s", schema),
		).Error; err != nil {
			return fmt.Errorf("failed creating schema %s: %w", schema, err)
		}
	}
	if err := database.AutoMigrate(
		&model.MotorType{},
		&model.Motor{},
	); err != nil {
		return fmt.Errorf("auto migrate failed: %w", err)
	}
	log.Println("✅ Auto migration completed")
	return nil
}
