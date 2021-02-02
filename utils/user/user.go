package user

import (
	db "github.com/JIeeiroSst/store/component"
	users "github.com/JIeeiroSst/store/models"
	"github.com/JIeeiroSst/store/utils/hash"
	"github.com/JIeeiroSst/store/utils/jwt"
	"log"
)

func CheckAccount(username string) (bool, int, string,string) {
	var accounts []users.Users
	db.GetConn().Find(&accounts)
	for _,account:=range accounts{
		if account.Username==username{
			return true, account.ID, account.Password,account.Permission
		}
	}
	return false, 0, "",""
}

func CheckAccountExists(username string) bool {
	var account [] users.Users
	db.GetConn().Find(&account)
	for _,item:=range account{
		if item.Username == username {
			return false
		}
	}
	return true
}

func Login(username,password string) string{
	check, id, hashPassword,permission := CheckAccount(username)
	if check == false {
		return "User does not exist"
	}
	if checkPass := hash.CheckPassowrd(password, hashPassword); checkPass != nil {
		return "password entered incorrectly"
	}
	token, _ := jwt.GenerateToken(id, username,permission)
	return token
}

func SignUp(username,password string) string{
	check := CheckAccountExists(username)
	if check == false {
		return "user already exists"
	}
	hashPassword, err := hash.HashPassword(password)
	if err != nil {
		log.Println("error server", err)
	}

	if username == "admin"{
		account := users.Users{
			Username:   username,
			Password:   hashPassword,
			Permission: "private",
		}
		_ = db.GetConn().Create(&account)
	}else {
		account := users.Users{
			Username:   username,
			Password:   hashPassword,
			Permission: "public",
		}
		_ = db.GetConn().Create(&account)
	}

	return username+":"+hashPassword
}
