package server

import (
	"github.com/go-playground/validator/v10"
	"github.com/kurdilesmana/backend-product-service/deps"
	handler "github.com/kurdilesmana/backend-product-service/internal/adapters/v1/handlers"
)

type Handler struct {
	healtHandler   handler.CheckHandler
	userHandler    handler.UserHandler
	productHandler handler.ProductHandler
}

func SetupHandler(dep deps.Dependency) Handler {
	//init validator
	validator := validator.New()

	return Handler{
		healtHandler:   handler.NewHealthCheckHandler(dep.HealthCheckService, dep.Logger),
		userHandler:    handler.NewUserHandler(dep.UserService, dep.Logger, validator),
		productHandler: handler.NewProductHandler(dep.ProductService, dep.Logger, validator),
	}
}
