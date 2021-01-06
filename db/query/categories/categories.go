package categories

import (
	db "github.com/JIeeiroSst/store/component"
	"github.com/JIeeiroSst/store/models/categories"
)

func CategoriesAll(categories []categories.Categories){
	db.GetConn().Find(&categories)
}

func CategoriesById(categories categories.Categories,id int){
	db.GetConn().Find(&categories,id)
}

func CreateCategories(categories categories.Categories){
	db.GetConn().Create(&categories)
}

func DeleteCategories(categories categories.Categories,id int){
	db.GetConn().First(&categories,id)
	db.GetConn().Delete(&categories)
}

func UpdateCategories(categories categories.Categories,id int){
	db.GetConn().Where("id = ?",id).First(&categories)
	db.GetConn().Save(&categories)
}