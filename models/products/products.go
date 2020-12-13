package products

import "time"

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

