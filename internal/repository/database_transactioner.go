package repository

import (
	"database/sql"

	"gorm.io/gorm"
)

type TxBeginner interface {
	// Begin begin a database transaction
	Begin(opts ...*sql.TxOptions) (Transactioner, error)
}

// Transaction transaction is an interface that should be implemented with a sql transaction
type Transactioner interface {
	// Commit commit a transaction
	Commit() error

	// Rollback rollback a transaction
	Rollback() error
}

type txBeginner struct {
	db *gorm.DB
}

func NewTxBeginner(db *gorm.DB) TxBeginner {
	return &txBeginner{db: db}
}

func (t *txBeginner) Begin(opts ...*sql.TxOptions) (Transactioner, error) {
	tx := t.db.Begin(opts...)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &transactioner{tx: tx}, nil
}

type transactioner struct {
	tx *gorm.DB
}

func (t *transactioner) Commit() error {
	err := t.tx.Commit().Error
	if err != nil {
		return err
	}

	return nil
}

func (t *transactioner) Rollback() error {
	err := t.tx.Rollback().Error
	if err != nil {
		return err
	}

	return nil
}
