package news

import (
	db "github.com/JIeeiroSst/store/component"
	"github.com/JIeeiroSst/store/models/news"
)

func NewAll(news []news.News){
	db.GetConn().Find(&news)
}

func NewById(news news.News,id int){
	db.GetConn().Find(&news,id)
}

func CreateNews(news news.News){
	db.GetConn().Create(&news)
}

func DeleteNews(news news.News,id int){
	db.GetConn().First(&news,id)
	db.GetConn().Delete(&news)
}

func UpdateNews(news news.News,id int){
	db.GetConn().Where("id = ?",id).First(&news)
	db.GetConn().Save(&news)
}