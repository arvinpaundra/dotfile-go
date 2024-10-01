package factory

import (
	"github.com/arvinpaundra/dotfile-go/internal/repository"
	"github.com/arvinpaundra/dotfile-go/pkg/database"
	"github.com/arvinpaundra/dotfile-go/pkg/tracer"

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
