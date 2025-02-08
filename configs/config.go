package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"

	"github.com/kurdilesmana/backend-product-service/pkg/constants"
	"github.com/kurdilesmana/backend-product-service/pkg/logging"
)

type EnvironmentConfig struct {
	Env       string
	AppConfig AppConfig
	Database  Database
	Cache     Redis
	Log       logging.Logger
}

type AppConfig struct {
	Name           string
	Version        string
	Port           int
	MaxRequestTime int
}

type Log struct {
	Path      string
	Prefix    string
	Extension string
}

type Auth struct {
	Token
	Otp
	Link
}

type Link struct {
	Duration string
}

type Otp struct {
	Duration string
}

type Token struct {
	Issuer string
	AccessToken
	RefreshToken
}

type AccessToken struct {
	Secret   string
	Duration string
}

type RefreshToken struct {
	Secret       string
	Duration     string
	LongDuration string
}

func LoadENVConfig() (config EnvironmentConfig, err error) {
	err = godotenv.Load(".env")
	if err != nil {
		err = fmt.Errorf(constants.ErrLoadENV, err)
		return EnvironmentConfig{}, err
	}

	port, err := strconv.Atoi(os.Getenv("APP_PORT"))
	if err != nil {
		err = fmt.Errorf(constants.ErrConvertStringToInt, err)
		return EnvironmentConfig{}, err
	}

	config = EnvironmentConfig{
		Env: os.Getenv("ENV"),
		AppConfig: AppConfig{
			Name:           os.Getenv("APP_NAME"),
			Version:        os.Getenv("APP_VERSION"),
			Port:           port,
			MaxRequestTime: cast.ToInt(os.Getenv("APP_MAX_REQUEST_TIME")),
		},
		Database: Database{
			Engine:          os.Getenv("_DATABASE_ENGINE"),
			Host:            os.Getenv("_DATABASE_HOST"),
			Port:            cast.ToInt(os.Getenv("_DATABASE_PORT")),
			Username:        os.Getenv("_DATABASE_USERNAME"),
			Password:        os.Getenv("_DATABASE_PASSWORD"),
			DBName:          os.Getenv("_DATABASE_NAME"),
			Schema:          os.Getenv("_DATABASE_SCHEMA"),
			MaxIdle:         cast.ToInt(os.Getenv("_DATABASE_MAX_IDLE")),
			MaxConn:         cast.ToInt(os.Getenv("_DATABASE_MAX_CONN")),
			ConnMaxLifetime: cast.ToInt(os.Getenv("_DATABASE_CONN_LIFETIME")),
		},
		Cache: Redis{
			Host:         os.Getenv("REDIS_HOST"),
			Port:         cast.ToInt(os.Getenv("REDIS_PORT")),
			Username:     os.Getenv("REDIS_USERNAME"),
			Password:     os.Getenv("REDIS_PASSWORD"),
			DB:           cast.ToInt(os.Getenv("REDIS_DB")),
			UseTLS:       cast.ToBool(os.Getenv("REDIS_USE_TLS")),
			MaxRetries:   cast.ToInt(os.Getenv("REDIS_MAX_RETRIES")),
			MinIdleConns: cast.ToInt(os.Getenv("REDIS_MIN_IDLE_CONNS")),
			PoolSize:     cast.ToInt(os.Getenv("REDIS_POOL_SIZE")),
			PoolTimeout:  cast.ToInt(os.Getenv("REDIS_POOL_TIMEOUT")),
			MaxConnAge:   cast.ToInt(os.Getenv("REDIS_MAX_CONN_AGE")),
			ReadTimeout:  cast.ToInt(os.Getenv("REDIS_READ_TIMEOUT")),
			WriteTimeout: cast.ToInt(os.Getenv("REDIS_WRITE_TIMEOUT")),
		},
	}

	return
}
