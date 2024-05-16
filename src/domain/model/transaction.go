package model

import (
	"context"
	"database/sql"
)

type TransactionRepository interface {
	Transaction(ctx context.Context, txFn func(*sql.Tx) error) error
}
