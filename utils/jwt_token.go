package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const TokenSecret = "jwt"

func GenerateToken(id int64, email string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{"id": id, "email": email, "exp": time.Now().Add(2 * time.Hour).Unix()})
	return token.SignedString([]byte(TokenSecret))
}
