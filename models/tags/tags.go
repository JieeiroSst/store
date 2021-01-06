package tags

import "github.com/JIeeiroSst/store/models/paginations"

type Tags struct {
	Id int
	Name string
	NewId int
}

type TagsEdge struct {
	Cursor string
	Node Tags
}

type TagsConnection struct {
	PageInfo paginations.PageInfo
	Edges [] TagsEdge
	totalCount int 
}

type InputTags struct {
	Name string
	NewId int
}

