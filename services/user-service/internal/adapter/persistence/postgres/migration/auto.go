package migration

import (
	"log"

	"github.com/zelalem-12/onetab/internal/adapter/persistence/postgres/model"
	"gorm.io/gorm"
)

func runPreMigrationScript(db *gorm.DB) {

	db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"")
}

func MigrateDatabaseSchema(db *gorm.DB) {

	runPreMigrationScript(db)

	userEntities := []interface{}{
		&model.User{},
		&model.Token{},
		&model.LinkedAccount{},
	}

	db.AutoMigrate(userEntities...)

	log.Println("DB Schema Migrated...")
}
