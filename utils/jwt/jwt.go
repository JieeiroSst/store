package jwt

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
	"log"
	"os"
	"time"
)

func GenerateToken(id int,username string,permission string) (string, error) {
	e := godotenv.Load()
	if e != nil {
		log.Print(e)
	}
	key := os.Getenv("KEY")
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_id"] = id
	atClaims["username"]=username
	atClaims["permission"]=permission
	atClaims["exp"] = time.Now().Add(time.Hour * 60 *60).Unix()
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
		username := claims["username"].(string)
		return username, nil
	} else {
		return "Missing Authentication Token", err
	}
}

func GetPermission(tokenStr string)(string,string){
	key := os.Getenv("KEY")
	token, _ := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		username:=claims["username"].(string)
		permission := claims["permission"].(string)
		return username,permission
	} else {
		return "Missing Authentication Token","Missing Authentication Token"
	}
	return "",""
}

func GetIdUser(tokenStr string) float64 {
	key := os.Getenv("KEY")
	token, _ := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		id:=claims["user_id"].(float64)
		return id
	} else {
		return -1
	}
	return 0
}