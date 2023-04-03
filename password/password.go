// Package password provide some utils to help with passwords
package password

import (
	"golang.org/x/crypto/bcrypt"
)

// PasswordHash return hashed password
func PasswordHash(password string) (string, error) {
	newPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(newPassword), err
}

// PasswordCheck check if hashed password and password are egale
func PasswordCheck(hashedPassword string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
