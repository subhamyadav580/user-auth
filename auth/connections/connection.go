package connections

import (
	"context"
	"database/sql"

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
	Create(tx *sqlx.Tx, query string, data interface{}) (sql.Result, error)
}
