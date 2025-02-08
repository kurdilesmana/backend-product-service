package deps

import (
	"log"

	config "github.com/kurdilesmana/backend-product-service/configs"
	"github.com/kurdilesmana/backend-product-service/internal/infra/db"
	"github.com/kurdilesmana/backend-product-service/pkg/kbslog"
	"github.com/kurdilesmana/backend-product-service/pkg/kbsvalidator"
)

const (
	keyTransaction = "kbs-ctx"
	timeout        = 60
)

type Dependency struct {
	Cfg       config.EnvironmentConfig
	Validator kbsvalidator.Validator
	Logger    kbslog.Logger
}

func SetupDependencies() Dependency {
	validator := kbsvalidator.New()

	// init config
	config, err := config.LoadENVConfig()
	if err != nil {
		log.Panic(err)
	}

	// load logger
	logger := kbslog.New("info", "stdout")
	defer logger.Sync() // This script will be executed last
	defer logger.Info("Done cleanup tasks...")

	// BIG DEPENDENCY STAGE =======================================
	database, err := db.OpenPgsqlConnection(&config.KBSDatabase)
	if err != nil {
		log.Panic(err)
	}

	// redis, err := db.RedisNewClient(&config.Cache)
	// if err != nil {
	// 	log.Panic(err)
	// }

	// BIG DEPENDENCY STAGE END =======================================

	// init repository

	//init middleware

	// init service

	return Dependency{
		Cfg:    config,
		Logger: logger,
	}
}
