package productPort

import (
	"context"

	"github.com/kurdilesmana/backend-product-service/internal/core/models/productModel"
	"github.com/kurdilesmana/backend-product-service/pkg/paginate"
)

type IProductRepository interface {
	StoreData(ctx context.Context, input productModel.Product) (err error)
	ListData(ctx context.Context, paging paginate.Datapaging, opt productModel.Filter) (total int64, products []productModel.Product, err error)
	GetDataByID(ctx context.Context, productID int64) (*productModel.Product, error)
	UpdateData(ctx context.Context, productID int64, input productModel.Product) (err error)
	SoftDeleteData(ctx context.Context, productID int64) (err error)
	CheckIsExist(ctx context.Context, productCode string) bool
}
