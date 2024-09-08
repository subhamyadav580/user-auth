package auth

import (
	"fmt"

	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
	"github.com/subhamyadav580/user-auth/auth/connections"
	"github.com/subhamyadav580/user-auth/auth/db/common"
	"github.com/subhamyadav580/user-auth/auth/db/sqlite"
	"github.com/subhamyadav580/user-auth/utils"
)

func SetupTables(userInput utils.UserInput) {
	sqliteConn := &connections.Sqlite3Connection{DSN: "./example_database.db"}
	sqlite.CreateTables(sqliteConn)
	userID := uuid.New()
	hashPassword, err := utils.HashPassword(userInput.Password)
	if err != nil {
		fmt.Println("error while hashing password: ", err.Error())
		return
	}
	CurrentTime := utils.GetCurrentTime()
	accountData := common.Account{
		UserID:       userID.String(),
		Email:        userInput.Email,
		Username:     userInput.Username,
		HashPassword: hashPassword,
		CreatedAt:    CurrentTime,
		UpdatedAt:    CurrentTime,
	}
	profileData := common.UserProfile{
		UserID:     userID.String(),
		Email:      userInput.Email,
		Username:   userInput.Username,
		FirstName:  userInput.FirstName,
		MiddleName: userInput.MiddleName,
		LastName:   userInput.LastName,
		CreatedAt:  CurrentTime,
		UpdatedAt:  CurrentTime,
	}
	fmt.Println("Account Data:: ", accountData)
	fmt.Println("Profile Data:: ", profileData)
	sqlite.CreateUser(sqliteConn, accountData, profileData)
}
