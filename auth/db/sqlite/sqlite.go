package sqlite

import (
	"context"
	"embed"
	"fmt"

	"github.com/midir99/sqload"
	"github.com/subhamyadav580/user-auth/auth/connections"
	"github.com/subhamyadav580/user-auth/auth/db/common"
)

//go:embed migrations
var fsys embed.FS

var Q = sqload.MustLoadFromFS[struct {
	CreateUserTable   string `query:"CreateUserTable"`
	CreateUserProfile string `query:"CreateUserProfile"`
	CreateAccount     string `query:"CreateAccount"`
	CreateProfile     string `query:"CreateProfile"`
}](fsys)

func CreateTables(dbConn connections.DBConnection) {
	var err error
	err = dbConn.Connect()
	if err != nil {
		fmt.Println("Error while connecting sqlite3 db:", err.Error())
		return
	}
	defer dbConn.Close()
	ctx := context.Background()
	tx, err := dbConn.BeginTransaction(ctx)
	if err != nil {
		fmt.Println("Error while initiating transaction")
		return
	}
	err = dbConn.CreateTable(tx, Q.CreateUserTable)
	if err != nil {
		fmt.Println("Error while creating User table: ", err.Error())
		dbConn.RollbackTransaction(tx)
		return
	}
	err = dbConn.CreateTable(tx, Q.CreateUserProfile)
	if err != nil {
		fmt.Println("Error while creating UserProfile table: ", err.Error())
		dbConn.RollbackTransaction(tx)
		return
	}
	// Commit the transaction if no errors occurred
	err = dbConn.CommitTransaction(tx)
	if err != nil {
		fmt.Println("Error while committing transaction:", err.Error())
		dbConn.RollbackTransaction(tx)
		return
	}
	fmt.Println("Tables created successfully!")
}

func CreateUser(dbConn connections.DBConnection, AccountsData common.Account, ProfileData common.UserProfile) {
	var err error
	err = dbConn.Connect()
	if err != nil {
		fmt.Println("Error while connecting sqlite3 db:", err.Error())
		return
	}
	defer dbConn.Close()
	ctx := context.Background()
	tx, err := dbConn.BeginTransaction(ctx)
	if err != nil {
		fmt.Println("Error while initiating transaction")
		return
	}
	_, err = dbConn.Create(tx, Q.CreateAccount, AccountsData)
	if err != nil {
		fmt.Println("Error while inserting rows in Accounts table: ", err.Error())
		dbConn.RollbackTransaction(tx)
		return
	}
	_, err = dbConn.Create(tx, Q.CreateProfile, ProfileData)
	if err != nil {
		fmt.Println("Error while inserting rows in UserProfile table: ", err.Error())
		dbConn.RollbackTransaction(tx)
		return
	}
	// Commit the transaction if no errors occurred
	err = dbConn.CommitTransaction(tx)
	if err != nil {
		fmt.Println("Error while committing transaction:", err.Error())
		dbConn.RollbackTransaction(tx)
		return
	}
	fmt.Println("Profile created successfully!")

}
