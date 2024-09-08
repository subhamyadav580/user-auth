package connections

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

type Sqlite3Connection struct {
	DSN string
	db  *sqlx.DB
}

func (sq *Sqlite3Connection) Connect() error {
	var err error
	sq.db, err = sqlx.Connect("sqlite3", sq.DSN)
	return err
}

func (sq *Sqlite3Connection) Ping() error {
	return sq.db.Ping()
}

func (sq *Sqlite3Connection) Close() error {
	return sq.db.Close()
}

func (sq *Sqlite3Connection) BeginTransaction(ctx context.Context) (*sqlx.Tx, error) {
	return sq.db.BeginTxx(ctx, nil)
}

func (sq *Sqlite3Connection) CommitTransaction(tx *sqlx.Tx) error {
	return tx.Commit()
}

func (sq *Sqlite3Connection) RollbackTransaction(tx *sqlx.Tx) error {
	return tx.Rollback()
}

func (sq *Sqlite3Connection) CreateTable(tx *sqlx.Tx, schema string) error {
	fmt.Println("Query to execute: ", schema)
	_, err := tx.Exec(schema)
	return err
}

// func (sq *Sqlite3Connection) Create(tableName string, data interface{}) {
// 	// sq.db.Exec()
// }
