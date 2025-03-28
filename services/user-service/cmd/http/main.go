package main

import (
	_ "github.com/zelalem-12/bill-aggregation-system_onetab/user-service/docs/openapi" // Swagger docs
	clientService "github.com/zelalem-12/bill-aggregation-system_onetab/user-service/internal/adapter/client"
	"github.com/zelalem-12/bill-aggregation-system_onetab/user-service/internal/adapter/http"
	"github.com/zelalem-12/bill-aggregation-system_onetab/user-service/internal/adapter/persistence/postgres/migration"
	"github.com/zelalem-12/bill-aggregation-system_onetab/user-service/internal/adapter/persistence/postgres/repo"
	"github.com/zelalem-12/bill-aggregation-system_onetab/user-service/internal/adapter/persistence/postgres/seeder"
	"github.com/zelalem-12/bill-aggregation-system_onetab/user-service/internal/app"
	"github.com/zelalem-12/bill-aggregation-system_onetab/user-service/internal/infrastructure/client"
	"github.com/zelalem-12/bill-aggregation-system_onetab/user-service/internal/infrastructure/config"
	"github.com/zelalem-12/bill-aggregation-system_onetab/user-service/internal/infrastructure/database"
	"github.com/zelalem-12/bill-aggregation-system_onetab/user-service/internal/infrastructure/server"
	"go.uber.org/fx"
)

// @title  Bill Aggregation System API(Bill Service)
// @version 1.0
// @description This is the API for User and authenticaion
// @termsOfService http://example.com/terms/

// @contact.name API Support
// @contact.url http://www.example.com/support
// @contact.email support@example.com

// @license.name MIT
// @license.url http://opensource.org/licenses/MIT

// @host 127.0.0.1:8000
// @BasePath /api/v1/

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

// @BasePath /api/v1/
func main() {
	fx.New(
		fx.Provide(

			config.Load,
			server.NewEcho,
			database.InitPostgresDB,
			client.InitSMTPClient,
			clientService.NewEmailService,
			repo.NewUserRepo,
			repo.NewToken,
			repo.NewLinkedAccountRepo,
			clientService.NewBillServiceClient,
		),
		http.Module,
		fx.Invoke(
			migration.MigrateDatabaseSchema,
			seeder.SeedDemoData,
			app.RegisterCQRSHandlers,
			server.ManageServerLifecycle,
		),
	).Run()
}
