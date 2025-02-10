package productService

import (
	"context"
	"fmt"
	"math"

	"github.com/kurdilesmana/backend-product-service/internal/core/models/helperModel"
	"github.com/kurdilesmana/backend-product-service/internal/core/models/productModel"
	"github.com/kurdilesmana/backend-product-service/internal/core/ports/productPort"
	"github.com/kurdilesmana/backend-product-service/pkg/logging"
	"github.com/kurdilesmana/backend-product-service/pkg/middleware"
	"github.com/kurdilesmana/backend-product-service/pkg/paginate"
	"github.com/sirupsen/logrus"
)

type productService struct {
	productRepo productPort.IProductRepository
	log         logging.Logger
}

func NewproductService(
	ProductRepo productPort.IProductRepository,
	logger logging.Logger,
) productPort.IProductService {
	return &productService{
		productRepo: ProductRepo,
		log:         logger,
	}
}

func (s *productService) CreateProduct(ctx context.Context, request productModel.ProductRequest) (err error) {
	// request id
	requestID := middleware.GetIDx(ctx)
	defer s.log.Info(logrus.Fields{"request_id": requestID}, request, "process service completed..")

	// validate data is exist
	existData := s.productRepo.CheckIsExist(ctx, request.ProductCode)
	if existData {
		err = fmt.Errorf("data already exists...")
		s.log.Warn(logrus.Fields{}, nil, err.Error())
		return

	}

	product := productModel.Product{
		ProductCode: request.ProductCode,
		ProductName: request.ProductName,
	}
	s.log.Info(logrus.Fields{"request_id": requestID}, product, "Create product to db")
	if err = s.productRepo.StoreData(ctx, product); err != nil {
		s.log.Warn(logrus.Fields{"error": err}, nil, err.Error())
		err = fmt.Errorf("failed to create product")
		return
	}

	return nil
}

func (s *productService) DetailProduct(ctx context.Context, ProductID int64) (result *productModel.Product, err error) {
	// request id
	requestID := middleware.GetIDx(ctx)
	defer s.log.Info(logrus.Fields{"request_id": requestID}, &result, "process service completed..")

	result, err = s.productRepo.GetDataByID(ctx, ProductID)
	if err != nil {
		s.log.Warn(logrus.Fields{"error": err}, &result, "failed get data...")
		return nil, fmt.Errorf("product not found")
	}

	return
}

func (s *productService) UpdateProduct(ctx context.Context, ProductID int64, input productModel.ProductRequest) (err error) {
	// request id
	requestID := middleware.GetIDx(ctx)
	defer s.log.Info(logrus.Fields{"request_id": requestID}, nil, "process service completed..")

	// validate item category exists
	productData, err := s.productRepo.GetDataByID(ctx, ProductID)
	if err != nil {
		s.log.Warn(logrus.Fields{"error": err}, nil, "failed get data...")
		return fmt.Errorf("product not found")
	}

	// validate item feature is exist
	if productData.ProductCode != input.ProductCode {
		existData := s.productRepo.CheckIsExist(ctx, input.ProductCode)
		if existData {
			err = fmt.Errorf("data already exists...")
			s.log.Warn(logrus.Fields{}, nil, err.Error())
			return
		}
	}

	product := productModel.Product{
		ProductCode: input.ProductCode,
		ProductName: input.ProductName,
	}
	s.log.Info(logrus.Fields{"request_id": requestID}, product, "Create product to db")
	if err = s.productRepo.UpdateData(ctx, ProductID, product); err != nil {
		s.log.Warn(logrus.Fields{"error": err}, nil, err.Error())
		err = fmt.Errorf("failed to create product")
		return
	}

	return nil
}

func (s *productService) DeleteProduct(ctx context.Context, ProductID int64) (err error) {
	// request id
	requestID := middleware.GetIDx(ctx)
	defer s.log.Info(logrus.Fields{"request_id": requestID}, nil, "process service completed..")

	// validate product exist
	product, err := s.productRepo.GetDataByID(ctx, ProductID)
	if err != nil {
		s.log.Warn(logrus.Fields{"error": err}, nil, "failed get data...")
		return fmt.Errorf("product not found")
	}

	s.log.Info(logrus.Fields{"request_id": requestID}, product, "Create product to db")
	if err = s.productRepo.SoftDeleteData(ctx, ProductID); err != nil {
		s.log.Warn(logrus.Fields{"error": err}, nil, err.Error())
		err = fmt.Errorf("failed to create product")
		return
	}

	return nil
}

func (s *productService) ListProduct(ctx context.Context, paging paginate.Datapaging, opt productModel.Filter) (result *productModel.ProductList, err error) {
	// request id
	requestID := middleware.GetIDx(ctx)
	defer s.log.Info(logrus.Fields{"request_id": requestID}, &result, "process service completed..")

	s.log.Info(logrus.Fields{"request_id": requestID}, paging, "get data product from db...")
	totalCount, products, err := s.productRepo.ListData(ctx, paging, opt)
	if err != nil {
		return
	}

	var rows []productModel.Product
	rows = append(rows, products...)

	// response
	result = &productModel.ProductList{
		Rows: rows,
		Pagination: helperModel.Pagination{
			Limit:      paging.Limit,
			Page:       paging.Page,
			TotalRows:  totalCount,
			TotalPages: int(math.Ceil(float64(totalCount) / float64(paging.Limit))),
		},
	}

	result.HasMorePage = true
	if paging.Page == result.TotalPages || result.TotalPages < 1 {
		result.HasMorePage = false
	}

	return
}
