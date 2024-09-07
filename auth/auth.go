package auth

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"github.com/subhamyadav580/user-auth/auth/db/sqlite"
)

func Signup() {
	fmt.Println("hELOOOO")
	sql_instance, err := sqlx.Open("sqlite3", "./example_database.db")
	if err != nil {
		fmt.Println("Error while connecting database", err.Error())
	} else {
		defer sql_instance.Close()
		fmt.Println("sqlInstance", sql_instance)
		sqlite.CreateTables()
	}
}
