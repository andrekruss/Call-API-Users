package helpers

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	salt, err := bcrypt.GenerateFromPassword([]byte(password), 10)

	if err != nil {
		return "Password hashing error.", err
	}

	hashedPassword := string(salt)
	return hashedPassword, nil
}

func CheckPassword(password string, hashedPassword string) bool {

	passwordBytes := []byte(password)
	hashedPasswordBytes := []byte(hashedPassword)

	err := bcrypt.CompareHashAndPassword(hashedPasswordBytes, passwordBytes)

	return err == nil
}
