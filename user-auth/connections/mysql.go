package connections

import "github.com/jmoiron/sqlx"

type MySQLConnection struct {
	DSN string
	db  *sqlx.DB
}

func (sq *MySQLConnection) Connect() error {
	var err error
	sq.db, err = sqlx.Connect("mysql", sq.DSN)
	return err
}

func (sq *MySQLConnection) Ping() error {
	return sq.db.Ping()
}

func (sq *MySQLConnection) Close() error {
	return sq.db.Close()
}
