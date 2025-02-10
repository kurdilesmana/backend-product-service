package deps

import (
	"log"

	"github.com/go-playground/validator/v10"
	config "github.com/kurdilesmana/backend-product-service/configs"
	"github.com/kurdilesmana/backend-product-service/internal/adapters/v1/repositories/healthCheckRepo"
	"github.com/kurdilesmana/backend-product-service/internal/adapters/v1/repositories/userRepo"
	"github.com/kurdilesmana/backend-product-service/internal/core/ports/healthCheckPort"
	"github.com/kurdilesmana/backend-product-service/internal/core/ports/userPort"
	"github.com/kurdilesmana/backend-product-service/internal/core/services/healthCheckService"
	"github.com/kurdilesmana/backend-product-service/internal/core/services/userService"
	"github.com/kurdilesmana/backend-product-service/internal/infra/db"
	"github.com/kurdilesmana/backend-product-service/pkg/logging"
)

const (
	keyTransaction = "kbs-ctx"
	timeout        = 60
)

type Dependency struct {
	Cfg                config.EnvironmentConfig
	HealthCheckService healthCheckPort.IHealthCheckService
	UserService        userPort.IUserService
	Validator          *validator.Validate
	Logger             *logging.Logger
}

func SetupDependencies() Dependency {
	validator := validator.New()

	// init config
	config, err := config.LoadENVConfig()
	if err != nil {
		log.Panic(err)
	}

	// load logger
	logger := logging.NewLogger("Product-Service")

	// BIG DEPENDENCY STAGE =======================================
	database, err := db.OpenPgsqlConnection(&config.Database)
	if err != nil {
		log.Panic(err)
	}

	// redis, err := db.RedisNewClient(&config.Cache)
	// if err != nil {
	// 	log.Panic(err)
	// }

	// BIG DEPENDENCY STAGE END =======================================

	// init repository
	healthCheckRepository := healthCheckRepo.NewHealthCheckRepo(database, keyTransaction, timeout)
	userRepository := userRepo.NewUserRepo(database, keyTransaction, timeout, logger)

	//init middleware

	// init service
	healthCheckService := healthCheckService.NewHealthCheckService(healthCheckRepository, *logger)
	userService := userService.NewUserService(userRepository, logger)

	return Dependency{
		Cfg:                config,
		HealthCheckService: healthCheckService,
		UserService:        userService,
		Validator:          validator,
		Logger:             logger,
	}
}
