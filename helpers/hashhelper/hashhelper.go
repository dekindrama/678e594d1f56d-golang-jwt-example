package hashhelper

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func GenerateHashString(value string) (string, error) {
	hashedValue, err := bcrypt.GenerateFromPassword([]byte(value), bcrypt.DefaultCost)
	return string(hashedValue), err
}

func CompareHashString(value, hashedValue string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedValue), []byte(value))
	if err != nil {
		fmt.Printf("compareHashString: %v", err.Error())
	}

	return err == nil
}
