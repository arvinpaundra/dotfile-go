package config

import (
	"os"

	"github.com/spf13/viper"
)

var C *Config

type (
	Postgres struct {
		Host     string
		Port     string
		User     string
		Password string
		Database string
		Timezone string
		Sslmode  string
	}

	Rabbitmq struct {
		Host     string
		Port     string
		User     string
		Password string
		Vhost    string
	}

	GoogleCloudStorage struct {
		BaseUrl        string
		Bucket         string
		CredentialPath string
	}

	Config struct {
		//	application configurations
		Port, GinMode string

		//	postgres connection
		Postgres Postgres

		//	gcs conf
		GCS GoogleCloudStorage

		//	rabbitmq conf
		Rabbitmq Rabbitmq
	}
)

func LoadEnv(filename, ext, path string) {
	viper.SetConfigName(filename)
	viper.SetConfigType(ext)
	viper.AddConfigPath(path)

	// omit the error, cause if the conf file is not exist, it will look up to os env
	_ = viper.ReadInConfig()

	C = &Config{
		Port:    getEnv("APP_PORT"),
		GinMode: getEnv("GIN_MODE"),
		Postgres: Postgres{
			Host:     getEnv("POSTGRES_HOST"),
			Port:     getEnv("POSTGRES_PORT"),
			User:     getEnv("POSTGRES_USER"),
			Password: getEnv("POSTGRES_PASS"),
			Database: getEnv("POSTGRES_DBNAME"),
			Timezone: getEnv("POSTGRES_TIMEZONE"),
			Sslmode:  getEnv("POSTGRES_SSLMODE"),
		},
		GCS: GoogleCloudStorage{
			BaseUrl:        getEnv("GOOGLE_CLOUD_BASE_URL"),
			Bucket:         getEnv("GOOGLE_CLOUD_BUCKET"),
			CredentialPath: getEnv("GOOGLE_CLOUD_CREDENTIALS_PATH"),
		},
		Rabbitmq: Rabbitmq{
			Host:     getEnv("RABBITMQ_HOST"),
			Port:     getEnv("RABBITMQ_PORT"),
			User:     getEnv("RABBITMQ_USER"),
			Password: getEnv("RABBITMQ_PASSWORD"),
			Vhost:    getEnv("RABBITMQ_VHOST"),
		},
	}
}

func getEnv(key string) string {
	v := viper.GetString(key)

	if v == "" {
		v = os.Getenv(key)
	}

	return v
}
