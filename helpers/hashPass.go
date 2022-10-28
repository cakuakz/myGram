package helpers

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(pass string) string {
	salt := bcrypt.DefaultCost
	password := []byte(pass)
	hashedPassword, _ := bcrypt.GenerateFromPassword(password, salt)

	return string(hashedPassword)
}

func ComparePassword(h, p []byte) bool {
	hash, pass := []byte(h), []byte(p)

	err := bcrypt.CompareHashAndPassword(hash, pass)

	return err == nil
}