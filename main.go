package main

import (
	"fmt"

	auth "github.com/subhamyadav580/user-auth/auth"
	"github.com/subhamyadav580/user-auth/utils"
)

func main() {
	var (
		first_name  string
		last_name   string
		middle_name string
		email       string
		// username    string
		password string
	)
	fmt.Print("Enter First Name: ")
	fmt.Scanf("%s", &first_name)

	fmt.Print("Enter Middle Name: ")
	fmt.Scanf("%s", &middle_name)

	fmt.Print("Enter Last Name: ")
	fmt.Scanf("%s", &last_name)

	fmt.Print("Enter Email: ")
	fmt.Scanf("%s", &email)

	// fmt.Print("Enter Username: ")
	// fmt.Scanf("%s", &username)

	fmt.Print("Enter Password: ")
	fmt.Scanf("%s", &password)
	userInput := utils.UserInput{
		FirstName:  first_name,
		LastName:   last_name,
		MiddleName: middle_name,
		Email:      email,
		// Username:   username,
		Password: password,
	}
	fmt.Println("User Inputs: ", userInput)
	auth.SetupTables(userInput)
}
