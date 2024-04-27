package helper

import (
	"golang.org/x/crypto/bcrypt"
)

func CheckValidatePassword(dbPass string, inputPass string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(dbPass), []byte(inputPass)); err != nil {
		return err
	}
	return nil
}
