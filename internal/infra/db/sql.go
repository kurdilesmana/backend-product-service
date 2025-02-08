package db

import (
	"fmt"
	"log"
	"time"

	config "github.com/kurdilesmana/backend-product-service/configs"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func OpenPgsqlConnection(dbConfig *config.KBSDatabase) (*gorm.DB, error) {

	var dsn string

	if dbConfig.KBSPassword != "" {
		dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
			dbConfig.KBSHost,
			dbConfig.KBSUsername,
			dbConfig.KBSPassword,
			dbConfig.KBSDBName,
			dbConfig.KBSPort,
		)
	} else {
		dsn = fmt.Sprintf("host=%s user=%s dbname=%s port=%d sslmode=disable",
			dbConfig.KBSHost,
			dbConfig.KBSUsername,
			dbConfig.KBSDBName,
			dbConfig.KBSPort,
		)
	}

	gormConfig := &gorm.Config{}

	db, err := gorm.Open(postgres.Open(dsn), gormConfig)
	if err != nil {
		log.Fatalf("Failed to connect to PostgreSQL: %v", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Failed to access database connection: %v", err)
	}

	if err := sqlDB.Ping(); err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}

	sqlDB.SetMaxIdleConns(dbConfig.KBSMaxIdle)
	sqlDB.SetMaxOpenConns(dbConfig.KBSMaxConn)
	sqlDB.SetConnMaxLifetime(time.Second * time.Duration(dbConfig.KBSConnMaxLifetime))
	return db, nil
}
