package abouts

import (
	"time"
)

type Abouts struct {
	Id int
	Title string
	MetaTitle string
	Description string
	Image string
	Detail string
	CreatedAt time.Time
	CreatedBy string
	ModifiedDate time.Time
	ModifiedBy string
	MetaKeyword string
	MetaDescription string
	Status int
}
