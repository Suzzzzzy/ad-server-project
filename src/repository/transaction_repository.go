package repository

import (
	"ad-server-project/src/domain/model"
	"context"
	"database/sql"
)

type TransactionRepository struct {
	db *sql.DB
}

func NewTransactionRepository(db *sql.DB) model.TransactionRepository {
	return &TransactionRepository{db: db}
}

func (t *TransactionRepository) Transaction(ctx context.Context,txFn func(tx *sql.Tx) error) error {
	tx, err := t.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	err = txFn(tx)
	if err != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			return rollbackErr
		}
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}
