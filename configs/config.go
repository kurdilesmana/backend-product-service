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
	// Cache     Redis
	Token Token
	Log   logging.Logger
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
			Engine:          os.Getenv("DATABASE_ENGINE"),
			Host:            os.Getenv("DATABASE_HOST"),
			Port:            cast.ToInt(os.Getenv("DATABASE_PORT")),
			Username:        os.Getenv("DATABASE_USERNAME"),
			Password:        os.Getenv("DATABASE_PASSWORD"),
			DBName:          os.Getenv("DATABASE_NAME"),
			Schema:          os.Getenv("DATABASE_SCHEMA"),
			MaxIdle:         cast.ToInt(os.Getenv("DATABASE_MAX_IDLE")),
			MaxConn:         cast.ToInt(os.Getenv("DATABASE_MAX_CONN")),
			ConnMaxLifetime: cast.ToInt(os.Getenv("DATABASE_CONN_LIFETIME")),
		},
		Token: Token{
			Issuer:       os.Getenv("JWT_ISSUER"),
			AccessToken:  AccessToken{Secret: os.Getenv("JWT_ACCESS_SECRET"), Duration: os.Getenv("JWT_ACCESS_DURATION")},
			RefreshToken: RefreshToken{Secret: os.Getenv("JWT_REFRESH_SECRET"), Duration: os.Getenv("JWT_REFRESH_DURATION"), LongDuration: os.Getenv("JWT_REFRESH_LONG_DURATION")},
		},
		// Cache: Redis{
		// 	Host:         os.Getenv("REDIS_HOST"),
		// 	Port:         cast.ToInt(os.Getenv("REDIS_PORT")),
		// 	Username:     os.Getenv("REDIS_USERNAME"),
		// 	Password:     os.Getenv("REDIS_PASSWORD"),
		// 	DB:           cast.ToInt(os.Getenv("REDIS_DB")),
		// 	UseTLS:       cast.ToBool(os.Getenv("REDIS_USE_TLS")),
		// 	MaxRetries:   cast.ToInt(os.Getenv("REDIS_MAX_RETRIES")),
		// 	MinIdleConns: cast.ToInt(os.Getenv("REDIS_MIN_IDLE_CONNS")),
		// 	PoolSize:     cast.ToInt(os.Getenv("REDIS_POOL_SIZE")),
		// 	PoolTimeout:  cast.ToInt(os.Getenv("REDIS_POOL_TIMEOUT")),
		// 	MaxConnAge:   cast.ToInt(os.Getenv("REDIS_MAX_CONN_AGE")),
		// 	ReadTimeout:  cast.ToInt(os.Getenv("REDIS_READ_TIMEOUT")),
		// 	WriteTimeout: cast.ToInt(os.Getenv("REDIS_WRITE_TIMEOUT")),
		// },
	}

	return
}
