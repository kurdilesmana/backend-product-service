package server

import (
	"github.com/labstack/echo/v4"
)

func routerGroupV1(handler Handler, e *echo.Echo) {
	e.GET("/favicon.ico", func(c echo.Context) error { return nil })

	api := e.Group("/api") // /api

	v1 := api.Group("/v1", func(next echo.HandlerFunc) echo.HandlerFunc { // middleware for /api/v1
		return func(c echo.Context) error {
			c.Set("Version", "v1")
			return next(c)
		}
	})

	v1.GET("/health-check", handler.healtHandler.HealthCheck)

	discountGroup := v1.Group("/discount")
	discountGroup.GET("/list", handler.discountHandler.GetListDiscountHandler)
	discountGroup.POST("", handler.discountHandler.CreateDiscountHandler)
	discountGroup.GET("/:id/detail", handler.discountHandler.GetDetailDiscountDataHandler)
	discountGroup.PUT("/:id/update", handler.discountHandler.UpdateDiscountHandler)
	discountGroup.DELETE("/:id/delete", handler.discountHandler.DeleteDiscountHandler)

	stockTypeGroup := v1.Group("/stockType")
	stockTypeGroup.GET("/list", handler.stockTypeHandler.GetListStockTypeHandler)
	stockTypeGroup.POST("", handler.stockTypeHandler.CreateStockTypeHandler)
	stockTypeGroup.GET("/:id/detail", handler.stockTypeHandler.GetDetailStockTypeDataHandler)
	stockTypeGroup.PUT("/:id/update", handler.stockTypeHandler.UpdateStockTypeHandler)
	stockTypeGroup.DELETE("/:id/delete", handler.stockTypeHandler.DeleteStockTypeHandler)

	ticketGroup := v1.Group("/ticket")
	ticketGroup.GET("/list", handler.ticketHandler.GetListTicketHandler)
	ticketGroup.POST("", handler.ticketHandler.CreateTicketHandler)
	ticketGroup.GET("/:id/detail", handler.ticketHandler.GetDetailTicketDataHandler)
	ticketGroup.PUT("/:id/update", handler.ticketHandler.UpdateTicketHandler)
	ticketGroup.DELETE("/:id/delete", handler.ticketHandler.DeleteTicketHandler)

	customerTicketGroup := v1.Group("/customerTicket")
	customerTicketGroup.GET("/list", handler.customerTicketHandler.GetListCustomerTicketHandler)
	customerTicketGroup.POST("", handler.customerTicketHandler.CreateCustomerTicketHandler)
	customerTicketGroup.GET("/:id/detail", handler.customerTicketHandler.GetDetailCustomerTicketDataHandler)
	customerTicketGroup.PUT("/:id/update", handler.customerTicketHandler.UpdateCustomerTicketHandler)
	customerTicketGroup.DELETE("/:id/delete", handler.customerTicketHandler.DeleteCustomerTicketHandler)

	ticketCategoryGroup := v1.Group("/ticketCategory")
	ticketCategoryGroup.GET("/list", handler.ticketCategoryHandler.GetListTicketCategoryHandler)
	ticketCategoryGroup.POST("", handler.ticketCategoryHandler.CreateTicketCategoryHandler)
	ticketCategoryGroup.GET("/:id/detail", handler.ticketCategoryHandler.GetDetailTicketCategoryDataHandler)
	ticketCategoryGroup.PUT("/:id/update", handler.ticketCategoryHandler.UpdateTicketCategoryHandler)
	ticketCategoryGroup.DELETE("/:id/delete", handler.ticketCategoryHandler.DeleteTicketCategoryHandler)
}
