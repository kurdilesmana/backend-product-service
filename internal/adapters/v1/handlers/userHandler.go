package handlers

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/kurdilesmana/backend-product-service/internal/core/models/userModel"
	"github.com/kurdilesmana/backend-product-service/internal/core/ports/userPort"
	"github.com/kurdilesmana/backend-product-service/pkg/logging"
	"github.com/kurdilesmana/backend-product-service/pkg/middleware"
	"github.com/kurdilesmana/backend-product-service/pkg/web"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type UserHandler struct {
	UserService userPort.IUserService
	Logger      *logging.Logger
	Validator   *validator.Validate
}

func NewUserHandler(
	userService userPort.IUserService,
	logger *logging.Logger,
	validator *validator.Validate,
) UserHandler {
	return UserHandler{
		UserService: userService,
		Logger:      logger,
		Validator:   validator,
	}
}

// CreateHandler godoc
// @Summary API Create User
// @Description Create For User
// @Tags		User
// @Accept	json
// @Produce	json
// @Param CreateUserRequest body userModel.CreateUserRequest true "Request Parameters"
// @Success 200 {object} userModel.CreateUserResponse "Response Success"
// @Router		/user/register	[post]
func (h *UserHandler) CreateHandler(ctx echo.Context) error {
	requestID := middleware.GetID(ctx)
	userCtx := middleware.SetIDx(ctx.Request().Context(), requestID)

	var payload userModel.CreateUserRequest
	if err := ctx.Bind(&payload); err != nil {
		h.Logger.Error(logrus.Fields{"request_id": requestID, "error": err.Error()}, payload, "create user payload")
		return web.ResponseFormatter(ctx, http.StatusBadRequest, "Bad Request", nil, err)
	}

	// Validate slice payload
	err := h.Validator.Struct(payload)
	if err != nil {
		h.Logger.Error(logrus.Fields{"request_id": requestID, "error": err.Error()}, payload, "create user payload validation")
		return web.ResponseErrValidationWithCode(ctx, "bad request", err, http.StatusBadRequest)
	}

	respData, err := h.UserService.CreateUser(userCtx, payload)
	if err != nil {
		h.Logger.Error(logrus.Fields{"request_id": requestID, "error": err.Error()}, payload, "error create user")
		return web.ResponseFormatter(ctx, http.StatusBadRequest, err.Error(), nil, err)
	}

	return web.ResponseFormatter(ctx, http.StatusCreated, "Success", respData, nil)
}

// LoginHandler godoc
// @Summary API Login User
// @Description Login For User
// @Tags		User
// @Accept	json
// @Produce	json
// @Param LoginRequest body userModel.LoginRequest true "Request Parameters"
// @Success 200 {object} userModel.CreateUserResponse "Response Success"
// @Router		/user/login	[post]
func (h *UserHandler) LoginHandler(ctx echo.Context) error {
	requestID := middleware.GetID(ctx)
	userCtx := middleware.SetIDx(ctx.Request().Context(), requestID)

	var payload userModel.LoginRequest
	if err := ctx.Bind(&payload); err != nil {
		h.Logger.Error(logrus.Fields{"request_id": requestID, "error": err.Error()}, payload, "create user payload")
		return web.ResponseFormatter(ctx, http.StatusBadRequest, "Bad Request", nil, err)
	}

	// Validate slice payload
	err := h.Validator.Struct(payload)
	if err != nil {
		h.Logger.Error(logrus.Fields{"request_id": requestID, "error": err.Error()}, payload, "create user payload validation")
		return web.ResponseErrValidationWithCode(ctx, "bad request", err, http.StatusBadRequest)
	}

	respData, err := h.UserService.Login(userCtx, payload)
	if err != nil {
		h.Logger.Error(logrus.Fields{"request_id": requestID, "error": err.Error()}, payload, "error create user")
		return web.ResponseFormatter(ctx, http.StatusBadRequest, err.Error(), nil, err)
	}

	return web.ResponseFormatter(ctx, http.StatusCreated, "Success", respData, nil)
}
