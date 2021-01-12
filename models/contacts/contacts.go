package contacts

import "github.com/JIeeiroSst/store/models/paginations"

type Contacts struct {
	Id int
	Content string
}

type ContactEdge struct {
	Cursor string
	node Contacts
}

type ContactConnection struct {
	PageInfo paginations.PageInfo
	Edges []ContactEdge
	TotalCount int
}

type InputContact struct {
	Content string
}