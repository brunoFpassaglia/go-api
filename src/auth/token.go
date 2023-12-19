package auth

import (
	"api/src/config"
	"errors"
	"fmt"
	"net/http"
	"strconv"
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
	token, error := jwt.Parse(tokenString, getSecret)

	if error != nil {
		return error
	}

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return nil
	}
	return errors.New("invalid token")
}

func ExtractUserId(r *http.Request) (uint64, error) {

	tokenString := extractToken(r)
	token, error := jwt.Parse(tokenString, getSecret)

	if error != nil {
		return 0, error
	}

	if permissions, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		id, error := strconv.ParseUint(fmt.Sprintf("%.0f", permissions["userId"]), 10, 64)
		if error != nil {
			return 0, error
		}
		return id, nil
	}

	return 0, errors.New("Invalid token, user not found")

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
		return nil, fmt.Errorf("unexpected signing method! %v", token.Header["alg"])
	}
	return config.Secret, nil
}
