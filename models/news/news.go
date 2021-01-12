package news

import (
	"github.com/JIeeiroSst/store/models/paginations"
	"time"
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
	TopHot time.Time
	ViewCount int
	TagId int
	Email string
	Content string
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
	TopHot string
	ViewCount int
	TagId int
	Content string
}