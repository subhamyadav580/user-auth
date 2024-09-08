package utils

import (
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func GetCurrentTime() string {
	now := time.Now()
	formattedTime := now.Format("2006-01-02 15:04:05")
	fmt.Println("Formatted Time:", formattedTime)
	return formattedTime
}
