package migration

import (
	"log"

	"github.com/zelalem-12/bill-aggregation-system_onetab/bill-service/internal/adapter/persistence/postgres/model"
	"gorm.io/gorm"
)

func runPreMigrationScript(db *gorm.DB) {

	db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"")
	db.Exec(`
	DO $$ 
	BEGIN
		IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'status') THEN
			CREATE TYPE status AS ENUM ('paid', 'unpaid', 'overdue');
		END IF;
	END $$;
	`)
}

func MigrateDatabaseSchema(db *gorm.DB) {

	runPreMigrationScript(db)

	userEntities := []interface{}{
		&model.Bill{},
	}

	db.AutoMigrate(userEntities...)

	log.Println("DB Schema Migrated...")
}
