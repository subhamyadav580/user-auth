package connections

import "github.com/jmoiron/sqlx"

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
