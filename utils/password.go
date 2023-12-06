package utils

import (
	"golang.org/x/crypto/bcrypt"
)

// GenerateHashPassword generates a hashed password from the given plaintext password.
func GenerateHashPassword(password string) (string, error) {
	// Generate the hashed password from the plaintext password using bcrypt with a cost factor of 14.
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	// Convert the hashed bytes to a string and return it along with any error that occurred.
	return string(hashedBytes), nil
}

// CheckPasswordHash checks if the given plaintext password matches the hashed password.
func CheckPasswordHash(password, hash string) bool {
	// Compare the plaintext password with the hashed password using bcrypt.
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

	// Return true if the comparison was successful, false otherwise.
	return err == nil
}
