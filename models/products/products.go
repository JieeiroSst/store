package products

import (
	"time"
	"github.com/JIeeiroSst/store/models/paginations"
)

type Products struct {
	Id int
	Code string
	Name string
	Title string
	Description string
	Images string
	Price int
	Vat int
	CategoryId int
	Detail string
	CreatedDate time.Time
	CreatedBy string
	ModifiedDate time.Time
	ModifiedBy string
	Status int
	TopHot time.Time
	ViewCount int
}

type ProductEdge struct {
	Cursor string
	Node Products
}

type ProductConnection struct {
	PageInfo paginations.PageInfo
	Edges []ProductEdge
	TotalCount int
}

type InputProduct struct {
	Code string
	Name string
	Title string
	Description string
	Images string
	Price int
	Vat int
	CategoryId int
	Detail string
	CreatedDate time.Time
	CreatedBy string
	ModifiedDate time.Time
	ModifiedBy string
	Status int
	TopHot time.Time
}