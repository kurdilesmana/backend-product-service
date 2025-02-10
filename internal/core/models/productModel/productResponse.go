package productModel

import "github.com/kurdilesmana/backend-product-service/internal/core/models/helperModel"

type ProductList struct {
	helperModel.Pagination
	HasMorePage bool      `json:"has_more_page"`
	Rows        []Product `json:"rows"`
}

type ProductListResponse struct {
	helperModel.BaseResponseModel
	Data Product `json:"data"`
}

type ProductResponse struct {
	helperModel.BaseResponseModel
	Data Product `json:"data"`
}
