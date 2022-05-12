package jwt

import (
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/form3tech-oss/jwt-go"
	"github.com/golang-jwt/jwt"
)

// Secret key b eing used to sign tokens
var (
	SecretKey = []byte("secret")
)

// Generates jwt  token and assign a username to it's claims and return it
func GenerateToken(username string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	// Creating a map to store the claims
	claims := token.Claims.(jwt.MapClaims)
	// Setting token claims
	claims["username"] = username
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	tokenString, err := token.SignedString(SecretKey)
	if err != nil {
		log.Fatal("Error in Generating key")
		return "", err
	}
	return tokenString, nil
}

// Parses a jwt token and returns the username in it's claims
func ParseToken(tokenStr string) (string, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return SecretKey, nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		username := claims["username"].(string)
		return username, nil
	} else {
		return "", err
	}
}
