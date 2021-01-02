package profiles

import (
	db "github.com/JIeeiroSst/store/component"
	"github.com/JIeeiroSst/store/models/profiles"
)

func ProfileAll(profiles []profiles.Profiles){
	db.GetConn().Find(&profiles)
}

func ProfileById(profile profiles.Profiles,id int){
	db.GetConn().Find(&profile,id)
}

func CreateProfile(profile profiles.Profiles){
	db.GetConn().Create(&profile)
}

func DeleteProfile(profile profiles.Profiles,id int){
	db.GetConn().First(&profile,id)
	db.GetConn().Delete(&profile)
}

func UpdateProfile(profile profiles.Profiles,id int){
	db.GetConn().Where("userid = ?",id).First(&profile)
	db.GetConn().Save(&profile)
}