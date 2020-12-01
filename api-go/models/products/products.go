package products

import "time"

type Products struct {
	id int
	code string
	name string
	title string
	description string
	images string
	price int
	vat int
	categoryId int
	detail string
	createdDate time.Time
	createdBy string
	modifiedDate time.Time
	modifiedBy string
	status int
	topHot time.Time
	viewCount int
}
