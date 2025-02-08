package server

import (
	"github.com/kurdilesmana/backend-product-service/deps"
	handler "github.com/kurdilesmana/backend-product-service/internal/adapters/v1/handlers"
	"github.com/kurdilesmana/backend-product-service/pkg/kbsvalidator"
)

type Handler struct {
	healtHandler          handler.CheckHandler
	discountHandler       handler.DiscountHandler
	stockTypeHandler      handler.StockTypeHandler
	ticketHandler         handler.TicketHandler
	customerTicketHandler handler.CustomerTicketHandler
	ticketCategoryHandler handler.TicketCategoryHandler
}

func SetupHandler(dep deps.Dependency) Handler {
	//init validator
	validator := kbsvalidator.New()

	return Handler{
		healtHandler:          handler.NewHealthCheckHandler(dep.HealthCheckService, dep.Logger),
		discountHandler:       handler.NewDiscountHandler(dep.DiscountService, dep.Logger, validator),
		stockTypeHandler:      handler.NewStockTypeHandler(dep.StockTypeService, dep.Logger, validator),
		ticketHandler:         handler.NewTicketHandler(dep.TicketService, dep.Logger, validator),
		customerTicketHandler: handler.NewCustomerTicketHandler(dep.CustomerTicketService, dep.Logger, validator),
		ticketCategoryHandler: handler.NewTicketCategoryHandler(dep.TicketCategoryService, dep.Logger, validator),
	}
}
