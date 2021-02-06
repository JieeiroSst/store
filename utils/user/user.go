package user

import (
	"errors"
	db "github.com/JIeeiroSst/store/component"
	users "github.com/JIeeiroSst/store/models"
	"github.com/JIeeiroSst/store/utils/hash"
	"github.com/JIeeiroSst/store/utils/jwt"
	"log"
)

func CheckAccount(username string) (int, string,string,bool) {
	var accounts []users.Users
	_ = db.GetConn().Find(&accounts)
	for _,account:=range accounts{
		if account.Username == username{
			return account.ID, account.Password,account.Permission, true
		}
	}
	return  0, "","", false
}

func CheckAccountExists(username string) bool {
	var account [] users.Users
	_ = db.GetConn().Find(&account)
	for _,item:=range account{
		if item.Username == username {
			return false
		}
	}
	return true
}

func Login(username,password string) (string,string,string,error){
	id, hashPassword,permission,check := CheckAccount(username)
	if check == false {
		return "User does not exist","","", errors.New("not exist")
	}
	if checkPass := hash.CheckPassowrd(password, hashPassword); checkPass != nil {
		return "password entered incorrectly","","", errors.New("incorrectly")
	}
	token, _ := jwt.GenerateToken(id, username,permission)
	return "logged in successfully",token,permission ,nil
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
