package feed_backs

import "time"

type FeedBacks struct {
	Id int
	Name string
	Phone string
	Email string
	Address string
	Content string
	Status int
	CreatedAt time.Time
}
