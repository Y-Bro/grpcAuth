package jwt

import (
	"log"
	"time"

	"github.com/form3tech-oss/jwt-go"
)

var (
	SecretKey = []byte("secret")
)

func GenerateToken(username string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["username"] = username

	claims["exp"] = time.Now().Add(24 * time.Hour).Unix()

	tokenStr, err := token.SignedString(SecretKey)

	if err != nil {
		log.Fatalln("Error while siging token")
		return "", err
	}

	return tokenStr, nil
}

func ParseToken(tokenStr string) (string, error) {
	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		return SecretKey, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		username := claims["username"].(string)

		return username, nil
	} else {
		return "", err
	}
}
