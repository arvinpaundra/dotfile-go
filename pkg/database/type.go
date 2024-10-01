package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Define the database conn configuration
type (
	dbConfig struct {
		Host     string
		User     string
		Pass     string
		Port     string
		Name     string
		Sslmode  string
		TimeZone string
	}

	postgresConfig struct {
		dbConfig
	}
)

var err error

// Connect to postgres with the input configuration
func (conf postgresConfig) Connect() {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		conf.Host,
		conf.User,
		conf.Pass,
		conf.Name,
		conf.Port,
		conf.Sslmode,
		conf.TimeZone,
	)

	dbConn, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		PrepareStmt: true,
		Logger:      logger.Default,
	})

	if err != nil {
		panic(err)
	}
}
