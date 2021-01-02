package jwt

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
	"log"
	"os"
	"time"
)

func GenerateToken(id int,username string) (string, error) {
	e := godotenv.Load()
	if e != nil {
		log.Print(e)
	}
	key := os.Getenv("KEY")
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_id"] = id
	atClaims["username"]=username
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(key))
	if err != nil {
		return "", err
	}
	return token, nil
}

func ParseToken(tokenStr string) (string, error) {
	e := godotenv.Load()
	if e != nil {
		log.Print(e)
	}
	key := os.Getenv("KEY")
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println(claims)
		username := claims["username"].(string)
		return username, nil
	} else {
		return "Missing Authentication Token", err
	}
}