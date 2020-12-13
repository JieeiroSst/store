package jwt

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
	"log"
	"os"
	"time"
)

type Claims struct {
	username string
	role string
	jwt.StandardClaims
}

func GenerateToken(id int,username string) (string, error) {
	e := godotenv.Load() //Load .env file
	if e != nil {
		log.Print(e)
	}
	key := os.Getenv("KEY")
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["id"]=id
	claims["username"] = username
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	tokenString, err := token.SignedString(key)
	if err != nil {
		log.Fatal("Error in Generating key")
		return "", err
	}
	return tokenString, nil
}

func ParseToken(tokenStr string) (string, error) {
	e := godotenv.Load() //Load .env file
	if e != nil {
		log.Print(e)
	}
	key := os.Getenv("KEY")
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		username := claims["username"].(string)
		return username, nil
	} else {
		return "", err
	}
}