package di

import (
	"github.com/google/wire"
	"super-indo-api/controllers"
	"super-indo-api/database"
	"super-indo-api/repositories"
	"super-indo-api/services"
)

var SuperIndoSet = wire.NewSet(
	database.NewDB,
	repositories.NewProductRepository,
	services.NewProductService,
	controllers.NewProductController,
)