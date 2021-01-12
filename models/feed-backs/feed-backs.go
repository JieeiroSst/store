package feed_backs

import (
	"time"
	"github.com/JIeeiroSst/store/models/paginations"
)

type FeedBacks struct {
	Id int
	Name string
	Phone string
	Email string
	Address string
	Content string
	CreatedAt time.Time
}

type FeedBackEdge struct {
	Cursor string
	Node FeedBacks
}

type FeedBackConnection struct {
	PageInfo paginations.PageInfo
	Edges [] FeedBackEdge
	TotalCount int
}

type InputFeedBacks struct {
	Name string
	Phone string
	Email string
	Address string
	Content string
}