package healthCheckService

import (
	"context"

	"github.com/kurdilesmana/backend-product-service/internal/core/ports/healthCheckPort"
	"github.com/kurdilesmana/backend-product-service/pkg/logging"
)

type healthCheckService struct {
	HealthCheckRepo healthCheckPort.IHealthCheckRepository
	Logger          logging.Logger
}

func NewHealthCheckService(
	healthCheckRepo healthCheckPort.IHealthCheckRepository,
	logger logging.Logger,
) healthCheckPort.IHealthCheckService {
	return &healthCheckService{
		HealthCheckRepo: healthCheckRepo,
		Logger:          logger,
	}
}

func (s *healthCheckService) HealthCheck(ctx context.Context) (bool, error) {
	healthy, err := s.HealthCheckRepo.DatabaseCheck(ctx)
	if err != nil {
		return false, err
	}

	return healthy, nil
}
