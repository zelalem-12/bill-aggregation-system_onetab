package main

import (
	_ "github.com/zelalem-12/bill-aggregation-system_onetab/provider-service/docs/openapi" // Swagger docs
	"github.com/zelalem-12/bill-aggregation-system_onetab/provider-service/internal/adapter/http"
	"github.com/zelalem-12/bill-aggregation-system_onetab/provider-service/internal/adapter/persistence/postgres/migration"
	"github.com/zelalem-12/bill-aggregation-system_onetab/provider-service/internal/adapter/persistence/postgres/repo"
	"github.com/zelalem-12/bill-aggregation-system_onetab/provider-service/internal/app"
	"github.com/zelalem-12/bill-aggregation-system_onetab/provider-service/internal/app/service"
	"github.com/zelalem-12/bill-aggregation-system_onetab/provider-service/internal/infrastructure/config"
	"github.com/zelalem-12/bill-aggregation-system_onetab/provider-service/internal/infrastructure/database"
	"github.com/zelalem-12/bill-aggregation-system_onetab/provider-service/internal/infrastructure/server"

	"go.uber.org/fx"
)

// @title  Bill Aggregation System API
// @version 1.0
// @description This is the API for Bill Aggregation System.
// @termsOfService http://example.com/terms/

// @contact.name API Support
// @contact.url http://www.example.com/support
// @contact.email support@example.com

// @license.name MIT
// @license.url http://opensource.org/licenses/MIT

// @host 127.0.0.1:8080
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
			repo.NewProviderRepo,
		),
		http.Module,
		fx.Invoke(
			migration.MigrateDatabaseSchema,
			app.RegisterCQRSHandlers,
			service.SeedutilityProviders,
			server.ManageServerLifecycle,
		),
	).Run()
}
