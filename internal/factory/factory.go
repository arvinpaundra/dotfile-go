package factory

import (
	"kompack-go-api/internal/repository"
	"kompack-go-api/pkg/database"
	"kompack-go-api/pkg/tracer"

	"go.uber.org/zap"
)

type Factory struct {
	TxBeginner repository.TxBeginner
	Logger     *zap.Logger
}

func NewFactory() *Factory {
	sql := database.GetConnection()

	return &Factory{
		// Pass the db connection to repository package for database query calling
		TxBeginner: repository.NewTxBeginner(sql),

		// logger
		Logger: tracer.Log,
	}
}
