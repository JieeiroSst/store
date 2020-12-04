package users

import (
	"github.com/jieeiro/api/base/jwt"
	"github.com/jieeiro/api/services/users"
	"github.com/jieeiro/api/utils/hash"
	models "github.com/jieeiro/api/models/users"
	"log"
)

func Login(username,password string) string{
	check, id, hashPassword := users.CheckAccount(username)
	if check == false {
		return "User does not exist"
	}
	if checkPass := hash.CheckPassowrd(password, hashPassword); checkPass != nil {
		return "password entered incorrectly"
	}
	token, _ := jwt.GenerateToken(id, username)
	return token
}

func SignUp(username,password string) string{
	check := users.CheckAccountExists(username)
	if check == false {
		return "user already exists"
	}
	hashPassword, err := hash.HashPassword(password)
	if err != nil {
		log.Println("error server", err)
	}
	account := models.Users{
		Username: username,
		Password: hashPassword,
	}
	users.CreateAccount(account)
	return username+":"+hashPassword
}