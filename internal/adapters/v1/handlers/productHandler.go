package handlers

import (
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/kurdilesmana/backend-product-service/internal/core/models/productModel"
	"github.com/kurdilesmana/backend-product-service/internal/core/ports/productPort"
	"github.com/kurdilesmana/backend-product-service/pkg/logging"
	"github.com/kurdilesmana/backend-product-service/pkg/middleware"
	"github.com/kurdilesmana/backend-product-service/pkg/paginate"
	"github.com/kurdilesmana/backend-product-service/pkg/web"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type ProductHandler struct {
	ProductService productPort.IProductService
	Logger         *logging.Logger
	Validator      *validator.Validate
}

func NewProductHandler(
	ProductService productPort.IProductService,
	logger *logging.Logger,
	validator *validator.Validate,
) ProductHandler {
	return ProductHandler{
		ProductService: ProductService,
		Logger:         logger,
		Validator:      validator,
	}
}

// CreateProductHandler handles create Product
//
//	@Summary		API create Product
//	@Description	endpoint create Product
//	@Tags			Product
//	@Accept			json
//	@Produce		json
//	@Param			Authorization	header		string								false	"Authorization"	example(Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJqdGkiOiJiNWQxNjY5NS0xZjJlLTQ1...)
//	@Param			Request			body		productModel.ProductRequest			true	"Request Parameters"
//	@Success		200				{object}	helperModel.BaseResponseModel		"Response Success"
//	@Failure		400				{object}	helperModel.BaseResponseModel		"Response Error"
//	@Router			/product	[post]
func (h *ProductHandler) CreateProductHandler(ctx echo.Context) error {
	requestID := middleware.GetID(ctx)
	userCtx := middleware.SetIDx(ctx.Request().Context(), requestID)

	/* userMakingRequest, err := middleware.GetDataFromJWTContext(ctx)
	if err != nil {
		return web.ResponseFormatter(ctx, http.StatusUnauthorized, "Unauthorized", nil, err)
	} */

	var payload productModel.ProductRequest
	if err := ctx.Bind(&payload); err != nil {
		h.Logger.Error(logrus.Fields{"request_id": requestID, "error": err.Error()}, payload, "create product payload")
		return web.ResponseFormatter(ctx, http.StatusBadRequest, "Bad Request", nil, err)
	}

	// log payload
	h.Logger.Info(logrus.Fields{"request_id": requestID}, payload, "create product payload")

	// Validate slice payload
	err := h.Validator.Struct(payload)
	if err != nil {
		h.Logger.Error(logrus.Fields{"request_id": requestID, "error": err.Error()}, payload, "create product payload validation")
		return web.ResponseErrValidationWithCode(ctx, "bad request", err, http.StatusBadRequest)
	}

	if err := h.ProductService.CreateProduct(userCtx, payload); err != nil {
		h.Logger.Error(logrus.Fields{"request_id": requestID, "error": err.Error()}, payload, "error create product")
		return web.ResponseFormatter(ctx, http.StatusBadRequest, err.Error(), nil, err)
	}
	return web.ResponseFormatter(ctx, http.StatusOK, "Success", "", nil)
}

// UpdateProductHandler handles Update Product
//
//	@Summary		API Update Product
//	@Description	endpoint Update Product
//	@Tags			Product
//	@Accept			json
//	@Produce		json
//	@Param			Authorization				header		string								false	"Authorization"	example(Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJqdGkiOiJiNWQxNjY5NS0xZjJlLTQ1...)
//	@Param			id							path		string								true	"Product ID"	example(1)
//	@Param			Request						body		productModel.ProductRequest	true	"Request Parameters"
//	@Success		200							{object}	helperModel.BaseResponseModel		"Response Success"
//	@Failure		400							{object}	helperModel.BaseResponseModel		"Response Error"
//	@Router			/product/{id}/update	[put]
func (h *ProductHandler) UpdateProductHandler(ctx echo.Context) error {
	requestID := middleware.GetID(ctx)
	userCtx := middleware.SetIDx(ctx.Request().Context(), requestID)

	/* userMakingRequest, err := middleware.GetDataFromJWTContext(ctx)
	if err != nil {
		return web.ResponseFormatter(ctx, http.StatusUnauthorized, "Unauthorized", nil, err)
	} */

	// get path param
	idStr := ctx.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		h.Logger.Error(logrus.Fields{"request_id": requestID, "error": err.Error(), "product_id": id}, nil, "update product id")
		return web.ResponseFormatter(ctx, http.StatusBadRequest, "Bad Request", nil, err)
	}

	var payload productModel.ProductRequest
	if err := ctx.Bind(&payload); err != nil {
		h.Logger.Error(logrus.Fields{"request_id": requestID, "error": err.Error()}, payload, "update product payload")
		return web.ResponseFormatter(ctx, http.StatusBadRequest, "Bad Request", nil, err)
	}

	// log payload
	h.Logger.Info(logrus.Fields{"request_id": requestID}, payload, "update product payload")

	// Validate slice payload
	err = h.Validator.Struct(payload)
	if err != nil {
		h.Logger.Error(logrus.Fields{"request_id": requestID, "error": err.Error()}, payload, "update product payload validation")
		return web.ResponseErrValidationWithCode(ctx, "bad request", err, http.StatusBadRequest)
	}

	if err := h.ProductService.UpdateProduct(userCtx, id, payload); err != nil {
		h.Logger.Error(logrus.Fields{"request_id": requestID, "error": err.Error()}, payload, "error update product")
		return web.ResponseFormatter(ctx, http.StatusBadRequest, err.Error(), nil, err)
	}
	return web.ResponseFormatter(ctx, http.StatusOK, "Success", "", nil)
}

