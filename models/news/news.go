package news

import "time"

type News struct {
	id int
	title string
	metaLink string
	description string
	image string
	categoryId int
	detail string
	createdAt time.Time
	createdBy string
	modifiedData string
	status int
	topHot time.Time
	viewCount int
	tagId int
}

