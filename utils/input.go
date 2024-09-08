package utils

type UserInput struct {
	Email      string `json:"email" db:"email"`
	Username   string `json:"username" db:"username"`
	FirstName  string `json:"first_name" db:"first_name"`
	MiddleName string `json:"middle_name" db:"middle_name"`
	LastName   string `json:"last_name" db:"last_name"`
	Password   string `json:"password" db:"password"`
}
