package database

import (
	"log"
	"sync"

	config "github.com/arvinpaundra/dotfile-go/config"

	"gorm.io/gorm"
)

var (
	dbConn *gorm.DB
	once   sync.Once
)

func CreateConnection() {
	// Create database configuration information
	conf := dbConfig{
		User:     config.C.Postgres.User,
		Pass:     config.C.Postgres.Password,
		Host:     config.C.Postgres.Host,
		Port:     config.C.Postgres.Port,
		Name:     config.C.Postgres.Database,
		Sslmode:  config.C.Postgres.Sslmode,
		TimeZone: config.C.Postgres.Timezone,
	}

	postgres := postgresConfig{dbConfig: conf}
	// Create only one mysql Connection, not the same as mysql TCP connection
	once.Do(func() {
		postgres.Connect()
	})

	log.Println("connected to postgres")
}

func GetConnection() *gorm.DB {
	// Check db connection, if exist return the memory address of the db connection
	if dbConn == nil {
		CreateConnection()
	}
	return dbConn
}

func Close(db *gorm.DB) error {
	conn, err := db.DB()
	if err != nil {
		log.Printf("failed to get database connection: %e", err)
		return err
	}

	err = conn.Close()
	if err != nil {
		log.Printf("failed to close database: %e", err)
		return err
	}

	return nil
}
