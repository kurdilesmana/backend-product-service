package userPort

import (
	"context"

	"github.com/kurdilesmana/backend-product-service/internal/core/models/userModel"
)

type IUserRepository interface {
	CreateUser(ctx context.Context, input userModel.User) (err error)
	CheckUserExist(ctx context.Context, email, phoneNumber string) (exist bool, err error)
	GetByEmailPhoneNumber(ctx context.Context, email, phoneNumber string) (account *userModel.User, err error)

	// SetToken(secretKey, issuerKey string, duration time.Duration) error
	// CreateToken(expiredPayload userModel.ExpiredPayload) (string, error)
	// VerifyToken(token string) (*userModel.Payload, error)
}
