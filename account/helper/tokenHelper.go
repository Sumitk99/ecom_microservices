package helper

import (
	"errors"
	"fmt"

	jwt "github.com/dgrijalva/jwt-go"
	"log"
	"os"
	"time"
)

type SignedDetails struct {
	Email    string
	Name     string
	UserID   string
	Phone    string
	UserType string
	jwt.StandardClaims
}

var SECRET_KEY = []byte(os.Getenv("SECRET_KEY"))

func GenerateTokens(Name, Email, Phone, UserType, ID string) (singedToken string, singedRefreshToken string, err error) {
	claims := &SignedDetails{
		Email:    Email,
		Name:     Name,
		UserID:   ID,
		UserType: UserType,
		Phone:    Phone,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(24)).Unix(),
		},
	}

	refreshClaims := &SignedDetails{ // used to get a new token if a token expires
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(168)).Unix(),
		},
	}

	singedToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(SECRET_KEY)

	singedRefreshToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString(SECRET_KEY)
	if err != nil {
		return
	}
	return
}

func ValidateToken(signedToken string) (*SignedDetails, error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&SignedDetails{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(SECRET_KEY), nil
		},
	)

	if err != nil {
		log.Println(err)
		return nil, errors.New(fmt.Sprintf("Invalid token: %v", err))
	}

	claims, ok := token.Claims.(*SignedDetails)
	if !ok {
		log.Println("Token is invalid")
		return nil, errors.New("Token is invalid")
	}
	if claims.ExpiresAt < time.Now().Local().Unix() {
		log.Println("Token is expired")
		return nil, errors.New("Token is expired")
	}
	return claims, nil
}
