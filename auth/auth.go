package auth

import (
	_ "github.com/mattn/go-sqlite3"
	"github.com/subhamyadav580/user-auth/auth/connections"
	"github.com/subhamyadav580/user-auth/auth/db/sqlite"
)

func SetupTables() {
	var (
	// pgConn     *connections.PostgresSQLConnection
	// sqliteConn *connections.Sqlite3Connection
	)
	// pgConn = &connections.PostgresSQLConnection{DSN: ""}
	sqliteConn := &connections.Sqlite3Connection{DSN: "./example_database.db"}
	sqlite.CreateTables(sqliteConn)
}
