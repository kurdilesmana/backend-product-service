package healthCheckPort

import (
	"context"
)

type IHealthCheckRepository interface {
	DatabaseCheck(ctx context.Context) (status bool, err error)
}
