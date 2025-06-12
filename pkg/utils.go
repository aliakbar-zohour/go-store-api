package pkg

import (
	"time"
    "fmt"
	"golang.org/x/crypto/bcrypt"
	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("your-secret-key")

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func GenerateJWT(email string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add(24 * time.Hour).Unix(),
	})
	return token.SignedString(jwtKey)
}

func ParseJWT(tokenStr string) (string, error) {
    token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
        return jwtKey, nil
    })

    if err != nil || !token.Valid {
        return "", err
    }

    claims, ok := token.Claims.(jwt.MapClaims)
    if !ok {
        return "", fmt.Errorf("invalid claims")
    }

    email, ok := claims["email"].(string)
    if !ok {
        return "", fmt.Errorf("invalid email in token")
    }

    return email, nil
}
