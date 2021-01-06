package abouts

import (
	db "github.com/JIeeiroSst/store/component"
	"github.com/JIeeiroSst/store/models/abouts"
)

func AboutsAll(abouts []abouts.Abouts){
	db.GetConn().Find(&abouts)
}

func AboutsById(abouts abouts.Abouts,id int){
	db.GetConn().Find(&abouts,id)
}

func CreateAbouts(abouts abouts.Abouts){
	db.GetConn().Create(&abouts)
}

func DeleteAbouts(abouts abouts.Abouts,id int){
	db.GetConn().First(&abouts,id)
	db.GetConn().Delete(&abouts)
}

func UpdateAbouts(abouts abouts.Abouts,id int){
	db.GetConn().Where("id = ?",id).First(&abouts)
	db.GetConn().Save(&abouts)
}