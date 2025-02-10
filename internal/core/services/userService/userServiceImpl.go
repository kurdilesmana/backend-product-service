package userService

import (
	"context"
	"fmt"
	"strings"

	config "github.com/kurdilesmana/backend-product-service/configs"
	"github.com/kurdilesmana/backend-product-service/internal/core/models/userModel"
	"github.com/kurdilesmana/backend-product-service/internal/core/ports/userPort"
	"github.com/kurdilesmana/backend-product-service/pkg/convert"
	"github.com/kurdilesmana/backend-product-service/pkg/hash"
	"github.com/kurdilesmana/backend-product-service/pkg/logging"
	"github.com/kurdilesmana/backend-product-service/pkg/middleware"
	"github.com/sirupsen/logrus"
)

type userService struct {
	UserRepo userPort.IUserRepository
	Config   config.EnvironmentConfig
	log      *logging.Logger
}

func NewUserService(
	userRepo userPort.IUserRepository,
	config config.EnvironmentConfig,
	logger *logging.Logger,
) userPort.IUserService {
	return &userService{
		UserRepo: userRepo,
		Config:   config,
		log:      logger,
	}
}

func (s *userService) CreateUser(ctx context.Context, request userModel.CreateUserRequest) (response userModel.CreateUserResponse, err error) {
	// get request id
	requestID := middleware.GetIDx(ctx)
	defer s.log.Info(logrus.Fields{"request_id": requestID}, request, "process service completed..")

	// format phoneno
	phoneNumber := convert.NormalizePhoneNumber(request.PhoneNumber)

	// validate user exist
	isExist, err := s.UserRepo.CheckUserExist(ctx, request.Email, phoneNumber)
	if err != nil {
		err = fmt.Errorf("failed to check exist user")
		s.log.Warn(logrus.Fields{}, nil, err.Error())
		return
	}

	if isExist {
		err = fmt.Errorf("failed to create, user with email %s or phone number %s already exist", request.Email, request.PhoneNumber)
		s.log.Warn(logrus.Fields{}, request, err.Error())
		return
	}

	// Hash PIN
	hashedPassword, err := hash.HashPassword(request.Password)
	if err != nil {
		err = fmt.Errorf("failed to hashed PIN")
		s.log.Warn(logrus.Fields{}, nil, err.Error())
		return
	}

	// create user
	user := userModel.User{
		Name:        request.Name,
		Email:       request.Email,
		PhoneNumber: phoneNumber,
		Password:    string(hashedPassword),
	}
	s.log.Info(logrus.Fields{"request_id": requestID}, user, "Create user to db")
	if err = s.UserRepo.CreateUser(ctx, user); err != nil {
		err = fmt.Errorf("failed to create user")
		s.log.Warn(logrus.Fields{}, nil, err.Error())
		return
	}

	// response
	response = userModel.CreateUserResponse{
		UserPayload: userModel.UserPayload{
			Name:        request.Name,
			Email:       request.Email,
			PhoneNumber: request.PhoneNumber,
		},
	}

	return
}

func (s *userService) Login(ctx context.Context, request userModel.LoginRequest) (response userModel.LoginResponse, err error) {
	// get request id
	requestID := middleware.GetIDx(ctx)
	defer s.log.Info(logrus.Fields{"request_id": requestID}, request, "process service completed..")

	userName := request.Username
	if strings.HasPrefix(userName, "0") {
		userName = convert.NormalizePhoneNumber(request.Username)
	}

	userData, err := s.UserRepo.GetByEmailPhoneNumber(ctx, userName, userName)
	if err != nil {
		s.log.Warn(logrus.Fields{}, nil, err.Error())
		return
	}

	s.log.Info(logrus.Fields{}, userData, "user data")
	if !hash.CheckPasswordHash(request.Password, userData.Password) {
		remark := "password did not matched"
		s.log.Warn(logrus.Fields{}, nil, remark)
		err = fmt.Errorf(remark)
		return
	}

	/* var accessToken string
	var expAccessToken int64

	// Generate access token
	issuerKey := s.Config.Token.Issuer
	secretKey := s.Config.Token.AccessToken.Secret
	accessTokenDurationStr := s.Config.Token.AccessToken.Duration
	atTokenDuration, err := convert.ParseStringToDuration(accessTokenDurationStr)
	if err != nil {
		return "", "", 0, 0, errr.ErrCreateAccessToken
	}

	s.UserRepo.SetToken(secretKey, issuerKey, atTokenDuration)
	expAccessToken = time.Now().Add(atTokenDuration).Unix()
	expiredAccessToken := userModel.ExpiredPayload{
		Exp: expAccessToken,
	}
	accessToken, err = s.UserRepo.CreateToken(expiredAccessToken)
	if err != nil {
		return "", "", 0, 0, errr.ErrCreateAccessToken
	}

	response = userModel.LoginResponse{
		TokenType:   "bearer",
		AccessToken: accessToken,
	} */

	return
}
