package healthCheckPort

import (
	"context"
)

type IHealthCheckService interface {
	HealthCheck(ctx context.Context) (bool, error)
}
