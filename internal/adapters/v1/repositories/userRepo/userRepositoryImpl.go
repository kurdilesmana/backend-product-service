package userRepo

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/kurdilesmana/backend-product-service/internal/core/models/userModel"
	"github.com/kurdilesmana/backend-product-service/internal/core/ports/userPort"
	"github.com/kurdilesmana/backend-product-service/pkg/logging"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type userRepository struct {
	DB             *gorm.DB
	KeyTransaction string
	timeout        time.Duration
	log            *logging.Logger
	IssuerKey      string
	SecretKey      string
	Duration       time.Duration
}

func NewUserRepo(db *gorm.DB, keyTransaction string, timeout int, logger *logging.Logger) userPort.IUserRepository {
	return &userRepository{
		DB:             db,
		KeyTransaction: keyTransaction,
		timeout:        time.Duration(timeout) * time.Second,
		log:            logger,
	}
}

func (r *userRepository) CreateUser(ctx context.Context, input userModel.User) (err error) {
	r.log.Info(logrus.Fields{}, input, "start create user...")
	trx, ok := ctx.Value(r.KeyTransaction).(*gorm.DB)
	if !ok {
		trx = r.DB
	}

	ctxWT, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()

	res := trx.WithContext(ctxWT).Create(&input)
	if res.Error != nil {
		remark := "failed to create rekening"
		r.log.Error(logrus.Fields{
			"error": res.Error.Error(),
		}, input, remark)
		err = fmt.Errorf(remark)
		return
	}

	return
}

func (r *userRepository) CheckUserExist(ctx context.Context, Email, PhoneNumber string) (exist bool, err error) {
	r.log.Info(logrus.Fields{"PhoneNumber": PhoneNumber, "Email": Email}, nil, "check user exist...")
	trx, ok := ctx.Value(r.KeyTransaction).(*gorm.DB)
	if !ok {
		trx = r.DB
	}

	ctxWT, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()

	var user userModel.User
	res := trx.WithContext(ctxWT).Where("email = ? or phone_number = ?", Email, PhoneNumber).Find(&user)
	if res.Error != nil {
		remark := "failed to get user by email or phone number"
		r.log.Error(logrus.Fields{
			"error": res.Error.Error(),
		}, nil, remark)
		err = fmt.Errorf(remark)
		return
	}

	if res.RowsAffected > 0 {
		exist = true
	}

	return
}

func (r *userRepository) GetByEmailPhoneNumber(ctx context.Context, email, phoneNumber string) (*userModel.User, error) {
	r.log.Info(logrus.Fields{"email": email, "phoneNumber": phoneNumber}, nil, "get user by email phone number...")
	trx, ok := ctx.Value(r.KeyTransaction).(*gorm.DB)
	if !ok {
		trx = r.DB
	}

	ctxWT, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()

	var user userModel.User
	res := trx.WithContext(ctxWT).Where("email = ? or phone_number = ?", email, phoneNumber).First(&user)
	if res.Error != nil {
		remark := "failed to get user by email phone number"
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			remark = "user not found"
		}

		r.log.Error(logrus.Fields{
			"error": res.Error.Error(),
		}, nil, remark)
		return nil, fmt.Errorf(remark)
	}

	// Periksa apakah ID user kosong
	if user.ID == 0 {
		return nil, fmt.Errorf("user not found")
	}

	return &user, nil
}

func (r *userRepository) SetToken(secretKey, issuerKey string, duration time.Duration) error {
	r.SecretKey = secretKey
	r.IssuerKey = issuerKey
	r.Duration = duration

	return nil
}

/* func (r *userRepository) CreateToken(expiredPayload userModel.ExpiredPayload, sessionPayload userModel.SessionPayload, userPayload userModel.UserPayload) (token string, err error) {
	newPayload, err := userModel.NewPayload(expiredPayload, sessionPayload, userPayload, r.Duration)
	if err != nil {
		return "", err
	}

	tokenWithClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, newPayload)
	signedToken, err := tokenWithClaims.SignedString([]byte(r.SecretKey))
	if err != nil {
		return "", err
	}

	return signedToken, nil
} */
