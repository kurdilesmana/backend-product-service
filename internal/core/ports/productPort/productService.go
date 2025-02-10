package productPort

import (
	"context"

	"github.com/kurdilesmana/backend-product-service/internal/core/models/productModel"
	"github.com/kurdilesmana/backend-product-service/pkg/paginate"
)

type IProductService interface {
	CreateProduct(ctx context.Context, input productModel.ProductRequest) error
	DetailProduct(ctx context.Context, ProductID int64) (result *productModel.Product, err error)
	UpdateProduct(ctx context.Context, ProductID int64, input productModel.ProductRequest) error
	DeleteProduct(ctx context.Context, ProductID int64) error
	ListProduct(ctx context.Context, paging paginate.Datapaging, opt productModel.Filter) (result *productModel.ProductList, err error)
}
