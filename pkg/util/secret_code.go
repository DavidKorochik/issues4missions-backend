package util

import "golang.org/x/crypto/bcrypt"

func HashSecretCode(secretCode string) (string, error) {
	hashedSecretCode, err := bcrypt.GenerateFromPassword([]byte(secretCode), bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}

	return string(hashedSecretCode), nil
}

func CheckSecretCode(secretCode string, hashedSecretCode string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedSecretCode), []byte(secretCode))

	if err != nil {
		return err
	}

	return nil
}
