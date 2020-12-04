package news

import "time"

type News struct {
	Id int
	Title string
	MetaLink string
	Description string
	Image string
	CategoryId int
	Detail string
	CreatedAt time.Time
	CreatedBy string
	ModifiedData string
	Status int
	TopHot time.Time
	ViewCount int
	TagId int
}

