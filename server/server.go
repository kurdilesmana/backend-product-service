package server

import (
	"fmt"
	"net/http"

	config "github.com/kurdilesmana/backend-product-service/configs"
	_ "github.com/kurdilesmana/backend-product-service/docs"
	"github.com/kurdilesmana/backend-product-service/pkg/logging"
	custom_middleware "github.com/kurdilesmana/backend-product-service/pkg/middleware"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Http(handler Handler, logger logging.Logger, config config.AppConfig) *echo.Echo {
	e := echo.New()

	// Init middleware
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
	}))

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(custom_middleware.RequestID)

	// Request and Response Middleware
	reqResTrap := custom_middleware.NewReqResMiddleware(logger)
	e.Use(reqResTrap.Middle)

	// Timeout middleware
	middTimeout := custom_middleware.MiddlewareTimeout{
		Timeout: config.MaxRequestTime,
	}
	e.Use(middTimeout.HandlerFunc)

	e.HTTPErrorHandler = func(err error, ctx echo.Context) {
		if ctx.Response().Committed {
			return
		}
		code := http.StatusInternalServerError
		he, ok := err.(*echo.HTTPError)
		if ok {
			code = he.Code
		}
		errMSg := map[string]interface{}{
			"data":    nil,
			"error":   err,
			"message": err.Error(),
		}
		ctx.JSON(code, errMSg)
	}

	// Initialize router for docs
	routerDocs(e)

	// Initialize router for v1
	routerGroupV1(handler, e)

	// Handle not found routes
	e.HTTPErrorHandler = func(err error, c echo.Context) {
		resp := map[string]interface{}{
			"status":  fmt.Sprintf("route %s or method not allowed", http.StatusText(http.StatusNotFound)),
			"message": fmt.Sprintf("route %s", http.StatusText(http.StatusNotFound)),
		}
		if err := c.JSON(http.StatusNotFound, resp); err != nil {
			c.Logger().Error(err)
		}
	}

	return e
}
