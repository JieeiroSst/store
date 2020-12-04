package users

import (
	"github.com/jieeiro/api/base/db"
	"github.com/jieeiro/api/models/users"
)

func AccountAll(accounts []users.Users) {
	db.GetConn().Find(&accounts)
}

func AccountById(account users.Users, id int) {
	db.GetConn().Find(&account, id)
}

func CreateAccount(account users.Users) {
	db.GetConn().Create(&account)
}

func DeleteAccount(account users.Users, id int) {
	db.GetConn().First(&account, id)
	db.GetConn().Delete(&account)
}

func CheckAccount(username string) (bool, int, string) {
	var accounts []users.Users
	db.GetConn().Find(&accounts)
	for _,account:=range accounts{
		if account.Username==username{
			return true, account.Id, account.Password
		}
	}
	return false, 0, ""
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
