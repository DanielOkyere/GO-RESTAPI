package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const secretKey = "supersecret"

func GenerateToken(email string, name string, userId int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userId": userId,
		"name":   name,
		"exp":    time.Now().Add(time.Hour * 3).Unix(),
	})
	return token.SignedString([]byte(secretKey))

}

func VerifyJWT(token string) error {
	parsedToken, err:=jwt.Parse(token, func(token *jwt.Token) (interface {}, error){
		_, ok :=token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(secretKey), nil
	})
	if err != nil {
		return  errors.New("could not parse token")
	}
	tokenIsValid := parsedToken.Valid
	if !tokenIsValid{
		return errors.New("invalid Token")
	}
	return nil
}