package auth

import (
	"api/src/config"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func MakeToken(userID uint64) (string, error) {
	permissions := jwt.MapClaims{}
	permissions["authorized"] = true
	permissions["expires"] = time.Now().Add(time.Hour * 6).Unix()
	permissions["userId"] = userID

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissions)
	return token.SignedString(config.Secret)
}
func ValidateToken(r *http.Request) error {
	tokenString := extractToken(r)
	_, error := jwt.Parse(tokenString, getSecret)

	if error != nil {
		return error
	}
	return nil
}

func extractToken(r *http.Request) string {
	token := r.Header.Get("Authorization")
	if splited := strings.Split(token, " "); len(splited) == 2 {
		return splited[1]
	}
	return ""
}
func getSecret(token *jwt.Token) (any, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("Unexpected signing method! %v", token.Header["alg"])
	}
	return config.Secret, nil
}
