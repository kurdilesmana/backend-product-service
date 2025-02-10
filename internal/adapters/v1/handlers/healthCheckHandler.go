package handlers

import (
	"net/http"

	"github.com/kurdilesmana/backend-product-service/internal/core/ports/healthCheckPort"
	"github.com/kurdilesmana/backend-product-service/pkg/logging"
	"github.com/kurdilesmana/backend-product-service/pkg/middleware"
	"github.com/kurdilesmana/backend-product-service/pkg/web"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type CheckHandler struct {
	HealthCheckService healthCheckPort.IHealthCheckService
	Logger             *logging.Logger
}

func NewHealthCheckHandler(
	healthCheckService healthCheckPort.IHealthCheckService,
	logger *logging.Logger,
) CheckHandler {
	return CheckHandler{
		HealthCheckService: healthCheckService,
		Logger:             logger,
	}
}

// HealthCheck godoc
//
//	@Summary		Get Health Check
//	@Description	LOV untuk health check
//	@Tags			HealthCheck
//	@Accept			json
//	@Produce		json
//	@Router			/health-check		[get]
func (h *CheckHandler) HealthCheck(c echo.Context) error {
	requestID := middleware.GetID(c)
	userCtx := middleware.SetIDx(c.Request().Context(), requestID)

	resp, err := h.HealthCheckService.HealthCheck(userCtx)
	if err != nil {
		h.Logger.Error(logrus.Fields{"requestID": requestID}, err, "error health check")
		return web.ResponseFormatter(c, http.StatusBadRequest, err.Error(), nil, err)
	}
	return web.ResponseFormatter(c, http.StatusOK, "Success", resp, nil)
}
