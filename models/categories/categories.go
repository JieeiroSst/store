package categories

import (
	"time"
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

