package connections

import "github.com/jmoiron/sqlx"

type PostgresSQLConnection struct {
	DSN string
	db  *sqlx.DB
}

func (sq *PostgresSQLConnection) Connect() error {
	var err error
	sq.db, err = sqlx.Connect("postgres", sq.DSN)
	return err
}

func (sq *PostgresSQLConnection) Ping() error {
	return sq.db.Ping()
}

func (sq *PostgresSQLConnection) Close() error {
	return sq.db.Close()
}
