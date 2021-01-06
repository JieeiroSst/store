package product_category

import "time"

type ProductCategory struct {
	Id int
	Name string
	MetaTitle string
	DisplayTitle string
	DisplayOrder int
	Title string
	CreatedAt time.Time
	CreatedBy string
	ModifiedDate time.Time
	ModifiedBy string
	Description string
	Status int
}

type InputProductCategory struct {
	Name string
	MetaTitle string
	DisplayTitle string
	DisplayOrder int
	Title string
	CreatedAt time.Time
	CreatedBy string
	ModifiedDate time.Time
	ModifiedBy string
	Description string
	Status int
}