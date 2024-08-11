package utils

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	"github.com/nuublx/react-go-todo-app/types"
)

var jwtSecret []byte

type Claims struct {
	ID          string `json:"id"`
	Username    string `json:"username"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phonenumber"`

	jwt.RegisteredClaims
}

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf(fmt.Sprintf("error: %s", err.Error()))
	}

	secretKey := os.Getenv("JWT_SECRET")
	if secretKey == "" {
		log.Fatal("JWT_SECRET is not set in .env file")
	}

	jwtSecret = []byte(secretKey)

}

// GenerateJWT generates a new JWT token
func GenerateJWT(user types.User) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		ID:          user.ID.Hex(),
		Username:    user.UserName,
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber,

		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

// ValidateJWT validates the JWT token
func ValidateJWT(tokenStr string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, err
	}
	return claims, nil
}
