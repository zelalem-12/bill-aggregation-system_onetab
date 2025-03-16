package database

import (
	"fmt"

	"github.com/zelalem-12/bill-aggregation-system_onetab/bill-service/internal/infrastructure/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitPostgresDB(cfg *config.Config) (*gorm.DB, error) {

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", cfg.POSTGRES_HOST, cfg.POSTGRES_PORT, cfg.POSTGRES_USER, cfg.POSTGRES_PASSWORD, cfg.POSTGRES_DATABASE)

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
