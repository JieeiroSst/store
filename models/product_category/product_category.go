package product_category

import "time"

type ProductCategory struct {
	id int
	name string
	metaTitle string
	displayTitle string
	displayOrder int
	title string
	createdAt time.Time
	createdBy string
	modifiedDate time.Time
	modifiedBy string
	description string
	status int
}

