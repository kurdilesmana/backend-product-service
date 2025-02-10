package productModel

import "github.com/kurdilesmana/backend-product-service/internal/core/models/helperModel"

type ProductRequest struct {
	ProductCode string `json:"product_code" validate:"required"`
	ProductName string `json:"product_name" validate:"required"`
}

type Filter struct {
	Keyword string `json:"keyword" query:"keyword"`
}

type ProductListRequest struct {
	helperModel.PaginationRequest
	Filter
}
