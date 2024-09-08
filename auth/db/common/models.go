package common

type Account struct {
	UserID       string `json:"user_id" db:"user_id"`
	Email        string `json:"email" db:"email"`
	Username     string `json:"username" db:"username"`
	HashPassword string `json:"hash_password" db:"hash_password"`
	UpdatedAt    string `json:"updated_at" db:"updated_at"`
	CreatedAt    string `json:"created_at" db:"created_at"`
}

type UserProfile struct {
	UserID     string `json:"user_id" db:"user_id"`
	Email      string `json:"email" db:"email"`
	Username   string `json:"username" db:"username"`
	FirstName  string `json:"first_name" db:"first_name"`
	MiddleName string `json:"middle_name" db:"middle_name"`
	LastName   string `json:"last_name" db:"last_name"`
	UpdatedAt  string `json:"updated_at" db:"updated_at"`
	CreatedAt  string `json:"created_at" db:"created_at"`
}
