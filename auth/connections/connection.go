package connections

import (
	"context"

	"github.com/jmoiron/sqlx"
)

type DBConnection interface {
	Connect() error
	Ping() error
	Close() error

	BeginTransaction(ctx context.Context) (*sqlx.Tx, error)
	CommitTransaction(tx *sqlx.Tx) error
	RollbackTransaction(tx *sqlx.Tx) error

	CreateTable(tx *sqlx.Tx, schema string) error
	// Create()
}
