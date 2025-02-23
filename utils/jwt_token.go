package utils

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const TokenSecret = "jwt"

func GenerateToken(id int64, email string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{"userId": id, "email": email, "exp": time.Now().Add(2 * time.Hour).Unix()})
	return token.SignedString([]byte(TokenSecret))
}

func ValidToken(token string) (int64, error) {
	parsedToken, error := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(TokenSecret), nil

	})
	if error != nil {
		return 0, error
	}
	var id int64
	if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok {
		id = int64(claims["userId"].(float64))
	}
	return id, nil
}
