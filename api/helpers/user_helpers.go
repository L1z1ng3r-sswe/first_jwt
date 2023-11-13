package helpers

import "golang.org/x/crypto/bcrypt"

func PasswordHasher(password string) (string, error) {
	hashedPassword , err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err!= nil {
		return "", err
	}

	return string(hashedPassword), nil
}