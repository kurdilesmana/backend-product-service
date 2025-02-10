package healthCheckRepo

import (
	"context"
	"time"

	"github.com/kurdilesmana/backend-product-service/internal/core/ports/healthCheckPort"
	"gorm.io/gorm"
)

type HealthCheckRepository struct {
	DB             *gorm.DB
	KeyTransaction string
	timeout        time.Duration
}

func NewHealthCheckRepo(db *gorm.DB, keyTransaction string, timeout int) healthCheckPort.IHealthCheckRepository {
	return &HealthCheckRepository{
		DB:             db,
		KeyTransaction: keyTransaction,
		timeout:        time.Duration(timeout) * time.Second,
	}
}

func (r *HealthCheckRepository) DatabaseCheck(ctx context.Context) (status bool, err error) {
	trx, ok := ctx.Value(r.KeyTransaction).(*gorm.DB)
	if !ok {
		trx = r.DB
	}

	ctxWT, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()

	var result int
	if query := trx.WithContext(ctxWT).Raw("SELECT 1").Row().Scan(&result).Error(); query != "" {
		return false, err
	}

	return true, nil
}
