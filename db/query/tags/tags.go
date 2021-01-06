package tags

import (
	db "github.com/JIeeiroSst/store/component"
	"github.com/JIeeiroSst/store/models/tags"
)

func TagsAll(tags []tags.Tags){
	db.GetConn().Find(&tags)
}

func TagsById(tags tags.Tags,id int)  {
	db.GetConn().Find(&tags,id)
}

func CreateTags(tags tags.Tags){
	db.GetConn().Create(&tags)
}

func DeleteTags(tags tags.Tags,id int){
	db.GetConn().First(&tags,id)
	db.GetConn().Delete(&tags)
}

func UpdateTags(tags tags.Tags,id int){
	db.GetConn().Where("id = ?",id).First(&tags)
	db.GetConn().Save(&tags)
}