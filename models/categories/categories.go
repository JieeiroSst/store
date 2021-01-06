package categories

import (
	"time"
	"github.com/JIeeiroSst/store/models/paginations"
)

type Categories struct {
	Id int
	Name string
	MetaTitle string
	DisplayOrder int
	CreatedAt time.Time
	CreatedBy string
	ModifiedDate time.Time
	MetaKeyword string
	MetaDescription string
	Status int
}

type CategoriesEdge struct {
	Cursor string
	Node Categories
}

type CategoriesConnection struct {
	PageInfo paginations.PageInfo
	Edges []CategoriesEdge
	TotalCount int
}

type InoutCategories struct {
	Name string
	MetaTitle string
	DisplayOrder int
	CreatedAt time.Time
	CreatedBy string
	ModifiedDate time.Time
	MetaKeyword string
	MetaDescription string
	Status int
}