// DeleteProductHandler handles delete Product
//
//	@Summary		API delete Product
//	@Description	endpoint delete Product
//	@Tags			Product
//	@Accept			json
//	@Produce		json
//	@Param			Authorization				header		string							false	"Authorization"	example(Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJqdGkiOiJiNWQxNjY5NS0xZjJlLTQ1...)
//	@Param			id							path		string							true	"Product ID"	example(1)
//	@Success		200							{object}	helperModel.BaseResponseModel	"Response Success"
//	@Failure		400							{object}	helperModel.BaseResponseModel	"Response Error"
//	@Router			/product/{id}/delete	[delete]
func (h *ProductHandler) DeleteProductHandler(ctx echo.Context) error {
	requestID := middleware.GetID(ctx)
	userCtx := middleware.SetIDx(ctx.Request().Context(), requestID)

	/* userMakingRequest, err := middleware.GetDataFromJWTContext(ctx)
	if err != nil {
		return web.ResponseFormatter(ctx, http.StatusUnauthorized, "Unauthorized", nil, err)
	} */

	// get path param
	idStr := ctx.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		h.Logger.Error(logrus.Fields{"request_id": requestID, "error": err.Error(), "product_id": id}, nil, "update product id")
		return web.ResponseFormatter(ctx, http.StatusBadRequest, "Bad Request", nil, err)
	}

	// log payload
	h.Logger.Info(logrus.Fields{"request_id": requestID, "product_id": id}, nil, "delete product payload")
	if err := h.ProductService.DeleteProduct(userCtx, id); err != nil {
		h.Logger.Error(logrus.Fields{"request_id": requestID, "error": err.Error(), "product_id": id}, nil, "error delete product")
		return web.ResponseFormatter(ctx, http.StatusBadRequest, err.Error(), nil, err)
	}
	return web.ResponseFormatter(ctx, http.StatusOK, "Success", "", nil)
}

// DetailProductHandler handles detail Product
//
//	@Summary		API Detail Product
//	@Description	endpoint detail Product
//	@Tags			Product
//	@Accept			json
//	@Produce		json
//	@Param			Authorization				header		string							false	"Authorization"	example(Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJqdGkiOiJiNWQxNjY5NS0xZjJlLTQ1...)
//	@Param			id							path		string							true	"Product ID"	example(1)
//	@Success		200							{object}	productModel.ProductResponse	"Response Success"
//	@Failure		400							{object}	helperModel.BaseResponseModel	"Response Error"
//	@Router			/product/{id}/detail	[get]
func (h *ProductHandler) DetailProductHandler(ctx echo.Context) error {
	requestID := middleware.GetID(ctx)
	userCtx := middleware.SetIDx(ctx.Request().Context(), requestID)

	/* userMakingRequest, err := middleware.GetDataFromJWTContext(ctx)
	if err != nil {
		return web.ResponseFormatter(ctx, http.StatusUnauthorized, "Unauthorized", nil, err)
	} */

	// get path param
	idStr := ctx.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		h.Logger.Error(logrus.Fields{"request_id": requestID, "error": err.Error(), "product_id": id}, nil, "update product id")
		return web.ResponseFormatter(ctx, http.StatusBadRequest, "Bad Request", nil, err)
	}

	productData, err := h.ProductService.DetailProduct(userCtx, id)
	if err != nil {
		h.Logger.Error(logrus.Fields{"request_id": requestID, "error": err.Error(), "product_id": id}, nil, "error get product")
		return web.ResponseFormatter(ctx, http.StatusBadRequest, err.Error(), nil, err)
	}

	return web.ResponseFormatter(ctx, http.StatusOK, "Success", productData, nil)
}

// ListProductHandler handles Product
//
//	@Summary		API List Product
//	@Description	endpoint Product list
//	@Tags			Product
//	@Accept			json
//	@Produce		json
//	@Param			Authorization		header		string								false	"Authorization"	example(Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJqdGkiOiJiNWQxNjY5NS0xZjJlLTQ1...)
//	@Param			request				query		productModel.ProductListRequest		false	"Request Parameters"
//	@Success		200					{object}	productModel.ProductListResponse	"Response Success"
//	@Failure		400					{object}	helperModel.BaseResponseModel		"Response Error"
//	@Router			/product/list	[get]
func (h *ProductHandler) ListProductHandler(ctx echo.Context) error {
	requestID := middleware.GetID(ctx)
	userCtx := middleware.SetIDx(ctx.Request().Context(), requestID)

	/* userMakingRequest, err := middleware.GetDataFromJWTContext(ctx)
	if err != nil {
		return web.ResponseFormatter(ctx, http.StatusUnauthorized, "Unauthorized", nil, err)
	} */

	// get query param
	paging := paginate.PreparePagination(map[string]string{
		"limit": ctx.QueryParam("limit"),
		"page":  ctx.QueryParam("page"),
		"sort":  ctx.QueryParam("sort"),
	}, []string{
		"id",
		"created_at",
		"updated_at",
	})

	// log payload
	h.Logger.Info(logrus.Fields{"request_id": requestID}, paging, "request list product")

	keyword := ctx.QueryParam("keyword")
	payload := productModel.Filter{
		Keyword: keyword,
	}
	// log payload
	h.Logger.Info(logrus.Fields{"request_id": requestID}, payload, "request list product")

	responseData, err := h.ProductService.ListProduct(userCtx, paging, payload)
	if err != nil {
		h.Logger.Error(logrus.Fields{"request_id": requestID, "error": err.Error(), "payload": payload}, paging, "error get list product")
		return web.ResponseFormatter(ctx, http.StatusBadRequest, err.Error(), nil, err)
	}

	return web.ResponseFormatter(ctx, http.StatusOK, "Success", responseData, nil)
}
