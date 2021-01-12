package sliders

import "time"

type Sliders struct {
	Id int
	Image string
	DisplayOrder int
	Link string
	Description string
	CreatedAt time.Time
	CreatedBy string
	ModifiedAt time.Time
	ModifiedBy string
}

type InputSliders struct {
	Image string
	DisplayOrder int
	Link string
	Description string
	CreatedAt time.Time
	CreatedBy string
	ModifiedAt time.Time
	ModifiedBy string
}