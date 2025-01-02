package postgres

import (
	"fmt"
	"time"

	"github.com/uluk001/spend/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	db *gorm.DB
}

func NewPostgresDB(cfg *config.Config) (*Database, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		cfg.Database.Host, cfg.Database.User, cfg.Database.Password, cfg.Database.Name,
		cfg.Database.Port,
	)

	config := &gorm.Config{}

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,  // data source name, refer https://github.com/jackc/pgx
		PreferSimpleProtocol: true, // disables implicit prepared statement usage. By default pgx automatically uses the extended protocol
	}), config)

	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	// Настройка пула соединений
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	return &Database{db: db}, nil
}

func (d *Database) GetDB() *gorm.DB {
	return d.db
}
