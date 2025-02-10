package productRepo

import (
	"context"
	"fmt"
	"time"

	"github.com/kurdilesmana/backend-product-service/internal/core/models/productModel"
	"github.com/kurdilesmana/backend-product-service/internal/core/ports/productPort"
	"github.com/kurdilesmana/backend-product-service/pkg/logging"
	"github.com/kurdilesmana/backend-product-service/pkg/paginate"
	"gorm.io/gorm"
)

type ProductRepository struct {
	DB             *gorm.DB
	KeyTransaction string
	timeout        time.Duration
	log            *logging.Logger
}

func NewProductRepo(db *gorm.DB, keyTransaction string, timeout int, logger *logging.Logger) productPort.IProductRepository {
	return &ProductRepository{
		DB:             db,
		KeyTransaction: keyTransaction,
		timeout:        time.Duration(timeout) * time.Second,
		log:            logger,
	}
}

func (r *ProductRepository) StoreData(ctx context.Context, input productModel.Product) (err error) {
	trx, ok := ctx.Value(r.KeyTransaction).(*gorm.DB)
	if !ok {
		trx = r.DB
	}

	ctxWT, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()

	err = trx.WithContext(ctxWT).Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&input).Error; err != nil {
			return err
		}

		return nil
	})

	return
}

func (r *ProductRepository) ListData(ctx context.Context, paging paginate.Datapaging, opt productModel.Filter) (total int64, products []productModel.Product, err error) {
	trx, ok := ctx.Value(r.KeyTransaction).(*gorm.DB)
	if !ok {
		trx = r.DB
	}

	ctxWT, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()

	query := trx.WithContext(ctxWT).
		Model(&products).
		Where("deleted_at is NULL")

	// add other option
	if len(opt.Keyword) > 0 {
		query = query.Where("(product_code like ? or product_name like ?)", "%"+opt.Keyword+"%", "%"+opt.Keyword+"%")
	}

	query = query.Count(&total)
	if err = query.Error; err != nil {
		return
	}

	paging.BuildQueryGORM(query).Find(&products)

	return
}

func (r *ProductRepository) GetDataByID(ctx context.Context, productID int64) (*productModel.Product, error) {
	ctxWT, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()

	var result productModel.Product
	query := r.DB.WithContext(ctxWT).
		Where("id = ? and deleted_at is NULL", productID).
		First(&result)
	if query.Error != nil {
		return nil, query.Error

	} else if result.ID == 0 {
		return nil, fmt.Errorf("product data not found")
	}

	return &result, nil
}

func (r *ProductRepository) UpdateData(ctx context.Context, productID int64, input productModel.Product) (err error) {
	trx, ok := ctx.Value(r.KeyTransaction).(*gorm.DB)
	if !ok {
		trx = r.DB
	}

	ctxWT, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()

	err = trx.WithContext(ctxWT).Transaction(func(tx *gorm.DB) error {
		var products productModel.Product
		if err := tx.Model(products).
			Where("id = ?", productID).
			Updates(input).Error; err != nil {
			return err
		}

		return nil
	})
	return
}

func (r *ProductRepository) SoftDeleteData(ctx context.Context, productID int64) (err error) {
	trx, ok := ctx.Value(r.KeyTransaction).(*gorm.DB)
	if !ok {
		trx = r.DB
	}

	ctxWT, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()

	err = trx.WithContext(ctxWT).Transaction(func(tx *gorm.DB) error {
		var products productModel.Product
		if err := tx.Model(products).
			Where("id = ?", productID).
			Updates(map[string]interface{}{"deleted_at": time.Now()}).Error; err != nil {
			return err
		}

		return nil
	})

	return nil
}

func (r *ProductRepository) CheckIsExist(ctx context.Context, productCode string) bool {
	ctxWT, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()

	var isAvailable bool
	query := r.DB.WithContext(ctxWT).Model(productModel.Product{}).
		Select("count(*) > 0").
		Where("product_code = ? and deleted_at is NULL", productCode).
		Find(&isAvailable)
	if query.Error != nil {
		return false
	}

	return isAvailable
}
