package categories

import (
	"time"
)

type Categories struct {
	id int
	name string
	metaTitle string
	displayOrder int
	createdAt time.Time
	createdBy string
	modifiedDate time.Time
	metaKeyword string
	metaDescription string
	status int
}
