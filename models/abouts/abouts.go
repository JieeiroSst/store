package abouts

import (
	"time"
)

type Abouts struct {
	id int
	title string
	metaTitle string
	description string
	image string
	detail string
	createdAt time.Time
	createdBy string
	modifiedDate time.Time
	modifiedBy string
	metaKeyword string
	metaDescription string
	status int
}

