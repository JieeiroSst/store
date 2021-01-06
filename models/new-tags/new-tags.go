package new_tags

import "github.com/JIeeiroSst/store/models/paginations"

type NewTag struct {
	Id int
	TagId int
	Name string
}

type NewTagsEdge struct {
	cursor string
	node NewTag
}

type NewTagsConnection struct {
	PageInfo paginations.PageInfo
	Edges []NewTagsEdge
	TotalCount int
}

type InputNewTag struct {
	TagId int
	Name string
}