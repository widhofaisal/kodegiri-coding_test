package utils

import (
	"time"
	
	"github.com/dgrijalva/jwt-go"

	"kodegiri/kodingtest/constant"
)

func Create_token(userId uint, username string, role string) (string, error) {
	// create the claims
	claims := jwt.MapClaims{}
	claims["user_id"] = userId
	claims["username"] = username
	claims["role"] = role
	claims["exp"] = time.Now().AddDate(0, 0, 7).Unix()

	// Create the token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign and get the complete encoded token as a string
	tokenString, err := token.SignedString([]byte(constant.SECRET_JWT))

	if err != nil {
		return "broken_token", err
	}

	return tokenString, nil
}



