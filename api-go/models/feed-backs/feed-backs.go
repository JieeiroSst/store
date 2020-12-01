package feed_backs

import "time"

type FeedBacks struct {
	id int
	name string
	phone string
	email string
	address string
	content string
	status int
	createdAt time.Time
}