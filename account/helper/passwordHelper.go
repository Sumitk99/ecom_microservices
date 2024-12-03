package helper

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"log"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)

	if err != nil {
		log.Println(err)
		return "", err
	}

	return string(bytes), nil
}

func VerifyPassword(hashedPassword string, providedPassword string) (bool, string) {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(providedPassword))
	check := true
	var msg string

	if err != nil {
		log.Println(err)
		check = false
		msg = fmt.Sprintf("email or password is incorrect")
	}
	return check, msg
}
