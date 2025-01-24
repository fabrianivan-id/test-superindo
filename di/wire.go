package di

import (
	"github.com/google/wire"
	"github.com/fabrianivan-id/test-superindo/controllers"
	"github.com/fabrianivan-id/test-superindo/database"
	"github.com/fabrianivan-id/test-superindo/repositories"
	"github.com/fabrianivan-id/test-superindo/services"
)

var SuperIndoSet = wire.NewSet(
	database.NewDB,
	repositories.NewProductRepository,
	services.NewProductService,
	controllers.NewProductController,
)