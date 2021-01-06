package news

import (
	"time"
	"github.com/JIeeiroSst/store/models/paginations"
)

type News struct {
	Id int
	Title string
	MetaTitle string
	Description string
	Image string
	CategoryId int
	Detail string
	CreatedAt time.Time
	CreatedBy string
	ModifiedData string
	ModifiedBy string
	MetaKeyWord string
	MetaDescription string
	Status int
	TopHot time.Time
	ViewCount int
	TagId int
	Email string
}

type NewsEdge struct {
	Cursor string
	Node News
}

type NewsConnection struct {
	PageInfo paginations.PageInfo
	Edges []NewsEdge
	TotalCount int 
}

type InputNews struct {
	Title string
	MetaTitle string
	Description string
	Image string
	CategoryId int
	Detail string
	CreatedAt string
	CreatedBy string
	ModifiedData string
	ModifiedBy string
	MetaKeyWord string
	MetaDescription string
	Status int
	TopHot string
	ViewCount int
	TagId int
}