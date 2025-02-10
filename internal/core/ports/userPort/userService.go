package userPort

import (
	"context"

	"github.com/kurdilesmana/backend-product-service/internal/core/models/userModel"
)

type IUserService interface {
	CreateUser(ctx context.Context, request userModel.CreateUserRequest) (response userModel.CreateUserResponse, err error)
	Login(ctx context.Context, request userModel.LoginRequest) (response userModel.LoginResponse, err error)
}
