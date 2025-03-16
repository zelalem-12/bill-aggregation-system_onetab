package migration

import (
	"log"

	"github.com/zelalem-12/bill-aggregation-system_onetab/provider-service/internal/adapter/persistence/postgres/model"
	"gorm.io/gorm"
)

func runPreMigrationScript(db *gorm.DB) {

	db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"")
}

func MigrateDatabaseSchema(db *gorm.DB) {

	runPreMigrationScript(db)

	userEntities := []interface{}{
		&model.Provider{},
	}

	db.AutoMigrate(userEntities...)

	log.Println("DB Schema Migrated...")
}
